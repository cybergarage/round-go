// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impl

import (
	"github.com/cybergarage/go-net-upnp/net/upnp"
)

type Server struct {
	*upnp.Device
	Target string
	Status bool
}

func NewServer() (*Server, error) {
	dev, err := upnp.NewDeviceFromDescription(roundServerDeviceDescription)
	if err != nil {
		return nil, err
	}

	service, err := dev.GetServiceByType(UpnpServiceType)
	if err != nil {
		return nil, err
	}

	err = service.LoadDescriptionBytes([]byte(roundServerServiceDescription))
	if err != nil {
		return nil, err
	}

	server := &Server{
		Device: dev,
	}
	server.ActionListener = server

	return server, nil
}

func (self *Server) ActionRequestReceived(action *upnp.Action) upnp.Error {
	switch action.Name {
	}

	return upnp.NewErrorFromCode(upnp.ErrorOptionalActionNotImplemented)
}
