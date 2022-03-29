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

package main

import (
	"context"
	"log"

	"github.com/golang/glog"

	"magma/orc8r/cloud/go/blobstore"
	"magma/orc8r/cloud/go/obsidian"
	swagger_protos "magma/orc8r/cloud/go/obsidian/swagger/protos"
	swagger_servicers "magma/orc8r/cloud/go/obsidian/swagger/servicers/protected"
	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/service"
	"magma/orc8r/cloud/go/services/tenants"
	"magma/orc8r/cloud/go/services/tenants/obsidian/handlers"
	"magma/orc8r/cloud/go/services/tenants/protos"
	servicers "magma/orc8r/cloud/go/services/tenants/servicers/protected"
	"magma/orc8r/cloud/go/services/tenants/servicers/storage"
	"magma/orc8r/cloud/go/sqorc"
	storage2 "magma/orc8r/cloud/go/storage"
	"magma/orc8r/cloud/go/tracing"
)

func main() {
	tp := tracing.Init("tenants")
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	srv, err := service.NewOrchestratorService(orc8r.ModuleName, tenants.ServiceName)
	if err != nil {
		glog.Fatalf("Error creating tenants service %s", err)
	}
	sqldriver := tracing.InitTracingDBHook(storage2.GetSQLDriver())
	db, err := sqorc.Open(sqldriver, storage2.GetDatabaseSource())
	if err != nil {
		glog.Fatalf("Failed to connect to database: %s", err)
	}
	factory := blobstore.NewSQLStoreFactory(tenants.DBTableName, db, sqorc.GetSqlBuilder())
	err = factory.InitializeFactory()
	if err != nil {
		glog.Fatalf("Error initializing tenant database: %s", err)
	}
	store := storage.NewBlobstoreStore(factory)

	server, err := servicers.NewTenantsServicer(store)
	if err != nil {
		glog.Fatalf("Error creating tenants server: %s", err)
	}
	protos.RegisterTenantsServiceServer(srv.GrpcServer, server)

	swagger_protos.RegisterSwaggerSpecServer(srv.GrpcServer, swagger_servicers.NewSpecServicerFromFile(tenants.ServiceName))

	obsidian.AttachHandlers(srv.EchoServer, handlers.GetObsidianHandlers())

	err = srv.Run()
	if err != nil {
		glog.Fatalf("Error running service: %s", err)
	}
}
