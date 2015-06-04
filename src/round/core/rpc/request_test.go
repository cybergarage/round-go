// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"errors"
	"testing"
)

const (
	nullRequestError = "Request is null"
)

func TestNewRequest(t *testing.T) {
	req := NewRequest()
	if req == nil {
		t.Error(errors.New(nullRequestError))
	}
}
