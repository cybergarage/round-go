###################################################################
#
# Round for Go
#
# Copyright (C) Satoshi Konno 2015
#
# This is licensed under BSD-style license, see file COPYING.
#
###################################################################

PRODUCT=round
GITHUB=github.com/cybergarage/${PRODUCT}-go

PACKAGES = ${GITHUB}/${PRODUCT} ${GITHUB}/${PRODUCT}/core ${GITHUB}/${PRODUCT}/common/ ${GITHUB}/${PRODUCT}/core/rpc ${GITHUB}/${PRODUCT}/config ${GITHUB}/${PRODUCT}/log ${GITHUB}/${PRODUCT}/script ${GITHUB}/${PRODUCT}/impl
CONST_FILES = ./${PRODUCT}/version.go ./${PRODUCT}/const.go ./${PRODUCT}/impl/server_desc.go

.PHONY: ${CONST_FILES}

all: build

const: ${CONST_FILES}

./${PRODUCT}/const.go: ./${PRODUCT}/const.go.gen
	$< > $@
 
./${PRODUCT}/version.go: ./${PRODUCT}/version.go.gen
	$< > $@

./${PRODUCT}/impl/server_desc.go: ./${PRODUCT}/impl/server_desc.go.gen
	$< > $@

setup:
	go get -u github.com/cybergarage/go-net-upnp/net/upnp
	go get -u ${GITHUB}/${PRODUCT}

format:
	gofmt -w src

cgo: 
	go tool cgo ./${PRODUCT}/script/global_script_func.go

build: cgo format
	go build -v ${PACKAGES}

test: build
	go test -v -cover ${PACKAGES}

install: build
	go install ${PACKAGES}

clean:
	rm -rf _obj
	go clean -i ${PACKAGES}
