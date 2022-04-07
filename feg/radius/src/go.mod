// Copyright 2022 The Magma Authors.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

module fbc/cwf/radius

replace (
	fbc/lib/go/machine => ../lib/go/machine
	fbc/lib/go/radius => ../lib/go/radius
	magma/orc8r/lib/go => ../../../orc8r/lib/go
	magma/orc8r/lib/go/protos => ../../../orc8r/lib/go/protos
)

require (
	contrib.go.opencensus.io/exporter/prometheus v0.1.0
	fbc/lib/go/machine v0.0.0-00010101000000-000000000000
	fbc/lib/go/radius v0.0.0-00010101000000-000000000000
	github.com/alicebob/miniredis v2.5.0+incompatible
	github.com/donovanhide/eventsource v0.0.0-20171031113327-3ed64d21fb0b
	github.com/go-redis/redis v6.15.8+incompatible
	github.com/google/uuid v1.1.2
	github.com/mitchellh/mapstructure v1.4.3
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/stretchr/testify v1.7.0
	go.opencensus.io v0.23.0
	go.uber.org/atomic v1.7.0
	go.uber.org/zap v1.19.0
	golang.org/x/net v0.0.0-20220121210141-e204ce36a2ba
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.28.0
	magma/orc8r/lib/go v0.0.0
)

require (
	github.com/alicebob/gopher-json v0.0.0-20200520072559-a9ecdc9d1d3a // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.18.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/yuin/gopher-lua v0.0.0-20200816102855-ee81675732da // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

go 1.18
