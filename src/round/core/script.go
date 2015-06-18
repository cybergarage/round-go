// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A Script represents a script.
type Script struct {
	Language string
	Name string
	Code []byte
	CodeEncoding string
}

// NewScript returns a new Script.
func NewScript() *Script {
	script := &Script{}
	return script
}
