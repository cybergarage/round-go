// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

import (
	"crypto/sha256"
	"fmt"
)

const (
	HashCodeSize = (sha256.Size * 2)

	HashCodeCompareEqual   = 0
	HashCodeCompareGreater = 1
	HashCodeCompareLess    = -1
)

const (
	errorNoHashObject = "[%#v] has no HashObjecter interface"
)

// HashCodeCompare returns a compare result with the specifice object.
func HashCodeCompare(selfHashCode, otherHashCode string) int {
	if selfHashCode == otherHashCode {
		return HashCodeCompareEqual
	}
	if selfHashCode < otherHashCode {
		return HashCodeCompareGreater
	}
	return HashCodeCompareLess
}

// A HashObjecter represents a listener of HashObjecter.
type HashObject interface {
	GetHashSeed() string
	GetHashCode() string
}

// A HashObject represents a ContentHashObject.
type HashBaseObject struct {
}

// NewHashObject returns a new NewHashObject.
func NewHashBaseObject() *HashBaseObject {
	hashObj := &HashBaseObject{}
	return hashObj
}

// GetHashCode returns a hash code of the specified object.
func (self *HashBaseObject) GetHashCode() string {
	hashSeed := self.GetHashSeed()
	hashByte := sha256.Sum256([]byte(hashSeed))
	var hashCode string
	for _, b := range hashByte {
		hashCode += fmt.Sprintf("%x", b)
	}
	return hashCode
}

// GetHashSeed returns a blnak seed.
func (self *HashBaseObject) GetHashSeed() string {
	return ""
}

// Compare returns a compare result with the specifice object.
func (self *HashBaseObject) Compare(otherObj HashObject) int {
	return HashCodeCompare(self.GetHashCode(), otherObj.GetHashCode())
}
