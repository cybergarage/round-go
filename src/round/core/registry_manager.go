// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"errors"
	"fmt"
	"time"
)

const (
	ErrorRegistryNotFound = "%s is not found"
)

// A RegistryManager represents a registry manager.
type RegistryManager struct {
	Map map[string]Registry
}

// NewRegistryManager returns a new RegistryManager.
func NewRegistryManager() *RegistryManager {
	regMgr := &RegistryManager{}
	regMgr.Map = map[string]Registry{}
	return regMgr
}

// Set sets a registry.
func (self *RegistryManager) Set(reg Registry) error {
	reg.Timestamp = time.Now()
	self.Map[reg.Key] = reg
	return nil
}

// Get gets a registry by the specified key.
func (self *RegistryManager) Get(key string) (*Registry, error) {
	reg, ok := self.Map[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf(ErrorRegistryNotFound, key))
	}
	return &reg, nil
}
