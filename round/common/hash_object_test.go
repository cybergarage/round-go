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
	errorHashObjectCompareInvalid          = "(%s, %s) = %d: expected %d"
)

func TestNewHashObject(t *testing.T) {
	NewHashBaseObject()
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

func TestHashObjectCompare(t *testing.T) {
	// number strings

	codes := []string{"0", "0"}
	cmp := HashCodeCompare(codes[0], codes[1])
	if cmp != HashCodeCompareEqual {
		t.Errorf(errorHashObjectCompareInvalid, codes[0], codes[1], cmp, HashCodeCompareEqual)
	}

	codes = []string{"0", "1"}
	cmp = HashCodeCompare(codes[0], codes[1])
	if cmp != HashCodeCompareGreater {
		t.Errorf(errorHashObjectCompareInvalid, codes[0], codes[1], cmp, HashCodeCompareGreater)
	}

	// number, alpha string

	codes = []string{"0", "a"}
	cmp = HashCodeCompare(codes[0], codes[1])
	if cmp != HashCodeCompareGreater {
		t.Errorf(errorHashObjectCompareInvalid, codes[0], codes[1], cmp, HashCodeCompareGreater)
	}

	codes = []string{"a", "0"}
	cmp = HashCodeCompare(codes[0], codes[1])
	if cmp != HashCodeCompareLess {
		t.Errorf(errorHashObjectCompareInvalid, codes[0], codes[1], cmp, HashCodeCompareLess)
	}

	// number, alpha strings

	codes = []string{"0000000000", "ffffffffff"}
	cmp = HashCodeCompare(codes[0], codes[1])
	if cmp != HashCodeCompareGreater {
		t.Errorf(errorHashObjectCompareInvalid, codes[0], codes[1], cmp, HashCodeCompareGreater)
	}

	codes = []string{"ffffffffff", "0000000000"}
	cmp = HashCodeCompare(codes[0], codes[1])
	if cmp != HashCodeCompareLess {
		t.Errorf(errorHashObjectCompareInvalid, codes[0], codes[1], cmp, HashCodeCompareLess)
	}
}
