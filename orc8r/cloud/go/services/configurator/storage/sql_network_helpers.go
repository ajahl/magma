/*
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package storage

import (
	"database/sql"
	"fmt"
	"sort"

	sq "github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
	"google.golang.org/protobuf/proto"

	"magma/orc8r/cloud/go/sqorc"
)

func getNetworkQueryColumns(criteria *NetworkLoadCriteria) []string {
	criteriaCopy := proto.Clone(criteria).(*NetworkLoadCriteria)
	ret := []string{
		fmt.Sprintf("%s.%s", networksTable, nwIDCol),
		fmt.Sprintf("%s.%s", networksTable, nwTypeCol),
	}
	if criteriaCopy.LoadMetadata {
		ret = append(
			ret,
			fmt.Sprintf("%s.%s", networksTable, nwNameCol),
			fmt.Sprintf("%s.%s", networksTable, nwDescCol),
		)
	}
	if criteriaCopy.LoadConfigs {
		ret = append(
			ret,
			fmt.Sprintf("%s.%s", networkConfigTable, nwcTypeCol),
			fmt.Sprintf("%s.%s", networkConfigTable, nwcValCol),
		)
	}
	ret = append(ret, fmt.Sprintf("%s.%s", networksTable, nwVerCol))
	return ret
}

func (store *sqlConfiguratorStorage) getLoadNetworksSelectBuilder(filter *NetworkLoadFilter, criteria *NetworkLoadCriteria) sq.SelectBuilder {
	filterCopy := proto.Clone(filter).(*NetworkLoadFilter)
	criteriaCopy := proto.Clone(criteria).(*NetworkLoadCriteria)

	selectBuilder := store.builder.Select(getNetworkQueryColumns(criteriaCopy)...).From(networksTable)
	if funk.NotEmpty(filterCopy.Ids) {
		selectBuilder = selectBuilder.Where(sq.Eq{
			fmt.Sprintf("%s.%s", networksTable, nwIDCol): filterCopy.Ids,
		})
	} else if funk.NotEmpty(filterCopy.TypeFilter) {
		selectBuilder = selectBuilder.Where(sq.Eq{fmt.Sprintf("%s.%s", networksTable, nwTypeCol): filterCopy.TypeFilter.Value})
	}
	return selectBuilder
}

func scanNetworkRows(rows *sql.Rows, loadCriteria *NetworkLoadCriteria) (map[string]*Network, []string, error) {
	loadCriteriaCopy := proto.Clone(loadCriteria).(*NetworkLoadCriteria)
	// Pointer values because we're modifying .Config in-place
	loadedNetworksByID := map[string]*Network{}
	for rows.Next() {
		nwResult, err := scanNextNetworkRow(rows, loadCriteriaCopy)
		if err != nil {
			return nil, nil, err
		}

		existingNetwork, wasLoaded := loadedNetworksByID[nwResult.ID]
		if wasLoaded {
			for k, v := range nwResult.Configs {
				existingNetwork.Configs[k] = v
			}
		} else {
			loadedNetworksByID[nwResult.ID] = nwResult
		}
	}

	// Sort map keys so we return deterministically
	loadedNetworkIDs := funk.Keys(loadedNetworksByID).([]string)
	sort.Strings(loadedNetworkIDs)
	return loadedNetworksByID, loadedNetworkIDs, nil
}

func scanNextNetworkRow(rows *sql.Rows, criteria *NetworkLoadCriteria) (*Network, error) {
	var id string
	var networkType, name, description sql.NullString
	var cfgType sql.NullString
	var cfgValue []byte
	var version uint64

	criteriaCopy := proto.Clone(criteria).(*NetworkLoadCriteria)
	scanArgs := []interface{}{
		&id,
		&networkType,
	}
	if criteriaCopy.LoadMetadata {
		scanArgs = append(scanArgs, &name, &description)
	}
	if criteriaCopy.LoadConfigs {
		scanArgs = append(scanArgs, &cfgType, &cfgValue)
	}
	scanArgs = append(scanArgs, &version)

	err := rows.Scan(scanArgs...)
	if err != nil {
		return &Network{}, fmt.Errorf("error while scanning network row: %s", err)
	}

	ret := &Network{ID: id, Type: nullStringToValue(networkType), Name: nullStringToValue(name), Description: nullStringToValue(description), Configs: map[string][]byte{}, Version: version}
	if criteriaCopy.LoadConfigs && cfgType.Valid {
		ret.Configs[cfgType.String] = cfgValue
	}
	return ret, nil
}

func getNetworkIDsNotFound(networksByID map[string]*Network, queriedIDs []string) []string {
	ret := []string{}
	for _, id := range queriedIDs {
		if _, ok := networksByID[id]; !ok {
			ret = append(ret, id)
		}
	}
	sort.Strings(ret)
	return ret
}

func (store *sqlConfiguratorStorage) doesNetworkExist(id string) (bool, error) {
	var count int
	err := store.builder.Select("COUNT(1)").
		From(networksTable).
		Where(sq.Eq{"id": id}).
		RunWith(store.tx).
		QueryRow().
		Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error checking if network id %s exists: %s", id, err)
	}

	return count > 0, nil
}

func validateNetworkUpdates(updates []*NetworkUpdateCriteria) error {
	updatesByID := funk.ToMap(updates, "ID").(map[string]*NetworkUpdateCriteria)
	if len(updatesByID) < len(updates) {
		return errors.New("multiple updates for a single network are not allowed")
	}
	return nil
}

func (store *sqlConfiguratorStorage) updateNetwork(update *NetworkUpdateCriteria, stmtCache *sq.StmtCache) error {
	updateCopy := proto.Clone(update).(*NetworkUpdateCriteria)
	// Update the network table first
	updateBuilder := store.builder.Update(networksTable).Where(sq.Eq{nwIDCol: updateCopy.ID})
	if updateCopy.NewName != nil {
		updateBuilder = updateBuilder.Set(nwNameCol, stringPtrToVal(updateCopy.NewName))
	}
	if updateCopy.NewDescription != nil {
		updateBuilder = updateBuilder.Set(nwDescCol, stringPtrToVal(updateCopy.NewDescription))
	}
	if updateCopy.NewType != nil {
		updateBuilder = updateBuilder.Set(nwTypeCol, stringPtrToVal(updateCopy.NewType))
	}
	updateBuilder = updateBuilder.Set(nwVerCol, sq.Expr(fmt.Sprintf("%s.%s+1", networksTable, nwVerCol)))
	_, err := updateBuilder.RunWith(stmtCache).Exec()
	if err != nil {
		return errors.Wrapf(err, "error updating network %s", updateCopy.ID)
	}

	// Sort config keys for deterministic behavior on upserts
	configUpdateTypes := funk.Keys(updateCopy.ConfigsToAddOrUpdate).([]string)
	sort.Strings(configUpdateTypes)
	for _, configType := range configUpdateTypes {
		configValue := updateCopy.ConfigsToAddOrUpdate[configType]

		// INSERT INTO %s (network_id, type, value) VALUES ($1, $2, $3)
		// ON CONFLICT (network_id, type) DO UPDATE SET value = $4
		_, err := store.builder.Insert(networkConfigTable).
			Columns(nwcIDCol, nwcTypeCol, nwcValCol).
			Values(updateCopy.ID, configType, configValue).
			OnConflict(
				[]sqorc.UpsertValue{{Column: nwcValCol, Value: configValue}},
				nwcIDCol, nwcTypeCol,
			).
			RunWith(stmtCache).
			Exec()
		if err != nil {
			return errors.Wrapf(err, "error updating config %s on network %s", configType, updateCopy.ID)
		}
	}

	// Finally delete configs
	if funk.IsEmpty(updateCopy.ConfigsToDelete) {
		return nil
	}

	orClause := make(sq.Or, 0, len(updateCopy.ConfigsToDelete))
	for _, configType := range updateCopy.ConfigsToDelete {
		orClause = append(orClause, sq.Eq{nwcIDCol: updateCopy.ID, nwcTypeCol: configType})
	}
	_, err = store.builder.Delete(networkConfigTable).Where(orClause).RunWith(store.tx).Exec()
	if err != nil {
		return errors.Wrapf(err, "failed to delete configs for network %s", updateCopy.ID)
	}

	return nil
}

func stringPtrToVal(value *wrappers.StringValue) interface{} {
	if value == nil {
		return ""
	}
	return value.Value
}

func nullStringToValue(in sql.NullString) string {
	if in.Valid {
		return in.String
	}
	return ""
}
