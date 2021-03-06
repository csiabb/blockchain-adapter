# Copyright ArxanChain Ltd. 2020 All Rights Reserved.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#                  http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# 

BINARY=blockchain-adapter
VERSION=0.1.5
BUILD=`date +%FT%T%z`

BUILDPATH=build
BUILDBINPATH=${BUILDPATH}/bin

PACKAGES=`go list ./...`
GOFILES=`find . -name "*.go"`

default:
	@CGO_ENABLED=0 go build -o ${BUILDBINPATH}/${BINARY}
	@cp sampleconfig/blockchain-adapter.yaml ${BUILDBINPATH}

list:
	@echo ${PACKAGES}
	@echo ${GOFILES}

fmt:
	@gofmt -s -w ${GOFILES}

fmt-check:
	@diff=$$(gofmt -s -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

mod:
	@go mod tidy

test:
	@go test -cover -v -tags integration ./...

vet:
	@go vet $(PACKAGES)

lint:
	@golint $(PACKAGES)

docker: default
	@docker build -t csiabb/blockchain-adapter -f dockerfile/Dockerfile ./
	@docker tag csiabb/blockchain-adapter:latest csiabb/blockchain-adapter:$(VERSION)

clean:
	@rm -rf ${BUILDPATH}

.PHONY: default fmt fmt-check mod test vet lint docker clean
