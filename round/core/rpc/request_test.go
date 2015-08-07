// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"testing"
)

const (
	errorRequestParserError = "%s != %#v"
)

func TestNewRequestExamples01(t *testing.T) {
	exsamples := []string{
		"{\"jsonrpc\": \"2.0\", \"method\": \"subtract\", \"params\": [42,23], \"id\": 1}",
	}

	expectedMethods := []string{
		"subtract",
	}

	expectedParams := []string{
		"[42,23]",
	}

	for n, example := range exsamples {
		req, err := NewRequestFromString(example)
		if err != nil {
			t.Errorf("%s : %s", err.Error(), example)
			return
		}

		expectedMethod := expectedMethod[n]
		if req.Method != expectedMethod {
			t.Errorf(errorRequestParserError, req.Method, expectedMethod)
		}

		params := req.GetParams()
		expectedParams := expectedParams[n]
		if params != expectedParams {
			t.Errorf(errorRequestParserError, params, expectedParams)
		}
	}
}
