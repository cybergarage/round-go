// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

// A Response represents a JSON-RPC reqponse.
type Response struct {
	Version string `json:"jsonrpc"`
	Result  string `json:"result"`
	Id      string `json:"id"`
}

// NewResponse returns a new Response.
func NewResponse() *Response {
	res := &Response{}
	return res
}
