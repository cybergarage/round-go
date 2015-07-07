###################################################################
#
# Round for Go
#
# Copyright (C) Satoshi Konno 2015
#
# This is licensed under BSD-style license, see file COPYING.
#
###################################################################

packages = round round/core round/common/ round/core/rpc round/config round/log round/script

.PHONY: ./src/round/version.go
	
all: build

./src/round/version.go: ./src/round/version.gen
	$< > $@
 
version: ./src/round/version.go

format:
	gofmt -w src

cgo: format
	go tool cgo src/round/script/global_script_func.go

build: cgo
	go build -v ${packages}

test: build
	go test -v ${packages}

install: build
	go install ${packages}

clean:
	rm -rf _obj
	go clean -i ${packages}
