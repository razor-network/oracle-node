#!/bin/bash

echo "Building docker image"
docker build --file Dockerfile.testing -t razor-test .

echo "Stating running test cases"
docker run --rm -v $(pwd):/test --name go  razor-test go-acc ./... --ignore razor/accounts/mocks --ignore razor/cmd/mocks --ignore razor/utils/mocks --ignore pkg --output /test/coverage.txt