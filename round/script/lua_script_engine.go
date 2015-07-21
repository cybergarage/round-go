// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package script

// #cgo CFLAGS: -I/usr/local/include -DROUND_SUPPORT_LUA -I../../../_obj
// #cgo LDFLAGS: -L/usr/local/lib -lround -llua
// #include <roundc/round.h>
// int round_lua_setregistry(lua_State* L);
// int round_lua_getregistry(lua_State* L);
import "C"

import (
	//"fmt"
	"errors"

	"github.com/cybergarage/round-go/round"
	"github.com/cybergarage/round-go/round/core"
)

type LuaEngine struct {
	*core.ScriptEngine
	Engine *C.RoundLuaEngine
}

// NewLocalNode returns a new LocalNode.
func NewLuaEngine() *LuaEngine {
	lueEngine := &LuaEngine{}
	lueEngine.ScriptEngine = core.NewScriptEngine()
	lueEngine.Engine = C.round_lua_engine_new()

	C.round_lua_engine_register(lueEngine.Engine, C.CString(round.SystemMethodSetRegistry), C.lua_CFunction(C.round_lua_setregistry))
	C.round_lua_engine_register(lueEngine.Engine, C.CString(round.SystemMethodGetRegistry), C.lua_CFunction(C.round_lua_getregistry))

	return lueEngine
}

// Compile compiles the specified script.
func (self *LuaEngine) Compile(script *core.Script) error {
	return nil
}

// Run runs the specified script.
func (self *LuaEngine) Run(script *core.Script, params string) (string, error) {
	var result, errmsg string

	name := script.Name
	code := string(script.Code)

	C.round_lua_engine_lock(self.Engine)

	ok := bool(C.round_lua_engine_run(self.Engine, C.CString(code), C.CString(name), C.CString(params)))
	if ok {
		result = C.GoString(C.round_lua_engine_getresult(self.Engine))
	} else {
		errmsg = C.GoString(C.round_lua_engine_geterror(self.Engine))
	}

	C.round_lua_engine_unlock(self.Engine)

	//fmt.Printf("%t %s %s %s\n", ok, code, result, errmsg)

	if !ok {
		return "", errors.New(errmsg)
	}

	return result, nil
}
