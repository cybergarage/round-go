// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impl

import (
	"github.com/cybergarage/go-net-upnp/net/upnp/ssdp"
)

// A Node represents a Node.
type Node struct {
	Address string
	Port    int
}

// NewNode returns a new Node.
func NewNode() *Node {
	node := &Node{}
	return node
}

// newNodeFromSSDPPacket returns a new Node from the specified SSDP packet
func newNodeFromSSDPPacket(ssdpPkt *ssdp.Packet) (*Node, error) {
	from := ssdpPkt.From
	node := &Node{
		Address: from.IP.String(),
		Port:    from.Port,
	}
	return node, nil
}

// NewNodeFromSSDPPacket returns a new Node from the specified SSDP request
func NewNodeFromSSDPRequest(req *ssdp.Request) (*Node, error) {
	return newNodeFromSSDPPacket(req.Packet)
}

// NewNodeFromSSDPResponse returns a new Node from the specified SSDP response
func NewNodeFromSSDPResponse(res *ssdp.Response) (*Node, error) {
	return newNodeFromSSDPPacket(res.Packet)
}

func (self *Node) GetRequestAddress() string {
	return self.Address
}

func (self *Node) GetRequestPort() int {
	return self.Port
}

func (self *Node) PostMessage(reqMsg string) (resMsg string, err error) {
	return "", nil
}

func (self *Node) GetClock() uint64 {
	return 0
}
