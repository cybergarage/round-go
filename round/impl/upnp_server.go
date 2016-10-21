// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impl

import (
	"github.com/cybergarage/round-go/round/core"

	"github.com/cybergarage/go-net-upnp/net/upnp"
)

type Server struct {
	*core.Server
	*upnp.Device
	Target string
	Status bool
}

func NewServer() *Server {
	dev, err := upnp.NewDeviceFromDescription(roundServerDeviceDescription)
	if err != nil {
		return nil
	}

	service, err := dev.GetServiceByType(UpnpServiceType)
	if err != nil {
		return nil
	}

	err = service.LoadDescriptionBytes([]byte(roundServerServiceDescription))
	if err != nil {
		return nil
	}

	server := &Server{
		Server: core.NewServer(),
		Device: dev,
	}
	server.ActionListener = server

	return server
}

func (self *Server) ActionRequestReceived(action *upnp.Action) upnp.Error {
	switch action.Name {
	}

	return upnp.NewErrorFromCode(upnp.ErrorOptionalActionNotImplemented)
}
