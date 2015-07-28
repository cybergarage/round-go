// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"testing"
)

const (
	testRequestExamples01 = "{\"jsonrpc\": \"2.0\", \"method\": \"subtract\", \"params\": [42,23], \"id\": 1}"
)

const (
	errorRequestParserError = "%s != %#v"
)

func TestNewRequestExamples01(t *testing.T) {
	req, err := NewRequestFromString(testRequestExamples01)
	if err != nil {
		t.Errorf("%s : %s", err.Error(), testRequestExamples01)
		return
	}

	if req.Method != "subtract" {
		t.Errorf(errorRequestParserError, TestNewRequestExamples01, req)
	}

	params := req.GetParams()
	expectedParams := "[42,23]"
	if params != expectedParams {
		t.Errorf(errorRequestParserError, params, expectedParams)
	}
}
