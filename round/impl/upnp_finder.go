// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impl

import (
	"fmt"

	"github.com/cybergarage/round-go/round/core"

	"github.com/cybergarage/go-net-upnp/net/upnp"
	"github.com/cybergarage/go-net-upnp/net/upnp/log"
	"github.com/cybergarage/go-net-upnp/net/upnp/ssdp"
	"github.com/cybergarage/go-net-upnp/net/upnp/util"
)

// A Finder represents a Finder.
type Finder struct {
	*core.Finder
	*upnp.ControlPoint
}

// NewFinder returns a new Finder.
func NewFinder() *Finder {
	finder := &Finder{
		Finder:       core.NewFinder(),
		ControlPoint: upnp.NewControlPoint(),
	}

	finder.ControlPoint.Listener = finder

	return finder
}

func (self *Finder) DeviceNotifyReceived(req *ssdp.Request) {
	usn, _ := req.GetUSN()
	log.Trace(fmt.Sprintf("ssdp notiry req : %s %s", usn, ssdpPacketToMessage(req.Packet)))

	if self.Finder.Listener == nil {
		return
	}

	node, err := NewNodeFromSSDPRequest(req)
	if err != nil {
		log.Warn(err)
		return
	}

	switch {
	case req.IsAlive():
		self.Finder.Listener.NodeAliveReceived(node)
	case req.IsByeBye():
		self.Finder.Listener.NodeDeadReceived(node)
	}
}

func (self *Finder) DeviceSearchReceived(req *ssdp.Request) {
}

func (self *Finder) DeviceResponseReceived(res *ssdp.Response) {
	url, _ := res.GetLocation()
	log.Trace(fmt.Sprintf("search res : %s %s", url, ssdpPacketToMessage(res.Packet)))

	if self.Finder.Listener == nil {
		return
	}

	node, err := NewNodeFromSSDPResponse(res)
	if err != nil {
		log.Warn(err)
		return
	}

	self.Finder.Listener.NodeAliveReceived(node)
}

func ssdpPacketToMessage(req *ssdp.Packet) string {
	fromAddr := req.From.String()
	toAddr := ""
	ifAddr, err := util.GetInterfaceAddress(req.Interface)
	if err == nil {
		toAddr = ifAddr
	}

	return fmt.Sprintf("(%s -> %s)", fromAddr, toAddr)
}
