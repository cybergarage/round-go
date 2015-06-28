// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"errors"
	"round"
)

const (
	MessageManagerMax  = 100
	errorChannelClosed = "Message queue chanele is closed"
)

// A MessageManagerListener represents a message listener interface.
type MessageManagerListener interface {
	MessageReceived(*Message) *round.Error
}

// A MessageManager represents a message queue.
type MessageManager struct {
	chanel   chan *Message
	listener MessageManagerListener
}

func messageQueueLoop(msgQueue *MessageManager) {
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

// NewMessageManager returns a new MessageManager.
func NewMessageManager() *MessageManager {
	msgQueue := &MessageManager{}
	msgQueue.chanel = nil
	msgQueue.listener = nil
	return msgQueue
}

// SetListener sets a message listener.
func (self *MessageManager) SetListener(listener MessageManagerListener) error {
	self.listener = listener
	return nil
}

// Start starts the message queue.
func (self *MessageManager) Start() error {
	err := self.Stop()
	if err != nil {
		return err
	}
	self.chanel = make(chan *Message, MessageManagerMax)
	go messageQueueLoop(self)
	return nil
}

// Stop starts the message queue.
func (self *MessageManager) Stop() error {
	if self.chanel != nil {
		close(self.chanel)
		self.chanel = nil
	}
	return nil
}

// Stop starts the message queue.
func (self *MessageManager) PostMessage(msg *Message) error {
	if self.chanel == nil {
		return errors.New(errorChannelClosed)
	}
	self.chanel <- msg
	return nil
}
