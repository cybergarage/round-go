// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"testing"
)

const (
	testResponseExamples01 = "{\"jsonrpc\": \"2.0\", \"result\": 19, \"id\": 1}"
)

const (
	errorResponseParserError = "%s != %#v"
)

func TestNewResponseExamples01(t *testing.T) {
	res, err := NewResponseFromString(testResponseExamples01)
	if err != nil {
		t.Errorf("%s : %s", err.Error(), testResponseExamples01)
		return
	}

	result := res.GetResult()
	expectedResult := "19"
	if result != expectedResult {
		t.Errorf(errorResponseParserError, result, expectedResult)
	}
}
