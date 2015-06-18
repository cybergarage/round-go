// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package script

import (
	"round/core"
)

type JavaScriptEngine struct {
}

// Compile compiles the specified script.
func (self *JavaScriptEngine) Compile(script *Script) error {
	return nil
}

// Run runs the specified script.
func (self *JavaScriptEngine) Run(script *Script, params string) (string, error) {
	return "", nil
}
	
	
