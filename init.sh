#!/bin/bash

export GOPATH=`pwd`
mkdir bin
go get -v -u github.com/spf13/cobra
