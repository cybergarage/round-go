// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A LocalNode represents a local node.
type LocalNode struct {
	*NodeBase
	regMgr    *RegistryManager
	msgMgr    *MessageManager
	scriptMgr *ScriptManager
}

// NewLocalNode returns a new LocalNode.
func NewLocalNode() *LocalNode {
	node := &LocalNode{NodeBase: &NodeBase{}}

	node.regMgr = NewRegistryManager()

	node.msgMgr = NewMessageManager()
	node.msgMgr.SetListener(node)

	node.scriptMgr = NewScriptManager()

	return node
}

// SetRegistry sets a registry by a specified key and value.
func (self *LocalNode) SetRegistry(key, value string) error {
	reg := NewRegistry(key, value)
	return self.regMgr.Set(reg)
}

// GetRegistry gets a registry by the specified key.
func (self *LocalNode) GetRegistry(key string) (*Registry, bool) {
	return self.regMgr.Get(key)
}

// Start starts the local node.
func (self *LocalNode) Start() error {
	err := self.msgMgr.Start()
	if err != nil {
		return err
	}

	return nil
}

// Stop starts the local node.
func (self *LocalNode) Stop() error {
	err := self.msgMgr.Stop()
	if err != nil {
		return err
	}

	return nil
}

func (self *LocalNode) MessageReceived(msg *Message) {
}
