#!/usr/bin/env bash

if [[ -n "$VIAM_TEST_FAIL_RUN_FIRST" ]]; then
    exit 1
fi

docker pull mongo:6

cat << EOF
-------------------------------------
Congratulations!

The setup script ran successfully!

This message is obnoxiously large for
testing purposes.

Sincerely,
First Run Script
-------------------------------------
EOF

exit 0
