// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"round"
	"round/core/rpc"
)

// A Method represents a node method.
type GetRegistry struct {
	*Method
}

// NewMethod returns a new Method.
func NewGetRegistry() *GetRegistry {
	method := &GetRegistry{NewMethod(round.SystemMethodGetRegistry)}
	return method
}

// Exec runs the specified request on the local node.
func (self *GetRegistry) Exec(node *LocalNode, req *rpc.Request) (*rpc.Response, *rpc.Error) {
	var reqReg Registry
	reqErr := req.GetJSONParams(&reqReg)
	if reqErr != nil {
		return nil, reqErr
	}

	nodeReg, ok := node.GetRegistry(reqReg.Key)
	if !ok {
		return nil, rpc.NewError(rpc.ErrorCodeInternalError)
	}

	return rpc.NewResponseFromRequestWithJSONResult(req, nodeReg), nil
}
