// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A Clock represents a Clock.
type Clock struct {
	Timestamp int64
}

// NewClock returns a new clock.
func NewClock() *Clock {
	clock := &Clock{0}
	return clock
}
