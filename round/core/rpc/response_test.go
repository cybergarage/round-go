// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"errors"
	"testing"
)

const (
	nullResponseError = "Response is null"
)

func TestNewResponse(t *testing.T) {
	res := NewResponse()
	if res == nil {
		t.Error(errors.New(nullResponseError))
	}
}
