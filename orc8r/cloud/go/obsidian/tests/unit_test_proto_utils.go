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

package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

const (
	Separator = "$$"
)

// CompareErrors checks if actualError contains relevant parts of expectedError.
// This method of comparison is necessary since string representations of proto messages contain
// non-deterministic whitespaces (https://github.com/golang/protobuf/issues/1269)
func CompareErrors(expectedError string, actualError string) error {
	// split expected error string at `separator`
	expectedErrorParts := strings.Split(expectedError, Separator)
	// check if actual message contains relevant parts of expected error
	for _, messagePart := range expectedErrorParts {
		if !strings.Contains(actualError, messagePart) {
			return fmt.Errorf("actual error does not contain all parts of expected error\nexpected: %v\nactual:   %v", strings.Join(expectedErrorParts, " "), actualError)
		}
	}
	return nil
}

// AssertMapsEqual compares two maps of with proto messages.
func AssertMapsEqual(t *testing.T, expected map[string]proto.Message, actual map[string]proto.Message) {
	assert.Equal(t, len(expected), len(actual))
	for key := range actual {
		assert.Equal(t, expected[key].String(), actual[key].String())
	}
}
