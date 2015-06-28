###################################################################
#
# Round for Go
#
# Copyright (C) Satoshi Konno 2015
#
# This is licensed under BSD-style license, see file COPYING.
#
###################################################################

packages = round round/core round/core/rpc round/config round/log round/script

	
all: build


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
