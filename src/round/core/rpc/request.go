// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

// A Request represents a JSON-RPC Request.
type Request struct {
	Version string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  string `json:"params"`
	Id      string `json:"id"`
}

// A Request represents a JSON-RPC Batch Request.
type BatchRequest struct {
	Requests []Request
}

// NewRequest returns a new Request.
func NewRequest() *Request {
	req := &Request{}
	return req
}
