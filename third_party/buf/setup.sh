#!/bin/bash

shopt -s nullglob

cd "$( dirname "$0" )" || exit

function die(){
  echo "$@"
  exit 1
}

if [ -z "$BUF_VERSION" ]; then
  die BUF_VERSION not set
fi
BUF_URL="https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(uname -s)-$(uname -m)"
mkdir -p bin
curl -sSL $BUF_URL -o "bin/buf-${BUF_VERSION}" && chmod +x bin/buf-$BUF_VERSION
