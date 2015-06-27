// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package method

import (
	"round"
	"round/core"
	"round/core/rpc"
)

// A Method represents a node method.
type GetRegistry struct {
	*core.Method
}

// NewMethod returns a new Method.
func NewGetRegistry() *GetRegistry {
	method := &GetRegistry{core.NewMethod(round.SystemMethodGetRegistry)}
	return method
}

// Exec runs the specified request on the local node.
func (self *GetRegistry) Exec(node *core.LocalNode, req *rpc.Request) (*rpc.Response, *rpc.Error) {
	var reqReg core.Registry
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
