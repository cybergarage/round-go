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

PACKAGES = ${GITHUB}/${PRODUCT} ${GITHUB}/${PRODUCT}/core ${GITHUB}/${PRODUCT}/common/ ${GITHUB}/${PRODUCT}/core/rpc ${GITHUB}/${PRODUCT}/config ${GITHUB}/${PRODUCT}/log ${GITHUB}/${PRODUCT}/script

.PHONY: ./${PRODUCT}/version.go
	
all: build

./${PRODUCT}/version.go: ./${PRODUCT}/version.gen
	$< > $@
 
version: ./${PRODUCT}/version.go

setup:
	go get -u ${GITHUB}/${PRODUCT}

format:
	gofmt -w src

cgo: 
	go tool cgo ./${PRODUCT}/script/global_script_func.go

build: cgo format
	go build -v ${PACKAGES}

test: build
	go test -v ${PACKAGES}

install: build
	go install ${PACKAGES}

clean:
	rm -rf _obj
	go clean -i ${PACKAGES}
