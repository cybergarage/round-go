# Copyright 2015 Satoshi Konno. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

all: const.go

const.csv:
  
const.go: const.csv
	./const.go.gen > const.go
	gofmt -w const.go
