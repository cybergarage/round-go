// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

import (
	"container/list"
	"sync"
)

// A ContentHashGraph represents a message queue.
type ContentHashGraph struct {
	*list.List
	*sync.Mutex
}

// NewContentHashGraph returns a new ContentHashGraph.
func NewContentHashGraph() *ContentHashGraph {
	cluster := &ContentHashGraph{}
	cluster.List = &list.List{}
	cluster.Mutex = &sync.Mutex{}
	return cluster
}

// AddNode adds a new node into the content hash graph.
func (self *ContentHashGraph) AddObject(obj interface{}) bool {
	hashObj, ok := IsHashObject(obj)
	if !ok {
		return false
	}
	_ = hashObj.GetHashSeed()

	self.Lock()
	defer self.Unlock()

	for e := self.Front(); e != nil; e = e.Next() {
	}

	return true
}

// RemveObject removes the specified node in the content hash graph.
func (self *ContentHashGraph) RemoveObject(obj interface{}) bool {
	hashObj, ok := IsHashObject(obj)
	if !ok {
		return false
	}
	_ = hashObj.GetHashSeed()

	self.Lock()
	defer self.Unlock()

	for e := self.Front(); e != nil; e = e.Next() {
	}

	return true
}
