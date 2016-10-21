// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import ()

// A FinderListener represents a listener for finder.
type FinderListener interface {
	NodeAliveReceived(node Node)
	NodeDeadReceived(node Node)
}

// A Finder represents a Finder.
type Finder struct {
	Listener FinderListener
}

// NewFinder returns a new Finder.
func NewFinder() *Finder {
	finder := &Finder{}
	return finder
}
