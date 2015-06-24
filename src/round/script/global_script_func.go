// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package script

import "C"

import (
	"round/core"
)

var currentScriptLocalNode *core.LocalNode = nil

func SetLocalNode(node *core.LocalNode) {
	currentScriptLocalNode = node
}

func GetLocalNode() *core.LocalNode {
	return currentScriptLocalNode
}

//export LocalNodeGetRegistry
func LocalNodeGetRegistry(key *C.char) *C.char {
	node := GetLocalNode()
	if node == nil {
		return C.CString("")
	}
	reg, ok := node.GetRegistry(C.GoString(key))
	if !ok {
		return C.CString("")
	}
	return C.CString(reg.Value)
}

//export LocalNodeSetRegistry
func LocalNodeSetRegistry(key, value *C.char) C.int {
	node := GetLocalNode()
	if node == nil {
		return C.int(0)
	}
	err := node.SetRegistry(C.GoString(key), C.GoString(value))
	if err != nil {
		return C.int(0)
	}
	return C.int(1)
}
