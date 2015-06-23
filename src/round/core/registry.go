// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"time"
)

// A Registry represents a script.
type Registry struct {
	Key       string
	Value     string
	Timestamp time.Time
}

// NewRegistry returns a new Registry.
func NewRegistry() *Registry {
	reg := &Registry{}
	return reg
}
