"""
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""

import sys
from distutils.util import strtobool
from time import sleep

from fabric.api import cd, env, execute, local, run, settings, sudo
from fabric.contrib.files import exists
from fabric.operations import get

sys.path.append('../../../orc8r')
import tools.fab.dev_utils as dev_utils
import tools.fab.pkg as pkg
from fabric.api import lcd
from tools.fab.hosts import ansible_setup, split_hoststring, vagrant_setup
from tools.fab.vagrant import setup_env_vagrant

"""
Magma Gateway packaging tool:

Magma packages released to different channels have different version schemes.

    - dev: used for development.

    - test: used for Continuous Integration (CI). Packages in the `test`
            channel should be built and released automatically.

# HOWTO build magma.deb

1. From your laptop, update the magma version number in `release/build-magma.sh`

2. From the dev VM, build the magma package. Dependency packages are recorded
in `release/magma.lockfile`

    fab dev package
    # optionally upload to aws (if you are configured for it)
    fab dev package upload_to_aws
"""

GATEWAY_IP_ADDRESS = "192.168.60.142"
AGW_ROOT = "$MAGMA_ROOT/lte/gateway"
AGW_PYTHON_ROOT = "$MAGMA_ROOT/lte/gateway/python"
FEG_INTEG_TEST_ROOT = AGW_PYTHON_ROOT + "/integ_tests/federated_tests"
FEG_INTEG_TEST_DOCKER_ROOT = FEG_INTEG_TEST_ROOT + "/docker"
ORC8R_AGW_PYTHON_ROOT = "$MAGMA_ROOT/orc8r/gateway/python"
AGW_INTEG_ROOT = "$MAGMA_ROOT/lte/gateway/python/integ_tests"
DEFAULT_CERT = "$MAGMA_ROOT/.cache/test_certs/rootCA.pem"
DEFAULT_PROXY = "$MAGMA_ROOT/lte/gateway/configs/control_proxy.yml"
TEST_SUMMARY_GLOB = "/var/tmp/test_results/*.xml"

# Look for keys as specified in our ~/.ssh/config
env.use_ssh_config = True
# Disable ssh known hosts to resolve key errors
# with multiple vagrant boxes in use.
env.disable_known_hosts = True


def _setup_vm(host, name, ansible_role, ansible_file, destroy_vm, provision_vm):
    ip = None
    if not host:
        host = vagrant_setup(
            name, destroy_vm, force_provision=provision_vm,
        )
    else:
        ansible_setup(host, ansible_role, ansible_file)
        ip = host.split('@')[1].split(':')[0]
    return host, ip


def _setup_gateway(gateway_host, name, ansible_role, ansible_file, destroy_vm, provision_vm):
    gateway_host, gateway_ip = _setup_vm(gateway_host, name, ansible_role, ansible_file, destroy_vm, provision_vm)
    if gateway_ip == None:
        gateway_ip = GATEWAY_IP_ADDRESS
    return gateway_host, gateway_ip


def integ_test_deb_installation(
    gateway_host=None, test_host=None, trf_host=None,
    destroy_vm='True', provision_vm='True',
):
    destroy_vm = bool(strtobool(destroy_vm))
    provision_vm = bool(strtobool(provision_vm))

    # Set up the gateway: use the provided gateway if given, else default to the
    # vagrant machine
    _setup_gateway(gateway_host, "magma_deb", "deb", "magma_deb.yml", destroy_vm, provision_vm)


