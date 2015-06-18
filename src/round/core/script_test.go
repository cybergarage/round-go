// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"errors"
	"testing"
)

const (
	nullScriptError = "Script is null"
)

func TestNewScript(t *testing.T) {
	script := NewScript()
	if script == nil {
		t.Error(errors.New(nullScriptError))
	}
}
