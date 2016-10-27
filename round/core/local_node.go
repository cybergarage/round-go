// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// #cgo CFLAGS: -I/usr/local/include -DROUND_SUPPORT_JS_SM
// #cgo LDFLAGS: -L/usr/local/lib -L/usr/local/opt/openssl/lib -lround -lmupnp -lcrypto -ljansson -lmozjs185
// #include <round/round.h>
import "C"

import (
	"encoding/json"
	"unsafe"

	"github.com/cybergarage/round-go/round"
	"github.com/cybergarage/round-go/round/core/rpc"
)

// A LocalNode represents a local node.
type LocalNode struct {
	*BaseNode
	msgMgr *MessageManager
	Engine *C.RoundLocalNode
}

// NewLocalNode returns a new LocalNode.
func NewLocalNode() *LocalNode {
	node := &LocalNode{BaseNode: NewBaseNode()}

	node.msgMgr = NewMessageManager()
	node.msgMgr.SetListener(node)

	node.Engine = (*C.RoundLocalNode)(C.round_local_node_new())

	return node
}

func (self *LocalNode) initDefaultMethods() bool {
	return true
}

// SetRegistry sets a registry by a specified registory.
func (self *LocalNode) SetRegistry(reg *Registry) error {
	return nil
}

// GetRegistry gets a registry by the specified key.
func (self *LocalNode) GetRegistry(key string) (*Registry, bool) {
	return nil, false
}

// Start starts the local node.
func (self *LocalNode) Start() error {
	err := self.msgMgr.Start()
	if err != nil {
		return err
	}

	ok, err := C.round_local_node_start((unsafe.Pointer)(self.Engine))
	if !ok {
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
	/*
	  // Set id and ts parameter

	  size_t msgId;
	  if (nodeReq->getId(&msgId)) {
	    nodeRes->setId(msgId);
	  }
	  nodeRes->setTimestamp(getLocalClock());

	  // Exec Message

	  std::string name;
	  if (!nodeReq->getMethod(&name) || (name.length() <= 0)) {
	    setError(RPC::JSON::ErrorCodeMethodNotFound, error);
	    return false;
	  }

	  bool isMethodExecuted = false;
	  bool isMethodSuccess = false;

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

	  if (!isMethodExecuted) {
	    setError(RPC::JSON::ErrorCodeMethodNotFound, error);
	    return false;
	  }

	  if (!isMethodSuccess)
	    return false;

	  if (!hasRoute(name)) {
	    return true;
	  }

	  NodeResponse routeNodeRes;
	  if (!execRoute(name, nodeRes, &routeNodeRes, error)) {
	    return false;
	  }

	  nodeRes->set(&routeNodeRes);
	*/

	return nil, nil
}

// MessageReceived is a listner for MessageManager.
func (self *LocalNode) MessageReceived(msg *Message) *round.Error {
	self.Lock()
	defer self.Unlock()

	var rpcReq rpc.Request
	err := json.Unmarshal([]byte(msg.Content), rpcReq)
	if err != nil {
		return round.NewErrorFromRPCError(rpc.NewError(rpc.ErrorCodeInvalidParams))
	}

	_, rpcErr := self.Exec(&rpcReq)
	if rpcErr != nil {
		return round.NewErrorFromRPCError(rpcErr)
	}

	return nil
}
