# Copyright 2021 The Magma Authors.

# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree.

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""Python repository"""

PY_VERSION = "3.8.5"

BUILD_DIR = "/tmp/bazel/external/python_{0}".format(PY_VERSION)

_py_from_source_build_file_content = """
exports_files(["python_bin"])
filegroup(
    name = "files",
    srcs = glob(["bazel_install/**"], exclude = ["**/* *"]),
    visibility = ["//visibility:public"],
)
"""

_py_configure = """
    if [[ "$OSTYPE" == "darwin"* ]]; then
        cd {0} && ./configure --prefix={0}/bazel_install --with-openssl=$(brew --prefix openssl)
    else
        cd {0} && ./configure --prefix={0}/bazel_install
    fi
    """.format(BUILD_DIR)

def _py_patch_cmds(configure):
    return [
        "mkdir -p {0}".format(BUILD_DIR),
        "cp -r * {0}".format(BUILD_DIR),
        _py_configure,
        "cd {0} && make install".format(BUILD_DIR),
        "rm -rf * && mv {0}/* .".format(BUILD_DIR),
        "ln -s bazel_install/bin/python3 python_bin",
    ]

PYTHON_PACKAGE = struct(
    name = "python_interpreter",
    sha256 = "e3003ed57db17e617acb382b0cade29a248c6026b1bd8aad1f976e9af66a83b0",
    strip_prefix = "Python-{0}".format(PY_VERSION),
    urls = ["https://www.python.org/ftp/python/{0}/Python-{0}.tar.xz".format(PY_VERSION)],
    build_file_content = _py_from_source_build_file_content,
    patch_cmds = _py_patch_cmds(_py_configure),
    python_interpreter = "python3_bin",
)
