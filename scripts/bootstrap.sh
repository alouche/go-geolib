#!/bin/sh -e

if [ ! -h src/github.com/alouche/go-geolib ]; then
  mkdir -p src/github.com/alouche/
  ln -s ../../.. src/github.com/alouche/go-geolib
fi

export GOBIN=${PWD}/bin
export GOPATH=${PWD}
