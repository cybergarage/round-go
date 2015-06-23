// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package script

import (
	"errors"
	"fmt"
	"round/core"
	"testing"
)

const (
	luaResultError = "Result (%s) is not equals (%s)"
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
		t.Error(errors.New(fmt.Sprintf(luaResultError, result, param)))
	}
}
