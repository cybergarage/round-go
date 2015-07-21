// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package script

import (
	"fmt"
	"testing"

	"github.com/cybergarage/round-go/round/core"
)

const (
	ErrorLuaScriptRun = "Result (%s) is not equals (%s)"
)

func TestNewLuaEngine(t *testing.T) {
	jsEngine := NewLuaEngine()

	script := core.NewScript()
	script.Name = "echo"
	script.Code = []byte("function echo(params)\n\treturn params\nend")
	param := "hello"

	result, err := jsEngine.Run(script, param)
	if err != nil {
		t.Error(err)
	}

	if result != param {
		t.Errorf(ErrorLuaScriptRun, result, param)
	}
}

func TestLuaEngineRegistry(t *testing.T) {
	node := core.NewLocalNode()
	SetLocalNode(node)
	defer SetLocalNode(nil)

	jsEngine := NewLuaEngine()

	script := core.NewScript()
	script.Name = "test_registry"
	script.Code = []byte("function test_registry(params)\n  set_registry(params, params)\n  ok, value = get_registry(params)\n  return value\nend")

	testLoopCnt := 10
	paramFmt := "param%d"
	for i := 0; i < testLoopCnt; i++ {
		param := fmt.Sprintf(paramFmt, i)

		result, err := jsEngine.Run(script, param)
		if err != nil {
			t.Error(err)
		}

		if result != param {
			t.Errorf(ErrorLuaScriptRun, result, param)
		}
	}
}
