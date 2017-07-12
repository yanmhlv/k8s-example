#!/usr/bin/env bash

VERSION=${VERSION:=1.0}
NAME='example'

docker build -t $NAME:$VERSION .

