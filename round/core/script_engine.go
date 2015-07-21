// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A ScriptEngine represents a base script engine.
type ScriptEngine struct {
	Language string
}

// A ScriptEnginer represents a script engine interface.
type ScriptEnginer interface {
	Compile(script *Script) error
	Run(script *Script, params string) (string, error)
}

// NewScriptEngine returns a new ScriptEngine.
func NewScriptEngine() *ScriptEngine {
	engine := &ScriptEngine{}
	return engine
}
