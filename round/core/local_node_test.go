// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"errors"
	"testing"
)

const (
	nullLocalNodeError = "LocalNode is null"
)

func TestNewLocalNode(t *testing.T) {
	req := NewLocalNode()
	if req == nil {
		t.Error(errors.New(nullLocalNodeError))
	}
}
