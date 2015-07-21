// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A ScriptManager represents a registry manager.
type ScriptManager struct {
	engines ScriptEngineMap
	scripts ScriptMap
}

// NewScriptManager returns a new ScriptManager.
func NewScriptManager() *ScriptManager {
	scriptMgr := &ScriptManager{}
	scriptMgr.engines = NewScriptEngineMap()
	scriptMgr.scripts = NewScriptMap()
	return scriptMgr
}

// SetScript sets a script.
func (self *ScriptManager) SetScript(script *Script) error {
	return self.scripts.Set(script.Name, script)
}

// GetScript gets a script by the specified name.
func (self *ScriptManager) GetScript(name string) (*Script, bool) {
	return self.scripts.Get(name)
}

// RemoveScript removes a script by the specified name.
func (self *ScriptManager) RemoveScript(name string) bool {
	return self.scripts.Remove(name)
}

// SetEngine sets a script engine.
func (self *ScriptManager) SetEngine(engine *ScriptEngine) error {
	return self.engines.Set(engine.Language, engine)
}

// GetEngine gets a script engine by the specified name.
func (self *ScriptManager) GetEngine(lang string) (*ScriptEngine, bool) {
	return self.engines.Get(lang)
}

// RemoveEngine removes a script engine by the specified name.
func (self *ScriptManager) RemoveEngine(lang string) bool {
	return self.engines.Remove(lang)
}
