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
