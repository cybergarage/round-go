// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

// A ContentHashObject represents a ContentHashObject.
type ContentHashObject struct {
	*HashObject
}

// NewContentHashObject returns a new ContentHashObject.
func NewContentHashObject() *ContentHashObject {
	cotentObj := &ContentHashObject{}
	cotentObj.HashObject = NewHashObject()
	return cotentObj
}
