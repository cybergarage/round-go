// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A Response represents a JSON-RPC reqponse.
type Response struct {
}

// NewResponse returns a new Response.
func NewResponse() *Response {
	res := &Response{}
	return res
}
