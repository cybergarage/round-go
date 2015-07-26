// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"fmt"

	"github.com/cybergarage/round-go/round/core/rpc"
)

// A ErrorObject represents a error object for round.
type Error interface {
	error
	GetCode() int
	GetMessage() string
	GetDetailCode() int
	GetDetailMessage() string
}

// A ErrorObject represents a error object for round.
type ErrorObject struct {
	error
	Code          int
	Message       string
	DetailCode    int
	DetailMessage string
}

// NewErrorObject returns a new error.
func NewError() Error {
	err := &ErrorObject{
		Code:          0,
		Message:       "",
		DetailCode:    0,
		DetailMessage: "",
	}
	return err
}

// NewErrorFromRPCError returns a new error from the specified RPC error.
func NewErrorFromRPCError(rpcError *rpc.Error) Error {
	err := &ErrorObject{
		Code:          rpcError.Code,
		Message:       rpcError.Message,
		DetailCode:    0,
		DetailMessage: "",
	}

	return err
}

// GetCode returns a error code.
func (self *ErrorObject) GetCode() int {
	return self.Code
}

// GetMessage returns a error message.
func (self *ErrorObject) GetMessage() string {
	return self.Message
}

// GetDetailCode returns a detail error code.
func (self *ErrorObject) GetDetailCode() int {
	return self.DetailCode
}

// GetDetailMessage returns a detail error message.
func (self *ErrorObject) GetDetailMessage() string {
	return self.DetailMessage
}

// ErrorObject returns a error message.
func (self *ErrorObject) Error() string {
	if self.DetailCode <= 0 {
		return fmt.Sprintf("[%d] %s", self.Code, self.Message)
	}

	return fmt.Sprintf("[%d] %s ([%d] %s)", self.Code, self.Message, self.DetailCode, self.DetailMessage)
}
