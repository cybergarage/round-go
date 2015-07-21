// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"errors"
	"fmt"
	"testing"

	"github.com/cybergarage/round-go/round/log"
)

const (
	TEST_KEY_NODE           = "node"
	TEST_KEY_BINDADDR       = "bind_addr"
	TEST_KEY_BINDPORT       = "bind_port"
	TEST_KEY_BINDADDR_VALUE = "127.0.0.1"
	TEST_KEY_BINDPORT_VALUE = "4649"

	TEST_CONFIG = "{\n" +
		"    \"" + TEST_KEY_NODE + "\": {\n" +
		"        \"" + TEST_KEY_BINDPORT + "\": \"" + TEST_KEY_BINDPORT_VALUE + "\",\n" +
		"        \"" + TEST_KEY_BINDADDR + "\": \"" + TEST_KEY_BINDADDR_VALUE + "\"\n" +
		"    }\n" +
		"}\n"
)

func TestLoadingSimpleConfig(t *testing.T) {
	log.SetSharedLogger(log.NewStdoutLogger(log.LoggerLevelTrace))
	defer log.SetSharedLogger(nil)

	const TEST_SIMPLE_KEY = TEST_KEY_BINDADDR
	const TEST_SIMPLE_VAL = TEST_KEY_BINDADDR_VALUE

	const TEST_CONFIG = "{\n" +
		"\"" + TEST_SIMPLE_KEY + "\": \"" + TEST_SIMPLE_VAL + "\"" +
		"}"

	config, err := NewConfig()
	if err != nil {
		t.Error(err)
	}

	err = config.ParseFromString(TEST_CONFIG)
	if err != nil {
		t.Error(err)
	}

	// TEST_SIMPLE_KEY

	keyValue, err := config.GetKeyStringByXPath(TEST_SIMPLE_KEY)
	if err != nil {
		t.Error(err)
	}

	if len(keyValue) <= 0 {
		t.Error(errors.New(fmt.Sprintf("%s is null", TEST_SIMPLE_KEY)))
	}

	if keyValue != TEST_SIMPLE_VAL {
		t.Error(errors.New(fmt.Sprintf("%s is not equals (%s)", TEST_SIMPLE_VAL, keyValue)))
	}
}

func ParseConfigTest(t *testing.T, s string) {
	log.SetSharedLogger(log.NewStdoutLogger(log.LoggerLevelTrace))
	defer log.SetSharedLogger(nil)

	config, err := NewConfig()
	if err != nil {
		t.Error(err)
	}

	err = config.ParseFromString(s)
	if err != nil {
		t.Error(err)
	}

	// /node/bind_addr

	xpath := TEST_KEY_NODE + "/" + TEST_KEY_BINDADDR
	keyValue, err := config.GetKeyStringByXPath(xpath)
	if err != nil {
		t.Error(err)
	}

	if len(keyValue) <= 0 {
		t.Error(errors.New(fmt.Sprintf("%s is null", xpath)))
	}

	if keyValue != TEST_KEY_BINDADDR_VALUE {
		t.Error(errors.New(fmt.Sprintf("%s is not equals (%s)", TEST_KEY_BINDADDR_VALUE, keyValue)))
	}

	// /node/bind_port

	xpath = TEST_KEY_NODE + "/" + TEST_KEY_BINDPORT
	keyValue, err = config.GetKeyStringByXPath(xpath)
	if err != nil {
		t.Error(err)
	}

	if len(keyValue) <= 0 {
		t.Error(errors.New(fmt.Sprintf("%s is null", xpath)))
	}

	if keyValue != TEST_KEY_BINDPORT_VALUE {
		t.Error(errors.New(fmt.Sprintf("%s is not equals (%s)", TEST_KEY_BINDPORT_VALUE, keyValue)))
	}
}

func TestLoadingConfig(t *testing.T) {
	ParseConfigTest(t, TEST_CONFIG)
}

func TestLoadingCommentedConfig(t *testing.T) {
	TEST_COMMENTED_CONFIG :=
		"####\n" +
			TEST_CONFIG
	ParseConfigTest(t, TEST_COMMENTED_CONFIG)

	TEST_COMMENTED_CONFIG =
		"####\n" +
			"####\n" +
			TEST_CONFIG
	ParseConfigTest(t, TEST_COMMENTED_CONFIG)

	TEST_COMMENTED_CONFIG =
		" ####\n" +
			TEST_CONFIG
	ParseConfigTest(t, TEST_COMMENTED_CONFIG)
}

func TestLoadingBlankConfig(t *testing.T) {
	TEST_COMMENTED_CONFIG :=
		"\n" +
			TEST_CONFIG
	ParseConfigTest(t, TEST_COMMENTED_CONFIG)

	TEST_COMMENTED_CONFIG =
		"\n" +
			"\n" +
			TEST_CONFIG
	ParseConfigTest(t, TEST_COMMENTED_CONFIG)
}

func TestLoadingCommentAndBlankConfig(t *testing.T) {
	TEST_COMMENTED_CONFIG :=
		"####\n" +
			"\n" +
			TEST_CONFIG
	ParseConfigTest(t, TEST_COMMENTED_CONFIG)
}
