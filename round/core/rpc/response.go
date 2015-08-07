// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"encoding/json"
)

// A Response represents a JSON-RPC reqponse.
type Response struct {
	Version   string `json:"jsonrpc"`
	Result    string `json:"result"`
	Id        uint64 `json:"id"`
	Timestamp uint64 `json:"ts"`
}

// NewResponse returns a new Response.
func NewResponse() *Response {
	res := &Response{Version: Version}
	return res
}

// NewResponseFromBytes returns a new Response from the specified bytes.
func NewResponseFromBytes(resBytes []byte) (*Response, *Error) {
	res := &Response{}
	err := res.ParseBytes(resBytes)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// NewResponseFromString returns a new Response from the specified string.
func NewResponseFromString(reqStr string) (*Response, *Error) {
	return NewResponseFromBytes([]byte(reqStr))
}

// NewResponseFromRequest returns a new Response which has a same id with the specified request.
func NewResponseFromRequest(req *Request) *Response {
	res := NewResponse()
	res.Id = req.Id
	return res
}

// NewResponseFromRequestWithResult returns a new Response which has a same id with the specified request and result.
func NewResponseFromRequestWithResult(req *Request, result string) *Response {
	res := NewResponseFromRequest(req)
	res.Result = result
	return res
}

// NewResponseFromRequestWithJSONResult returns a new Response which has a same id with the specified request and result.
func NewResponseFromRequestWithJSONResult(req *Request, result interface{}) *Response {
	res := NewResponseFromRequest(req)
	res.SetJSONResult(result)
	return res
}

// ParseBytes parses the specified bytes.
func (self *Response) ParseBytes(resBytes []byte) *Error {
	err := json.Unmarshal(resBytes, self)
	if err != nil {
		return NewError(ErrorCodeParserError)
	}

	return nil
}

// SetResult sets a object into the result.
func (self *Response) SetResult(v interface{}) *Error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return NewError(ErrorCodeInvalidParams)
	}
	self.Result = string(bytes)
	return nil
}

// SetJSONResult sets the specified struct into the result.
func (self *Response) SetJSONResult(v interface{}) *Error {
	return self.SetResult(v)
}

// SetErrorResult set an error into the result.
func (self *Response) SetErrorResult(rpcErr *Error) *Error {
	return self.SetResult(rpcErr)
}

// GetResult returns a the specified interface using params.
func (self *Response) GetResult() string {
	b, err := json.Marshal(self.Result)
	if err != nil {
		return ""
	}
	return string(b)
}

// GetJSONResult returns a the specified interface using params.
func (self *Response) GetJSONResult(v interface{}) *Error {
	/*
		err := json.Unmarshal([]byte(self.Result), v)
		if err != nil {
			return NewError(ErrorCodeInvalidResult)
		}
	*/
	return nil
}
