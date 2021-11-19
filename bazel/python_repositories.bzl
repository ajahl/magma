# Copyright 2021 The Magma Authors.

# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree.

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""All external repositories used for Python dependencies"""

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@rules_python//python:pip.bzl", "pip_parse")
load("//bazel:python_packages.bzl", "PYTHON_PACKAGE")

def python_repositories(name):
    http_archive(
        name = PYTHON_PACKAGE.name,
        build_file_content = PYTHON_PACKAGE.build_file_content,
        patch_cmds = PYTHON_PACKAGE.patch_cmds,
        sha256 = PYTHON_PACKAGE.sha256,
        strip_prefix = PYTHON_PACKAGE.strip_prefix,
        urls = PYTHON_PACKAGE.urls,
    )

    native.register_toolchains("//bazel:py_toolchain")

    pip_parse(
        name = "python_deps",
        extra_pip_args = ["--require-hashes"],
        python_interpreter_target = "@python_interpreter//:python_bin",
        requirements_lock = "//lte/gateway/python:requirements.txt",
        visibility = ["//visibility:public"],
    )

    git_repository(
        name = "subpar",
        remote = "https://github.com/google/subpar",
        tag = "2.0.0",
        visibility = ["//visibility:public"],
    )
