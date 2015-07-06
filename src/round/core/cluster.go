// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A Cluster represents a message queue.
type Cluster struct {
}

// NewCluster returns a new Cluster.
func NewCluster() *Cluster {
	cluster := &Cluster{}
	return cluster
}
