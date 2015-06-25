// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"fmt"
	"testing"
)

const (
	errorScriptEngineMapBadObject = "invalid object (%p) : expected (%p)"
	errorScriptEngineMapBadLength = "invalid engine len (%d) : expected (%d)"
)

func TestNewScriptEngineMap(t *testing.T) {
	engineMap := NewScriptEngineMap()

	if len(engineMap) != 0 {
		t.Errorf(errorScriptEngineMapBadLength, len(engineMap), 0)
	}

	testLoopCnt := 10
	nameFmt := "name%d"
	engines := make([]*ScriptEngine, 0)

	// Add

	for n := 0; n < testLoopCnt; n++ {
		name := fmt.Sprintf(nameFmt, n)
		engine := NewScriptEngine()
		engineMap.Set(name, engine)
		engines = append(engines, engine)
	}

	if len(engineMap) != testLoopCnt {
		t.Errorf(errorScriptEngineMapBadLength, len(engineMap), testLoopCnt)
	}

	for n := 0; n < testLoopCnt; n++ {
		name := fmt.Sprintf(nameFmt, n)
		engine, ok := engineMap.Get(name)
		if !ok {
			t.Errorf(errorScriptEngineMapBadObject, engine, engines[n])
		}
		if engine != engines[n] {
			t.Errorf(errorScriptEngineMapBadObject, engine, engines[n])
		}
	}

	// Remove

	for n := 0; n < testLoopCnt; n++ {
		name := fmt.Sprintf(nameFmt, n)
		engineMap.Remove(name)
	}

	if len(engineMap) != 0 {
		t.Errorf(errorScriptEngineMapBadLength, len(engineMap), 0)
	}
}
