// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"errors"
	"testing"
)

const (
	localNodeNullError = "LocalNode is null"
)

func TestNewLocalNode(t *testing.T) {
	node := NewLocalNode()
	if node == nil {
		t.Error(errors.New(localNodeNullError))
	}

	err := node.Start()
	if err != nil {
		t.Error(err)
	}

	err = node.Stop()
	if err != nil {
		t.Error(err)
	}
}
