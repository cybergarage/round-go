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
PRODUCT_DIR=./${PRODUCT}
GITHUB=github.com/cybergarage/${PRODUCT}-go

PACKAGES = ${GITHUB}/${PRODUCT} ${GITHUB}/${PRODUCT}/core ${GITHUB}/${PRODUCT}/common/ ${GITHUB}/${PRODUCT}/core/rpc ${GITHUB}/${PRODUCT}/config ${GITHUB}/${PRODUCT}/log ${GITHUB}/${PRODUCT}/script ${GITHUB}/${PRODUCT}/impl
CONST_FILES = ${PRODUCT_DIR}/version.go ${PRODUCT_DIR}/const.go ${PRODUCT_DIR}/impl/const.go ${PRODUCT_DIR}/impl/server_desc.go

.PHONY: ${CONST_FILES}

all: build

const: ${CONST_FILES}

${PRODUCT_DIR}/const.go: ${PRODUCT_DIR}/const.go.gen
	$< > $@
	gofmt -w $@
 
${PRODUCT_DIR}/version.go: ${PRODUCT_DIR}/version.go.gen
	$< > $@
	gofmt -w $@

${PRODUCT_DIR}/impl/const.go: ${PRODUCT_DIR}/impl/const.go.gen
	$< ${PRODUCT_DIR}/const.go > $@
	gofmt -w $@

${PRODUCT_DIR}/impl/upnp_server_desc.go: ${PRODUCT_DIR}/impl/upnp_server_desc.go.gen
	$< > $@
	gofmt -w $@

setup:
	go get -u github.com/cybergarage/go-net-upnp/net/upnp
	go get -u ${GITHUB}/${PRODUCT}

format:
	gofmt -w src

cgo: 
	go tool cgo ${PRODUCT_DIR}/script/global_script_func.go

build: cgo format
	go build -v ${PACKAGES}

test: build
	go test -v -cover ${PACKAGES}

install: build
	go install ${PACKAGES}

clean:
	rm -rf _obj
	go clean -i ${PACKAGES}
