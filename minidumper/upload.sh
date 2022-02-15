#!usr/bin

export SENTRY_AUTH_TOKEN
export SENTRY_URL=https://sentry.io/
export SENTRY_ORG=
export SENTRY_PROJECT=magma-cpp-native
export SENTRY_LOG_LEVEL=info
export _FOLDER_PATH=/home/vagrant/magma/bazel-bin/minidumper/minidumper

# if which sentry-cli >/dev/null; then
# ERROR=$(sentry-cli upload-dif "$_FOLDER_PATH" 2>&1 >/dev/null)
# if [ ! $? -eq 0 ]; then
# echo "warning: sentry-cli - $ERROR"
# fi
# else
# echo "warning: sentry-cli not installed, download from https://github.com/getsentry/sentry-cli/releases"
# fi

sudo cp /home/vagrant/magma/bazel-bin/minidumper/minidumper .
chmod 755 minidumper

sentry-cli difutil check minidumper

objcopy --only-keep-debug minidumper minidumper.debug
objcopy  --strip-debug --strip-unneeded minidumper
objcopy --add-gnu-debuglink=minidumper.debug minidumper


# sentry-cli difutil check minidumper.debug
# sentry-cli difutil check minidumper

# sentry-cli difutil check /home/vagrant/magma/bazel-bin/minidumper/minidumper

# sentry-cli difutil bundle-sources  /home/vagrant/magma/bazel-bin/minidumper/minidumper ./src

sentry-cli --auth-token $SENTRY_AUTH_TOKEN  upload-dif --log-level=info --org $SENTRY_ORG --project $SENTRY_PROJECT --include-sources ./

# sentry-cli --auth-token $SENTRY_AUTH_TOKEN  upload-dif --log-level=info --org $SENTRY_ORG --project $SENTRY_PROJECT --include-sources /home/vagrant/magma/bazel-bin/minidumper/minidumper ./src
