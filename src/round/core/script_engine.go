// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

type ScriptEngine interface {
	Compile(script *Script) error
	Run(script *Script, params string) (string, error)
}
