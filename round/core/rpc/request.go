// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"encoding/json"
)

// A Request represents a JSON-RPC Request.
type Request struct {
	Version   string `json:"jsonrpc"`
	Method    string `json:"method"`
	Params    string `json:"params"`
	Id        string `json:"id"`
	Timestamp uint64 `json:"ts"`
}

// A Request represents a JSON-RPC Batch Request.
type BatchRequest struct {
	Requests []Request
}

// NewRequest returns a new Request.
func NewRequest() *Request {
	req := &Request{
		Version: Version,
	}
	return req
}

// GetJSONParams returns a the specified interface using params.
func (self *Request) GetJSONParams(v interface{}) *Error {
	err := json.Unmarshal([]byte(self.Params), v)
	if err != nil {
		return NewError(ErrorCodeInvalidParams)
	}
	return nil
}
