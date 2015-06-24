// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"fmt"
	"testing"
)

const (
	ErrorRegistryManagerSet = "couldn't set (%s, %s)"
	ErrorRegistryManagerGet = "couldn't get (%s) : expected (%s)"
)

func TestNewRegistryManager(t *testing.T) {
	regMgr := NewRegistryManager()

	testLoopCnt := 10
	keyFmt := "key%d"
	valFmt := "val%d"

	for i := 0; i < testLoopCnt; i++ {
		key := fmt.Sprintf(keyFmt, i)
		val := fmt.Sprintf(valFmt, i)
		reg := NewRegistry(key, val)
		err := regMgr.Set(reg)
		if err != nil {
			t.Errorf(ErrorRegistryManagerSet, key, val)
		}
	}

	for i := 0; i < testLoopCnt; i++ {
		key := fmt.Sprintf(keyFmt, i)
		val := fmt.Sprintf(valFmt, i)
		reg, ok := regMgr.Get(key)
		if !ok {
			t.Errorf(ErrorRegistryManagerGet, key, val)
		}
		if reg.Key != key {
			t.Errorf(ErrorRegistryManagerGet, key, val)
		}
		if reg.Value != val {
			t.Errorf(ErrorRegistryManagerGet, key, val)
		}
	}
}
