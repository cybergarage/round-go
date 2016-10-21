// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

// A Error represents a error object of JSON-RPC.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewErrorWithMessage returns a new error of the specified code
func NewError(code int) *Error {
	msg := ""
	errMsg, ok := errorMessage[code]
	if ok {
		msg = errMsg
	}
	return NewErrorWithMessage(code, msg)
}

// NewErrorWithMessage returns a new error of the specified code and message
func NewErrorWithMessage(code int, msg string) *Error {
	err := &Error{
		Code:    code,
		Message: msg,
	}
	return err
}

// Error returns a error message.
func (self *Error) Error() string {
	return self.Message
}
