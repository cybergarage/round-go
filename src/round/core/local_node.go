// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import ()

// A LocalNode represents a local node.
type LocalNode struct {
	*NodeBase
	regMgr    *RegistryManager
	msgMgr    *MessageManager
	methodMgr *MethodManager
	scriptMgr *ScriptManager
}

// NewLocalNode returns a new LocalNode.
func NewLocalNode() *LocalNode {
	node := &LocalNode{NodeBase: &NodeBase{}}

	node.regMgr = NewRegistryManager()
	node.scriptMgr = NewScriptManager()

	node.methodMgr = NewMethodManager()
	node.initDefaultMethods()

	node.msgMgr = NewMessageManager()
	node.msgMgr.SetListener(node)

	return node
}

func (self *LocalNode) initDefaultMethods() bool {
	self.methodMgr.SetDynamicMethod(NewGetRegistryMethod())

	return true
}

// SetRegistry sets a registry by a specified registory.
func (self *LocalNode) SetRegistry(reg *Registry) error {
	newReg := NewRegistryFromRegistry(reg)
	return self.regMgr.Set(newReg)
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
