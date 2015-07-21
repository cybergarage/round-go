// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

import (
	"testing"
)

const (
	errorHashObjectInvalidLength = "HashObject (#v) is invalid length = %d: expected %d"
)

func TestNewHashObject(t *testing.T) {
	NewHashObject()
}

func TestNewTestHashObject(t *testing.T) {
	hashObj := NewTestHashObject()

	hashCode, err := hashObj.GetHashCode()
	if err != nil {
		t.Error(err)
	}

	if len(hashCode) == GetHashCodeLength() {
		t.Errorf(errorHashObjectInvalidLength, hashObj, len(hashCode), GetHashCodeLength())
	}
}
