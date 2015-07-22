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
)

const (
	errorNoHashObject = "[%#v] has no HashObjecter interface"
)

// A HashObjecter represents a listener of HashObjecter.
type HashObjecter interface {
	GetHashSeed() string
}

// A HashObject represents a ContentHashObject.
type HashObject struct {
}

// NewHashObject returns a new NewHashObject.
func NewHashObject() *HashObject {
	hashObj := &HashObject{}
	return hashObj
}

// GetHashCode returns a hash code of the specified object.
func (self *HashObject) GetHashCode() string {
	hashSeed := self.GetHashSeed()
	hashByte := sha256.Sum256([]byte(hashSeed))
	var hashCode string
	for _, b := range hashByte {
		hashCode += fmt.Sprintf("%x", b)
	}
	return hashCode
}

// GetHashSeed returns a blnak seed.
func (self *HashObject) GetHashSeed() string {
	return ""
}

// Compare returns a compare result with the specifice object.
func (self *HashObject) Compare(otherObj *HashObject) int {
	selfHashCode := self.GetHashCode()
	otherHashCode := otherObj.GetHashCode()
	if selfHashCode == otherHashCode {
		return 0
	}
	if selfHashCode < otherHashCode {
		return 1
	}
	return -1
}
