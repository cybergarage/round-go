# Copyright 2015 Satoshi Konno. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

all: errors.go

erros.csv:
  
errors.go: erros.csv
	./errors.go.gen > errors.go
	gofmt -w errors.go
