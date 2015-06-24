// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A LocalNode represents a local node.
type LocalNode struct {
	*NodeBase
	registryMgr *RegistryManager
}

// NewLocalNode returns a new LocalNode.
func NewLocalNode() *LocalNode {
	node := &LocalNode{NodeBase: &NodeBase{}}
	node.registryMgr = NewRegistryManager()
	return node
}

// SetRegistry sets a registry by a specified key and value.
func (self *LocalNode) SetRegistry(key, value string) error {
	reg := NewRegistry(key, value)
	return self.registryMgr.Set(reg)
}

// GetRegistry gets a registry by the specified key.
func (self *LocalNode) GetRegistry(key string) (*Registry, bool) {
	return self.registryMgr.Get(key)
}
