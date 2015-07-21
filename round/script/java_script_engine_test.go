// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package script

import (
	"testing"

	"github.com/cybergarage/round-go/round/core"
)

func TestNewJavaScriptEngine(t *testing.T) {
	jsEngine := NewJavaScriptEngine()

	script := core.NewScript()
	script.Code = []byte("var result = 1 + 2 + 3;\n result;")
	_, err := jsEngine.Run(script, "")
	if err != nil {
		t.Error(err)
	}
}
