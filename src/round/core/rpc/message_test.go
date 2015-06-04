// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"errors"
	"testing"
)

const (
	nullMessageError = "Message is null"
)

func TestNewMessage(t *testing.T) {
	msg := NewMessage()
	if msg == nil {
		t.Error(errors.New(nullMessageError))
	}
}
