// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impl

import (
	"github.com/cybergarage/go-net-upnp/net/upnp"
)

const (
	SetTarget      = "SetTarget"
	NewTargetValue = "newTargetValue"
	GetTarget      = "GetTarget"
	RetTargetValue = "RetTargetValue"
	GetStatus      = "GetStatus"
	ResultStatus   = "ResultStatus"

	DefaultTarget = "Living"
	DefaultStatus = true
)

type Server struct {
	*upnp.Device
	Target string
	Status bool
}

func NewServer() (*Server, error) {
	dev, err := upnp.NewDeviceFromDescription(roundServerDescription)
	if err != nil {
		return nil, err
	}

	service, err := dev.GetServiceByType("urn:schemas-upnp-org:service:SwitchPower:1")
	if err != nil {
		return nil, err
	}

	err = service.LoadDescriptionBytes([]byte(roundServiceDescription))
	if err != nil {
		return nil, err
	}

	lightDev := &Server{
		Device: dev,
		Target: DefaultTarget,
		Status: DefaultStatus,
	}
	lightDev.ActionListener = lightDev

	return lightDev, nil
}

func (self *Server) ActionRequestReceived(action *upnp.Action) upnp.Error {
	switch action.Name {
	case SetTarget:
		target, err := action.GetArgumentString(NewTargetValue)
		if err == nil {
			self.Target = target
		} else {
			return upnp.NewErrorFromCode(upnp.ErrorInvalidArgs)
		}
		return nil
	case GetTarget:
		err := action.SetArgumentString(RetTargetValue, self.Target)
		if err != nil {
			return upnp.NewErrorFromCode(upnp.ErrorInvalidArgs)
		}
		return nil
	case GetStatus:
		err := action.SetArgumentBool(ResultStatus, self.Status)
		if err != nil {
			return upnp.NewErrorFromCode(upnp.ErrorInvalidArgs)
		}
		return nil
	}

	return upnp.NewErrorFromCode(upnp.ErrorOptionalActionNotImplemented)
}
