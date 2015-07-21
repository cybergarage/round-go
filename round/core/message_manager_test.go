// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"round"
	"strconv"
	"testing"
	"time"
)

const (
	errorMessageManagerBadCnt = "invalid chan cnt (%d) : ecpected %d"
	errorMessageManagerBadSeq = "invalid chan seq (%d) : ecpected %d"
	errorMessageManagerBadSum = "invalid chan sum (%d) : ecpected %d"
)

type TestMessageManagerListener struct {
	Values []int
	Sum    int
}

func (self *TestMessageManagerListener) MessageReceived(msg *Message) *round.Error {
	value, err := strconv.Atoi(msg.Content)
	if err == nil {
		self.Values = append(self.Values, value)
		self.Sum += value
	}
	return nil
}

func TestNewMessageManager(t *testing.T) {
	msgMgr := NewMessageManager()
	msgMgrListener := &TestMessageManagerListener{Sum: 0, Values: make([]int, 0)}
	msgMgr.SetListener(msgMgrListener)

	err := msgMgr.Start()
	if err != nil {
		t.Error(err)
	}

	msgValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, msgValue := range msgValues {
		err := msgMgr.PostMessage(NewMessage(strconv.Itoa(msgValue)))
		if err != nil {
			t.Error(err)
		}
	}

	time.Sleep(time.Second)

	err = msgMgr.Stop()
	if err != nil {
		t.Error(err)
	}

	// Check FIFO

	if len(msgValues) == len(msgMgrListener.Values) {
		for n, msgValue := range msgValues {
			if msgValue != msgMgrListener.Values[n] {
				t.Errorf(errorMessageManagerBadSeq, msgMgrListener.Values[n], msgValue)
			}
		}
	} else {
		t.Errorf(errorMessageManagerBadCnt, len(msgMgrListener.Values), len(msgValues))
	}

	// Check Sum

	expectedResult := 0
	for _, msgValue := range msgValues {
		expectedResult += msgValue
	}

	if msgMgrListener.Sum != expectedResult {
		t.Errorf(errorMessageManagerBadSum, msgMgrListener.Sum, expectedResult)
	}
}
