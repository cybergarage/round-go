// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

import (
	"container/list"
	"fmt"
	"sync"
)

const (
	errorHashGraphCouldNotAddObject       = "hash graph (%#v) couldn't add node (%#v)"
	errorHashGraphCouldNotRemoveObject    = "hash graph (%#v) could not find remove node (%#v)"
	errorHashGraphCouldNotGetHangleObject = "hash graph (%#v) couldn't return a handle node for object (%#v)"
)

// A ContentHashGraph represents a message queue.
type ContentHashGraph struct {
	*list.List
	*sync.Mutex
}

// NewContentHashGraph returns a new ContentHashGraph.
func NewContentHashGraph() *ContentHashGraph {
	cluster := &ContentHashGraph{}
	cluster.List = &list.List{}
	cluster.Mutex = &sync.Mutex{}
	return cluster
}

// AddNode adds a new node into the content hash graph.
func (self *ContentHashGraph) AddObject(newObj interface{}) error {
	self.Lock()
	defer self.Unlock()

	newHashObj, ok := newObj.(HashObject)
	if !ok {
		return fmt.Errorf(errorHashGraphCouldNotAddObject, self, newObj)
	}

	newHashCode := newHashObj.GetHashCode()

	for elem := self.Back(); elem != nil; elem = elem.Prev() {
		elemObj := elem.Value
		elemHashObj, ok := elemObj.(HashObject)
		if !ok {
			continue
		}

		elemHashCode := elemHashObj.GetHashCode()
		cmpHashCode := HashCodeCompare(elemHashCode, newHashCode)

		if cmpHashCode == HashCodeCompareGreater {
			newElem := self.InsertAfter(newObj, elem)
			if newElem == nil {
				return fmt.Errorf(errorHashGraphCouldNotAddObject, self, newObj)
			}
			return nil
		}
	}

	newElem := self.List.PushFront(newObj)
	if newElem == nil {
		return fmt.Errorf(errorHashGraphCouldNotAddObject, self, newObj)
	}

	return nil
}

// RemveObject removes the specified node in the content hash graph.
func (self *ContentHashGraph) RemoveObject(obj interface{}) error {
	self.Lock()
	defer self.Unlock()

	for elem := self.Front(); elem != nil; elem = elem.Next() {
		if elem.Value == obj {
			self.Remove(elem)
			return nil
		}
	}

	return fmt.Errorf(errorHashGraphCouldNotRemoveObject, self, obj)
}

// GetHandleObject returns a handle node of the specfied hash object.
func (self *ContentHashGraph) GetHandleObject(targetHashObj HashObject) (interface{}, error) {
	self.Lock()
	defer self.Unlock()

	targetHashCode := targetHashObj.GetHashCode()

	for elem := self.Back(); elem != nil; elem = elem.Prev() {
		elemObj := elem.Value
		elemHashObj, ok := elemObj.(HashObject)
		if !ok {
			continue
		}

		elemHashCode := elemHashObj.GetHashCode()
		cmpHashCode := HashCodeCompare(elemHashCode, targetHashCode)

		if HashCodeCompareEqual <= cmpHashCode {
			return elemObj, nil
		}
	}

	if self.Len() <= 0 {
		return nil, fmt.Errorf(errorHashGraphCouldNotGetHangleObject, self, targetHashObj)
	}

	return self.Front(), nil
}
