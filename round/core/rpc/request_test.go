// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"errors"
	"testing"
)

// {"jsonrpc": "2.0", "method": "subtract", "params": [42, 23], "id": 1}
func TestNewRequestExamples01(t *testing.T) {
	NewRequest()
}
