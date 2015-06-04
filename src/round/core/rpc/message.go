// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

// A Message represents a JSON-RPC Message.
type Message struct {
}

// NewRequest returns a new Request.
func NewMessage() *Message {
	msg := &Message{}
	return msg
}
