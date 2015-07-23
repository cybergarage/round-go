// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impl

import (
	"github.com/cybergarage/go-net-upnp/net/upnp"
	"github.com/cybergarage/go-net-upnp/net/upnp/log"
	"github.com/cybergarage/go-net-upnp/net/upnp/ssdp"
)

// A FinderListener represents a listener for finder.
type FinderListener interface {
	NodeAdded(node *Node)
	NodeRemoved(node *Node)
}

// A Finder represents a Finder.
type Finder struct {
	*upnp.ControlPoint
	Listener FinderListener
}

// NewFinder returns a new Finder.
func NewFinder() *Finder {
	finder := &Finder{
		ControlPoint: upnp.NewControlPoint(),
	}

	finder.ControlPoint.Listener = finder

	return finder
}

func (self *Finder) DeviceNotifyReceived(req *ssdp.Request) {
	if self.Listener == nil {
		return
	}

	node, err := NewNodeFromSSDPRequest(req)
	if err != nil {
		log.Warn(err)
		return
	}

	self.Listener.NodeAdded(node)
}

func (self *Finder) DeviceSearchReceived(req *ssdp.Request) {
	if self.Listener == nil {
		return
	}

	node, err := NewNodeFromSSDPRequest(req)
	if err != nil {
		log.Warn(err)
		return
	}

	self.Listener.NodeAdded(node)
}

func (self *Finder) DeviceResponseReceived(res *ssdp.Response) {
	if self.Listener == nil {
		return
	}

	node, err := NewNodeFromSSDPResponse(res)
	if err != nil {
		log.Warn(err)
		return
	}

	self.Listener.NodeAdded(node)
}
