// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A RemoteNode represents a local node.
type RemoteNode struct {
	*BaseNode
	Address string
	Port    int
}

// NewRemoteNode returns a new RemoteNode.
func NewRemoteNode(addr string, port int) *RemoteNode {
	node := &RemoteNode{BaseNode: NewBaseNode(), Address: addr, Port: port}
	return node
}
