// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"encoding/json"
)

// A Response represents a JSON-RPC reqponse.
type Response struct {
	Version string `json:"jsonrpc"`
	Result  string `json:"result"`
	Id      string `json:"id"`
}

// NewResponse returns a new Response.
func NewResponse() *Response {
	res := &Response{Version: Version}
	return res
}

// NewResponseFromRequest returns a new Response which has a same id with the specified request.
func NewResponseFromRequest(req *Request) *Response {
	res := NewResponse()
	res.Id = req.Id
	return res
}

// SetJSONParams set the specified struct into the result.
func (self *Response) SetJSONParams(v interface{}) *Error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return NewError(ErrorCodeInvalidParams)
	}
	self.Result = string(bytes)
	return nil
}
