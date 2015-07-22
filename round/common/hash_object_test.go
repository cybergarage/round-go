// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

import (
	"testing"
)

const (
	errorHashObjectGetHashSeedNotOverrided = "GetHashSeed() is not overrided"
	errorHashObjectInvalidLength           = "HashObject (%#v) is invalid length = %d: expected %d"
)

func TestNewHashObject(t *testing.T) {
	NewHashObject()
}

func TestNewTestHashObject(t *testing.T) {
	hashObj := NewTestHashObject()

	hashSeed := hashObj.GetHashSeed()
	if len(hashSeed) <= 0 {
		t.Errorf(errorHashObjectGetHashSeedNotOverrided)
	}

	hashCode := hashObj.GetHashCode()

	if len(hashCode) != HashCodeSize {
		t.Errorf(errorHashObjectInvalidLength, hashObj, len(hashCode), HashCodeSize)
	}
}
