// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import "encoding/json"

// A base node  represents a JSON-RPC Request.
type Node interface {
	PostMessage(reqMsg string) (resMsg string, err error);
}

// A NodeBase represents a base node.
type NodeBase struct {
}

func (self *NodeBase) PostJsonMessage(jsonReq interface{}, jsonRes interface{}) error {
	_, err := json.Marshal(jsonReq)
	if err != nil {
		return err
		}
	//resMsg, err := PostMessage("")
	return  nil
}