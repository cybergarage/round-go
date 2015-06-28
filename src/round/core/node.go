// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"encoding/json"
	"sync"
)

// A base node  represents a JSON-RPC Request.
type Node interface {
	PostMessage(reqMsg string) (resMsg string, err error)
}

// A NodeBase represents a base node.
type BaseNode struct {
	mutex *sync.Mutex
}

// NewBaseNode returns a new BaseNode.
func NewBaseNode() *BaseNode {
	node := &BaseNode{}
	node.mutex = &sync.Mutex{}
	return node
}

// Lock locks a mutex of the node.
func (self *BaseNode) Lock() {
	self.mutex.Lock()
}

// Unlock unlocks a mutex of the node.
func (self *BaseNode) Unlock() {
	self.mutex.Unlock()
}

// PostMessage is dummy. The method have to be overrided.
func (self *BaseNode) PostMessage(reqMsg string) (resMsg string, err error) {
	return "", nil
}

// PostJsonMessage posts the specified json request.
func (self *BaseNode) PostJsonMessage(jsonReq interface{}) (jsonRes interface{}, err error) {
	reqMsg, err := json.Marshal(jsonReq)
	if err != nil {
		return nil, err
	}

	_, err = self.PostMessage(string(reqMsg))
	if err != nil {
		return nil, err
	}

	/*
		    jsonRes interface{}	= nil
			err := json.Unmarshal([]byte(resMsg), jsonRes)
			if err != nil {
				return err
			}
	*/

	return nil, nil
}
