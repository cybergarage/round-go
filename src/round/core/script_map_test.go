// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"fmt"
	"testing"
)

const (
	errorScriptMapBadObject = "invalid object (%p) : expected (%p)"
	errorScriptMapBadLength = "invalid script len (%d) : expected (%d)"
)

func TestNewScriptMap(t *testing.T) {
	scriptMap := NewScriptMap()

	if len(scriptMap) != 0 {
		t.Errorf(errorScriptMapBadLength, len(scriptMap), 0)
	}

	testLoopCnt := 10
	nameFmt := "name%d"
	scripts := make([]*Script, 0)

	// Add

	for n := 0; n < testLoopCnt; n++ {
		name := fmt.Sprintf(nameFmt, n)
		script := NewScript()
		scriptMap.Set(name, script)
		scripts = append(scripts, script)
	}

	if len(scriptMap) != testLoopCnt {
		t.Errorf(errorScriptMapBadLength, len(scriptMap), testLoopCnt)
	}

	for n := 0; n < testLoopCnt; n++ {
		name := fmt.Sprintf(nameFmt, n)
		script, ok := scriptMap.Get(name)
		if !ok {
			t.Errorf(errorScriptMapBadObject, script, scripts[n])
		}
		if script != scripts[n] {
			t.Errorf(errorScriptMapBadObject, script, scripts[n])
		}
	}

	// Remove

	for n := 0; n < testLoopCnt; n++ {
		name := fmt.Sprintf(nameFmt, n)
		scriptMap.Remove(name)
	}

	if len(scriptMap) != 0 {
		t.Errorf(errorScriptMapBadLength, len(scriptMap), 0)
	}
}
