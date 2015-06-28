// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"round"
	"round/core/rpc"
)

// A Method represents a node method.
type SetRegistry struct {
	*Method
}

// NewMethod returns a new Method.
func NewSetRegistry() *SetRegistry {
	method := &SetRegistry{NewMethod(round.SystemMethodSetRegistry)}
	return method
}

// Exec runs the specified request on the local node.
func (self *SetRegistry) Exec(node *LocalNode, req *rpc.Request) (*rpc.Response, *rpc.Error) {
	var reg Registry
	reqErr := req.GetJSONParams(&reg)
	if reqErr != nil {
		return nil, reqErr
	}

	nodeErr := node.SetRegistry(&reg)
	if nodeErr != nil {
		return nil, rpc.NewError(rpc.ErrorCodeInternalError)
	}

	return rpc.NewResponseFromRequest(req), nil
}
