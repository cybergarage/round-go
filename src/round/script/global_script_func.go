// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package script

import "C"

import (
	"round/core"
)

var currentScriptLocalNode *core.LocalNode = nil

func SetLocalNode(node *core.LocalNode)  {
	currentScriptLocalNode = node
}

func GetLocalNode() *core.LocalNode {
	return currentScriptLocalNode
}

//export GetLocalNodeRegistry
func GetLocalNodeRegistry(key *C.char) *C.char {
	node := GetLocalNode()
	if node == nil {
		return C.CString("")
	}
	
	return C.CString("hi")
}

//export SetLocalNodeRegistry
func SetLocalNodeRegistry(key, value *C.char) C.int {
	//node := GetLocalNode()	
	return C.int(1)
}