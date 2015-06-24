// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"errors"
)

const (
	MessageQueueMax    = 100
	errorChannelClosed = "Message queue chanele is closed"
)

// A MessageQueueListener represents a message listener interface.
type MessageQueueListener interface {
	MessageReceived(*Message)
}

// A MessageQueue represents a message queue.
type MessageQueue struct {
	chanel   chan *Message
	listener MessageQueueListener
}

func messageQueueLoop(msgQueue *MessageQueue) {
	for {
		msg := <-msgQueue.chanel
		if msg == nil {
			break
		}
		if msgQueue.listener != nil {
			msgQueue.listener.MessageReceived(msg)
		}
	}
}

// NewMessageQueue returns a new MessageQueue.
func NewMessageQueue() *MessageQueue {
	msgQueue := &MessageQueue{}
	msgQueue.chanel = nil
	msgQueue.listener = nil
	return msgQueue
}

// SetListener sets a message listener.
func (self *MessageQueue) SetListener(listener MessageQueueListener) error {
	self.listener = listener
	return nil
}

// Start starts the message queue.
func (self *MessageQueue) Start() error {
	err := self.Stop()
	if err != nil {
		return err
	}
	self.chanel = make(chan *Message, MessageQueueMax)
	go messageQueueLoop(self)
	return nil
}

// Stop starts the message queue.
func (self *MessageQueue) Stop() error {
	if self.chanel != nil {
		close(self.chanel)
		self.chanel = nil
	}
	return nil
}

// Stop starts the message queue.
func (self *MessageQueue) PostMessage(msg *Message) error {
	if self.chanel == nil {
		return errors.New(errorChannelClosed)
	}
	self.chanel <- msg
	return nil
}
