#!/bin/bash

rm -rf goravel
echo "✅ Clear old testing"
git clone git@github.com:goravel/goravel.git goravel
echo "✅ Clone Goravel"
mkdir goravel/testing
cp -r $(ls . | grep -v goravel | xargs) ./goravel/testing
cp .env ./goravel/testing
cd goravel
cp -af ./testing/stubs/* ./
echo "✅ Copy testing file"
go mod tidy
echo "✅ go mod"
echo "Testing"
go test -v ./testing/...
echo "✅ Complete"
