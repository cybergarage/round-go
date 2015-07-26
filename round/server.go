// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package round

import (
	"net/http"

	"github.com/cybergarage/round-go/round/impl"
	"github.com/cybergarage/round-go/round/method"
)

// A Server represents a server.
type Server struct {
	*impl.Server
}

// NewServer returns a new Server.
func NewServer() *Server {
	server := &Server{
		Server: impl.NewServer(),
	}

	server.init()

	return server
}

func (self *Server) init() bool {
	self.initDefaultMethods()
	return true
}

func (self *Server) initDefaultMethods() bool {

	self.SetStaticMethods(method.GetSystemStaticMethods())
	self.SetDynamicMethods(method.GetSystemDynamicMethods())

	return true
}

// ServeHTTP
func (server Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
