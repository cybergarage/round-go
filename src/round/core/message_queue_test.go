// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"strconv"
	"testing"
	"time"
)

const (
	errorMessageQueueBadSum = "invalid sum (%d) : ecpected %d"
)

var msgQueueSum int

type TestMessageQueueListener struct {
}

func (self TestMessageQueueListener) MessageReceived(msg *Message) {
	value, err := strconv.Atoi(msg.Content)
	if err == nil {
		msgQueueSum += value
	}
}

func TestNewMessageQueue(t *testing.T) {
	msgQueue := NewMessageQueue()
	msgQueueListener := TestMessageQueueListener{}
	msgQueue.SetListener(msgQueueListener)

	msgQueueSum = 0

	err := msgQueue.Start()
	if err != nil {
		t.Error(err)
	}

	msgValues := []int{1, 2, 3}
	for _, msgValue := range msgValues {
		err := msgQueue.PostMessage(NewMessage(strconv.Itoa(msgValue)))
		if err != nil {
			t.Error(err)
		}
	}

	time.Sleep(time.Second)

	err = msgQueue.Stop()
	if err != nil {
		t.Error(err)
	}

	expectedResult := 0
	for _, msgValue := range msgValues {
		expectedResult += msgValue
	}

	if msgQueueSum != expectedResult {
		t.Errorf(errorMessageQueueBadSum, msgQueueSum, expectedResult)
	}
}
