// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/cybergarage/round-go/round/common"
)

// A Cluster represents a message queue.
type Cluster struct {
	*common.ContentHashGraph
}

// NewCluster returns a new Cluster.
func NewCluster() *Cluster {
	cluster := &Cluster{
		ContentHashGraph: common.NewContentHashGraph(),
	}
	return cluster
}

// AddNode adds a new node into the content hash graph.
func (self *Cluster) AddNode(node interface{}) error {
	return self.AddObject(node)
}

// RemoveNode removes the specified node in the content hash graph.
func (self *Cluster) RemoveNode(node interface{}) error {
	return self.RemoveObject(node)
}

// GetHandleNode returns a handle node of the specfied hash object.
func (self *Cluster) GetHandleNode(targetHashObj common.HashObject) (interface{}, error) {
	return self.GetHandleObject(targetHashObj)
}
