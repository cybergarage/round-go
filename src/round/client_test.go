// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package round

import (
	"errors"
	"testing"
)

const (
	nullClientError         = "Clinet is null"
)

func TestNewClient(t *testing.T) {
	clinet := NewClinet()
	if clinet == nil {
		t.Error(errors.New(nullClientError))
	}
}
