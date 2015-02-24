// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A RemoteNode represents a local node.
type RemoteNode struct {
}

// NewRemoteNode returns a new RemoteNode.
func NewRemoteNode() *RemoteNode {
	node := &RemoteNode{}
	return node
}
