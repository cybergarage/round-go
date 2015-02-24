// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A Request represents a JSON-RPC Request.
type Request struct {
}

// NewRequest returns a new Request.
func NewRequest() *Request {
	req := &Request{}
	return req
}
