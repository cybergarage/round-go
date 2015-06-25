// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A ScriptEngineMap represents a script map.
type ScriptEngineMap map[string]*ScriptEngine

// NewScriptEngineMap returns a new ScriptEngineMap.
func NewScriptEngineMap() ScriptEngineMap {
	scriptMap := ScriptEngineMap{}
	return scriptMap
}

// Set sets a script engine with the specified language.
func (self ScriptEngineMap) Set(lang string, engine *ScriptEngine) error {
	self[lang] = engine
	return nil
}

// Get gets a script engine by the specified language.
func (self ScriptEngineMap) Get(lang string) (*ScriptEngine, bool) {
	script, ok := self[lang]
	return script, ok
}

// Remove removes a script engine by the specified language.
func (self ScriptEngineMap) Remove(lang string) bool {
	delete(self, lang)
	return true
}
