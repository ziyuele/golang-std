#!/bin/bash

go get -u
CUR_DIR=``
function do_build() {
   go build
}

if [[ -z go.mod ]]; then
  do_build
else
   do_build
fi