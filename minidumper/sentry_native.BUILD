# Based off of and modified from https://github.com/google/ml-metadata/blob/284af0c9d0d8467b6e69b632f4aedd0f40daac4c/ml_metadata/libmysqlclient.BUILD

load("@rules_cc//cc:defs.bzl", "cc_library")

# Source files generated with cmake.
configure_out_srcs = [
    "build/lib/libsentry.a",
    "build/lib/libbreakpad_client.a",
]

# A genrule to run cmake and generate configure_out_(srcs,hdrs).
genrule(
    name = "cmake_sentry",
    srcs = glob(
        ["**"],
    ),
    outs = configure_out_srcs,
    cmd = "\n".join([
        "export INSTALL_DIR=$$(pwd)/$(@D)/build",
        # Create a temp directory to keep the external/sentry_native directory clean
        "export TMP_DIR=$$(mktemp -d -t build.XXXXXX)",
        "mkdir -p $$TMP_DIR",
        "cp -R $$(pwd)/external/sentry_native/* $$TMP_DIR",
        "cd $$TMP_DIR",
        # See sentry-native/src/CMakeLists.txt @ GitHub for how these values are used
        # This CC=/usr/bin/gcc is important for the Magma VM
        "CC=/usr/bin/gcc cmake -B build -DCMAKE_BUILD_TYPE=RelWithDebInfo -DBUILD_SHARED_LIBS=OFF -DSENTRY_BUILD_TESTS=0 -DSENTRY_BUILD_EXAMPLES=0",
        "cmake --build build --parallel",
        "cmake --install build --prefix install --config RelWithDebInfo",
        # Copy out generated libraries into the path specified by outs
        "cp -R ./install/* $$INSTALL_DIR",
        "rm -rf $$TMP_DIR",
    ]),
)

cc_library(
    name = "sentry_h",
    srcs = ["include/sentry.h"],
    includes = ["include"],
)

cc_library(
    name = "sentry",
    # srcs / hdrs are generated by the :cmake_sentry genrule
    srcs = [":cmake_sentry"],
    linkopts = ["-lcurl -ldl"],
    linkstatic = 1,
    visibility = ["//visibility:public"],
    deps = ["sentry_h"],
)