// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

// A Error represents a error object of JSON-RPC.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewError returns a new error.
func NewError() *Error {
	err := &Error{}
	return err
}

// Error returns a error message.
func (self *Error) Error() string {
	return self.Message
}
