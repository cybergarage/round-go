// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package script

// #cgo CFLAGS: -I/usr/local/include -DROUND_SUPPORT_JS_SM
// #cgo LDFLAGS: -L/usr/local/lib -lround -lmozjs185
// #include <roundc/round.h>
import "C"

import (
	//"fmt"
	"errors"
	"round/core"
)

type JavaScriptEngine struct {
	*core.ScriptEngine
	Engine *C.RoundJavaScriptEngine
}

// NewLocalNode returns a new LocalNode.
func NewJavaScriptEngine() *JavaScriptEngine {
	jsEngine := &JavaScriptEngine{}
	jsEngine.ScriptEngine = core.NewScriptEngine()
	jsEngine.Engine = C.round_js_engine_new()
	return jsEngine
}

// Compile compiles the specified script.
func (self *JavaScriptEngine) Compile(script *core.Script) error {
	return nil
}

// Run runs the specified script.
func (self *JavaScriptEngine) Run(script *core.Script, params string) (string, error) {
	var result, errmsg string

	code := string(script.Code)

	C.round_js_engine_lock(self.Engine)

	ok := bool(C.round_js_engine_run(self.Engine, C.CString(code)))
	if ok {
		result = C.GoString(C.round_js_engine_getresult(self.Engine))
	} else {
		errmsg = C.GoString(C.round_js_engine_geterror(self.Engine))
	}

	C.round_js_engine_unlock(self.Engine)

	//fmt.Printf("%t %s %s %s\n", ok, code, result, errmsg)

	if !ok {
		return "", errors.New(errmsg)
	}

	return result, nil
}
