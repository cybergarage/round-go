// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package round

import (
	"errors"
	"testing"
)

const (
	nullServerError = "Server is null"
)

func TestNewServer(t *testing.T) {
	server := NewServer()
	if server == nil {
		t.Error(errors.New(nullServerError))
	}
}
