// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

import (
	"testing"
)

const (
	errorContentHashGraphInvalideSize = "content hash graph size is %d: expected %d"
	errorContentHashGraphInvalideNode = "content hash graph node is invalid %#v: expected %#v"
)

func TestNewContentHashGraph(t *testing.T) {
	NewContentHashGraph()
}

func TestStaticContentHashGraphClockwise(t *testing.T) {
	const testNodeCnt = 10
	testNodes := make([]*TestStaticContentHashObject, testNodeCnt)
	for n := 0; n < testNodeCnt; n++ {
		testNodes[n] = NewTestStaticContentHashObject(n)
	}

	// add nodes

	hg := NewContentHashGraph()

	if hg.Len() != 0 {
		t.Errorf(errorContentHashGraphInvalideSize, hg.Len(), 0)
	}

	for n := 0; n < testNodeCnt; n++ {
		hg.AddObject(testNodes[n])
		if hg.Len() != (n + 1) {
			t.Errorf(errorContentHashGraphInvalideSize, hg.Len(), (n + 1))
		}
	}
	if hg.Len() != testNodeCnt {
		t.Errorf(errorContentHashGraphInvalideSize, hg.Len(), testNodeCnt)
	}

	n := 0
	for obj := hg.Front(); obj != nil; obj = obj.Next() {
		node, _ := obj.Value.(*TestStaticContentHashObject)
		if node != testNodes[n] {
			t.Errorf(errorContentHashGraphInvalideNode, node, testNodes[n])
		}
		n++
	}

	// handle node

	n = 0
	for obj := hg.Front(); obj != nil; obj = obj.Next() {
		node := testNodes[n]
		handleObj, err := hg.GetHandleObject(node)
		if err != nil {
			t.Error(err)
		}
		if handleObj != node {
			t.Errorf(errorContentHashGraphInvalideNode, handleObj, node)
		}
		n++
	}

	// remove nodes

	for n := 0; n < testNodeCnt; n++ {
		hg.RemoveObject(testNodes[n])
		if hg.Len() != (testNodeCnt - (n + 1)) {
			t.Errorf(errorContentHashGraphInvalideSize, hg.Len(), (testNodeCnt - (n + 1)))
		}
	}

	if hg.Len() != 0 {
		t.Errorf(errorContentHashGraphInvalideSize, hg.Len(), 0)
	}
}

func TestStaticContentHashGraphCounterclockwise(t *testing.T) {
	const testNodeCnt = 10
	testNodes := make([]*TestStaticContentHashObject, testNodeCnt)
	for n := 0; n < testNodeCnt; n++ {
		testNodes[n] = NewTestStaticContentHashObject(n)
	}

	// add nodes

	hg := NewContentHashGraph()

	if hg.Len() != 0 {
		t.Errorf(errorContentHashGraphInvalideSize, hg.Len(), 0)
	}

	for n := 0; n < testNodeCnt; n++ {
		hg.AddObject(testNodes[(testNodeCnt - (n + 1))])
		if hg.Len() != (n + 1) {
			t.Errorf(errorContentHashGraphInvalideSize, hg.Len(), (n + 1))
		}
	}
	if hg.Len() != testNodeCnt {
		t.Errorf(errorContentHashGraphInvalideSize, hg.Len(), testNodeCnt)
	}

	n := 0
	for obj := hg.Front(); obj != nil; obj = obj.Next() {
		node, _ := obj.Value.(*TestStaticContentHashObject)
		if node != testNodes[n] {
			t.Errorf(errorContentHashGraphInvalideNode, node, testNodes[n])
		}
		n++
	}

	// handle node

	n = 0
	for obj := hg.Front(); obj != nil; obj = obj.Next() {
		node := testNodes[n]
		handleObj, err := hg.GetHandleObject(node)
		if err != nil {
			t.Error(err)
		}
		if handleObj != node {
			t.Errorf(errorContentHashGraphInvalideNode, handleObj, node)
		}
		n++
	}

	// remove nodes

	for n := 0; n < testNodeCnt; n++ {
		hg.RemoveObject(testNodes[(testNodeCnt - (n + 1))])
		if hg.Len() != (testNodeCnt - (n + 1)) {
			t.Errorf(errorContentHashGraphInvalideSize, hg.Len(), (testNodeCnt - (n + 1)))
		}
	}

	if hg.Len() != 0 {
		t.Errorf(errorContentHashGraphInvalideSize, hg.Len(), 0)
	}
}
