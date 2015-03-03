# Copyright 2015 Satoshi Konno. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

all:	constants.go

constants.csv:
	wget https://raw.githubusercontent.com/cybergarage/round-common/master/src/constants.csv constants.csv

constants.go:	constants.csv	
	./constants.build > constants.go
