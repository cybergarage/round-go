// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// #cgo CFLAGS: -I/usr/local/include -DROUND_SUPPORT_JS_SM
// #cgo LDFLAGS: -L/usr/local/lib -L/usr/local/opt/openssl/lib -lround -lmupnp -lcrypto -ljansson -lmozjs185
// #include <round/round.h>
import "C"

import (
	"encoding/json"
	"unsafe"

	"github.com/cybergarage/round-go/round"
	"github.com/cybergarage/round-go/round/core/rpc"
)

// A LocalNode represents a local node.
type LocalNode struct {
	*BaseNode
	msgMgr     *MessageManager
	coreEngine *C.RoundLocalNode
}

// NewLocalNode returns a new LocalNode.
func NewLocalNode() *LocalNode {
	node := &LocalNode{BaseNode: NewBaseNode()}

	node.msgMgr = NewMessageManager()
	node.msgMgr.SetListener(node)

	node.coreEngine = (*C.RoundLocalNode)(C.round_local_node_new())

	return node
}

func (self *LocalNode) initDefaultMethods() bool {
	return true
}

// SetRegistry sets a registry by a specified registory.
func (self *LocalNode) SetRegistry(reg *Registry) error {
	return nil
}

// GetRegistry gets a registry by the specified key.
func (self *LocalNode) GetRegistry(key string) (*Registry, bool) {
	return nil, false
}

// Start starts the local node.
func (self *LocalNode) Start() error {
	err := self.msgMgr.Start()
	if err != nil {
		return err
	}

	ok, err := C.round_local_node_start((unsafe.Pointer)(self.coreEngine))
	if !ok {
		return err
	}

	return nil
}

// Stop starts the local node.
func (self *LocalNode) Stop() error {
	err := self.msgMgr.Stop()
	if err != nil {
		return err
	}

	ok, err := C.round_local_node_stop((unsafe.Pointer)(self.coreEngine))
	if !ok {
		return err
	}

	return nil
}

// Convert round error to round.Error.
func (self *LocalNode) CoreError2Error(coreErr *C.RoundError) *round.Error {
	err := round.NewError()
	err.Code = int(C.round_error_getcode((unsafe.Pointer)(coreErr)))
	err.Message = C.GoString(C.round_error_getmessage((unsafe.Pointer)(coreErr)))
	err.DetailCode = int(C.round_error_getdetailcode((unsafe.Pointer)(coreErr)))
	err.DetailMessage = C.GoString(C.round_error_getdetailmessage((unsafe.Pointer)(coreErr)))
	return err
}

// Exec runs the specified JSON request string.
func (self *LocalNode) ExecJSONRequest(reqStr string) (string, *round.Error) {
	var coreRetStr *C.char
	var coreErr C.RoundError
	C.round_error_init((unsafe.Pointer)(&coreErr))
	ok, _ := C.round_local_node_execjsonrequest((unsafe.Pointer)(self.coreEngine), C.CString(reqStr), &coreRetStr, (unsafe.Pointer)(&coreErr))
	if !ok {
		return "", self.CoreError2Error(&coreErr)
	}

	retStr := C.GoString(coreRetStr)
	C.free((unsafe.Pointer)(coreRetStr))

	return retStr, nil
}

// MessageReceived is a listner for MessageManager.
func (self *LocalNode) MessageReceived(msg *Message) *round.Error {
	self.Lock()
	defer self.Unlock()

	var rpcReq rpc.Request
	err := json.Unmarshal([]byte(msg.Content), rpcReq)
	if err != nil {
		return round.NewErrorFromRPCError(rpc.NewError(rpc.ErrorCodeInvalidParams))
	}

	/*
		_, rpcErr := self.Exec(&rpcReq)
		if rpcErr != nil {
			return round.NewErrorFromRPCError(rpcErr)
		}
	*/

	return nil
}
