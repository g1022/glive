#!/usr/bin/env bash

cd test&&go test -test.v -test.run="^TestTsSmoke$" &&cd ..


