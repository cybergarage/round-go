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
	errorMessageQueueBadCnt = "invalid chan cnt (%d) : ecpected %d"
	errorMessageQueueBadSeq = "invalid chan seq (%d) : ecpected %d"
	errorMessageQueueBadSum = "invalid chan sum (%d) : ecpected %d"
)

type TestMessageQueueListener struct {
	Values []int
	Sum    int
}

func (self *TestMessageQueueListener) MessageReceived(msg *Message) {
	value, err := strconv.Atoi(msg.Content)
	if err == nil {
		self.Values = append(self.Values, value)
		self.Sum += value
	}
}

func TestNewMessageQueue(t *testing.T) {
	msgQueue := NewMessageQueue()
	msgQueueListener := &TestMessageQueueListener{Sum: 0, Values: make([]int, 0)}
	msgQueue.SetListener(msgQueueListener)

	err := msgQueue.Start()
	if err != nil {
		t.Error(err)
	}

	msgValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
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

	// Check FIFO

	if len(msgValues) == len(msgQueueListener.Values) {
		for n, msgValue := range msgValues {
			if msgValue != msgQueueListener.Values[n] {
				t.Errorf(errorMessageQueueBadSeq, msgQueueListener.Values[n], msgValue)
			}
		}
	} else {
		t.Errorf(errorMessageQueueBadCnt, len(msgQueueListener.Values), len(msgValues))
	}

	// Check Sum

	expectedResult := 0
	for _, msgValue := range msgValues {
		expectedResult += msgValue
	}

	if msgQueueListener.Sum != expectedResult {
		t.Errorf(errorMessageQueueBadSum, msgQueueListener.Sum, expectedResult)
	}
}
