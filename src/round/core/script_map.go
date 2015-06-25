// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A ScriptMap represents a script map.
type ScriptMap map[string]*Script

// NewScriptMap returns a new ScriptMap.
func NewScriptMap() ScriptMap {
	scriptMap := ScriptMap{}
	return scriptMap
}

// Set sets a script with the specified name.
func (self ScriptMap) Set(name string, script *Script) error {
	self[name] = script
	return nil
}

// Get gets a script by the specified name.
func (self ScriptMap) Get(name string) (*Script, bool) {
	script, ok := self[name]
	return script, ok
}

// Remove removes a script by the specified name.
func (self ScriptMap) Remove(name string) bool {
	delete(self, name)
	return true
}
