// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A LocalNode represents a local node.
type LocalNode struct {
	*NodeBase
}

// NewLocalNode returns a new LocalNode.
func NewLocalNode() *LocalNode {
	node := &LocalNode{NodeBase: &NodeBase{}}
	return node
}
