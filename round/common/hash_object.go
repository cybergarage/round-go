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
func (self *HashObject) GetHashCode() (string, error) {
	hashSeed := self.GetHashSeed()
	hashByte := sha256.Sum256([]byte(hashSeed))
	var hashCode string
	for _, b := range hashByte {
		hashCode += fmt.Sprintf("%x", b)
	}
	return hashCode, nil
}

// GetHashSeed returns a blnak seed.
func (self *HashObject) GetHashSeed() string {
	return ""
}
