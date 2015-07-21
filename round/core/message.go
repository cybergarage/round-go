// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A Message represents a JSON-RPC Message.
type Message struct {
	Content string
}

// NewMessage returns a new Message.
func NewMessage(content string) *Message {
	msg := &Message{Content: content}
	return msg
}
