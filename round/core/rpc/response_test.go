// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"testing"
)

const (
	errorResponseParserError = "%s != %#v"
)

func TestNewResponseExamples01(t *testing.T) {
	exsamples := []string{
		"{\"jsonrpc\": \"2.0\", \"result\": 19, \"id\": 1}",
	}

	expectedResults := []string{
		"19",
	}

	for n, example := range exsamples {
		res, err := NewResponseFromString(example)
		if err != nil {
			t.Errorf("%s : %s", err.Error(), example)
			return
		}

		result := res.GetResult()
		expectedResult := expectedResults[n]
		if result != expectedResult {
			t.Errorf(errorResponseParserError, result, expectedResult)
		}
	}
}
