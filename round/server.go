// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package round

import (
	"net/http"

	//"github.com/cybergarage/round-go/round/impl"
)

// A Server represents a server.
type Server struct {
	//*impl.Server
}

// NewServer returns a new Server.
func NewServer() *Server {
	server := &Server{}
	return server
}

// ServeHTTP
func (server Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
