// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A Clock represents a Clock.
type Clock struct {
	Timestamp uint64
}

// NewClock returns a new clock.
func NewClock() *Clock {
	clock := &Clock{0}
	return clock
}

func (self *Clock) SetClock(value uint64) {
	self.Timestamp = value
}

func (self *Clock) GetClock() uint64 {
	return self.Timestamp
}
