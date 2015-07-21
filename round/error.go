// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package round

import (
	"round/core/rpc"
)

// A Error represents a error object for round.
type Error struct {
	Code          int
	Message       string
	DetailCode    int
	DetailMessage string
}

// NewError returns a new error.
func NewError() *Error {
	err := &Error{}
	return err
}

// NewError returns a new error from the specified RPC error.
func NewErrorFromRPCError(rpcError *rpc.Error) *Error {
	err := &Error{}
	return err
}

// Error returns a error message.
func (self *Error) Error() string {
	return self.Message
}
