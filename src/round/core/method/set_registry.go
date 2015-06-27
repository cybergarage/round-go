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
type SetRegistry struct {
	*core.Method
}

// NewMethod returns a new Method.
func NewSetRegistry() *SetRegistry {
	method := &SetRegistry{core.NewMethod(round.SystemMethodSetRegistry)}
	return method
}

// Exec runs the specified request on the local node.
func (self *SetRegistry) Exec(node *core.LocalNode, req *rpc.Request) (*rpc.Response, *rpc.Error) {
	var reg core.Registry
	reqErr := req.GetJSONParams(&reg)
	if reqErr != nil {
		return reqErr
	}

	nodeErr := node.SetRegistry(&reg)
	if nodeErr != nil {
		return nil, NewError(ErrorCodeInternalError)
	}

	return nil, nil
}
