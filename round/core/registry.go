// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"time"
)

// A Registry represents a script.
type Registry struct {
	Key   string
	Value string
	Ts    int64
	Lts   int64
}

// NewRegistry returns a new Registry.
func NewRegistry(key, value string) *Registry {
	newReg := &Registry{Key: key, Value: value, Ts: time.Now().Unix(), Lts: 0}
	return newReg
}

// NewRegistry returns a new Registry.
func NewRegistryFromRegistry(reg *Registry) *Registry {
	newReg := &Registry{Key: reg.Key, Value: reg.Value, Ts: time.Now().Unix(), Lts: reg.Lts}
	return newReg
}
