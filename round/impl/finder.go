// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impl

import (
	"github.com/cybergarage/go-net-upnp/net/upnp"
	"github.com/cybergarage/go-net-upnp/net/upnp/ssdp"
)

// A Finder represents a Finder.
type Finder struct {
	*upnp.ControlPoint
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
}

func (self *Finder) DeviceSearchReceived(req *ssdp.Request) {
}

func (self *Finder) DeviceResponseReceived(res *ssdp.Response) {
}
