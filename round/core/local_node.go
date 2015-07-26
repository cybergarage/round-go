// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"encoding/json"

	"github.com/cybergarage/round-go/round/core/rpc"
)

// A LocalNode represents a local node.
type LocalNode struct {
	*BaseNode
	cluster   *Cluster
	regMgr    *RegistryManager
	msgMgr    *MessageManager
	methodMgr *MethodManager
	scriptMgr *ScriptManager
}

// NewLocalNode returns a new LocalNode.
func NewLocalNode() *LocalNode {
	node := &LocalNode{
		BaseNode:  NewBaseNode(),
		cluster:   NewCluster(),
		regMgr:    NewRegistryManager(),
		msgMgr:    NewMessageManager(),
		scriptMgr: NewScriptManager(),
		methodMgr: NewMethodManager(),
	}

	node.init()

	return node
}

func (self *LocalNode) init() bool {
	self.cluster.AddNode(self)
	self.msgMgr.SetListener(self)

	return true
}

// SetStaticMethod adds a method as a static method.
func (self *LocalNode) SetStaticMethod(method Method) bool {
	return self.methodMgr.SetStaticMethod(method)
}

// SetStaticMethods adds a method as a static method.
func (self *LocalNode) SetStaticMethods(method []Method) bool {
	return self.methodMgr.SetStaticMethods(method)
}

// SetDynamicMethod adds a method as a dynamic method.
func (self *LocalNode) SetDynamicMethod(method Method) bool {
	return self.methodMgr.SetDynamicMethod(method)
}

// SetDynamicMethods adds a method as a dynamic method.
func (self *LocalNode) SetDynamicMethods(method []Method) bool {
	return self.methodMgr.SetDynamicMethods(method)
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

// Exec runs the specified request.
func (self *LocalNode) Exec(req *rpc.Request) (*rpc.Response, *rpc.Error) {
	res := rpc.NewResponse()

	// Set id and ts parameter

	res.Id = req.Id
	res.Timestamp = self.GetClock()

	// Exec Message

	method := req.Method
	if len(method) <= 0 {
		err := rpc.NewError(rpc.ErrorCodeMethodNotFound)
		res.SetErrorResult(err)
		return res, err
	}

	/*
		res, rpcErr := self.methodMgr.ExecMethod(self, method, req) (*rpc.Response, error) {
	*/
	isMethodExecuted := false
	switch {
	case self.methodMgr.IsStaticMethod(method):
		isMethodExecuted = true
	case self.methodMgr.IsDynamicMethod(method):
		isMethodExecuted = true
	}
	/*
	  // Exec Message


	  if (isStaticMethod(name)) {
	    isMethodExecuted = true;
	    isMethodSuccess = execStaticMethod(nodeReq, nodeRes, error);
	  }
	  else if (isDynamicMethod(name)) {
	    isMethodExecuted = true;
	    isMethodSuccess = execDynamicMethod(nodeReq, nodeRes, error);
	  }
	  else if (isNativeMethod(name)) {
	    isMethodExecuted = true;
	    isMethodSuccess = execNativeMethod(nodeReq, nodeRes, error);
	  }
	  else if (isAliasMethod(name)) {
	    isMethodExecuted = true;
	    isMethodSuccess = execAliasMethod(nodeReq, nodeRes, error);
	  }
	*/
	if !isMethodExecuted {
		err := rpc.NewError(rpc.ErrorCodeMethodNotFound)
		res.SetErrorResult(err)
		return res, err
	}

	/*
		if !isMethodSuccess {
			return res, err
		}

		  if (!hasRoute(name)) {
		    return true;
		  }

		  NodeResponse routeNodeRes;
		  if (!execRoute(name, nodeRes, &routeNodeRes, error)) {
		    return false;
		  }

		  nodeRes->set(&routeNodeRes);
	*/

	return res, nil
}

// MessageReceived is a listner for MessageManager.
func (self *LocalNode) MessageReceived(msg *Message) Error {
	self.Lock()
	defer self.Unlock()

	var rpcReq rpc.Request
	err := json.Unmarshal([]byte(msg.Content), rpcReq)
	if err != nil {
		return NewErrorFromRPCError(rpc.NewError(rpc.ErrorCodeInvalidParams))
	}

	_, rpcErr := self.Exec(&rpcReq)
	if rpcErr != nil {
		return NewErrorFromRPCError(rpcErr)
	}

	return nil
}
