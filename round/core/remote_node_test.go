// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"errors"
	"testing"
)

const (
	nullRemoteNodeError = "RemoteNode is null"
)

func TestNewRemoteNode(t *testing.T) {
	req := NewRemoteNode("", 0)
	if req == nil {
		t.Error(errors.New(nullRemoteNodeError))
	}
}
