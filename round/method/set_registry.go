// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package method

import (
	"github.com/cybergarage/round-go/round/core"
	"github.com/cybergarage/round-go/round/core/rpc"
)

// A Method represents a node method.
type SetRegistry struct {
	*core.BaseMethod
}

// NewMethod returns a new Method.
func NewSetRegistryMethod() *SetRegistry {
	method := &SetRegistry{
		BaseMethod: core.NewBaseMethod(SystemMethodSetRegistry),
	}
	return method
}

// Exec runs the specified request on the local node.
func (self *SetRegistry) Exec(node *core.LocalNode, req *rpc.Request) (*rpc.Response, *rpc.Error) {
	var reg core.Registry
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
