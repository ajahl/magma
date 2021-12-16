# Copyright 2021 The Magma Authors.

# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree.

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""All external repositories not specific to a language"""

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def protobuf():
    http_archive(
        # The name is protobuf here as that is what prometheus-cpp expects
        # See https://github.com/jupp0r/prometheus-cpp.git @ d8326b2bba945a435f299e7526c403d7a1f68c1f
        name = "protobuf",
        strip_prefix = "protobuf-3.19.1",
        #sha256 = "6aff9834fd7c540875e1836967c8d14c6897e3785a2efac629f69860fb7834ff",
        urls = ["https://github.com/protocolbuffers/protobuf/archive/refs/tags/v3.19.1.tar.gz"],
    )

def grpc():
    http_archive(
        name = "rules_proto_grpc",
        sha256 = "8383116d4c505e93fd58369841814acc3f25bdb906887a2023980d8f49a0b95b",
        strip_prefix = "rules_proto_grpc-4.1.0",
        urls = ["https://github.com/rules-proto-grpc/rules_proto_grpc/archive/4.1.0.tar.gz"],
    )
