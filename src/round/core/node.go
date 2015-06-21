// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import "encoding/json"

// A base node  represents a JSON-RPC Request.
type Node interface {
	PostMessage(reqMsg string) (resMsg string, err error)
}

// A NodeBase represents a base node.
type NodeBase struct {
}

// PostMessage is dummy. The method have to be overrided.
func (self *NodeBase) PostMessage(reqMsg string) (resMsg string, err error) {
	return "", nil
}

// PostJsonMessage posts the specified json request.
func (self *NodeBase) PostJsonMessage(jsonReq interface{}) (jsonRes interface{}, err error) {
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
