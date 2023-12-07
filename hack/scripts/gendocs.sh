#!/usr/bin/env bash

version=${1}

sed "s/0.0.1/${version}/g" main.go > docs.go

swag init -g docs.go

rm -rf docs.go
