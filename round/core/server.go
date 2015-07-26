// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import ()

type Server struct {
	*LocalNode
}

func NewServer() *Server {
	server := &Server{
		LocalNode: NewLocalNode(),
	}
	return server
}
