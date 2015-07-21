// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package config loads configuration files for Round.

Config gets a setting value in the specified file by the given path like XPath.

	config, err := NewConfig()
	if err != nil {
		t.Error(err)
	}

	err = config.ParseFromFile("/etc/roundd.conf")
	if err != nil {
		t.Error(err)
	}

	keyValue, err := config.GetStringByXPath("/bind_port")
	if err != nil {
		t.Error(err)
	}

The configuration file fomat is based on JSON as the following.

	#
	#  roundd.conf
	#

	{
		"bind_port": <port number>,
		"bind_addr": "<if address>",
		"cluster": "<name>",
		"log_file": "file name",
		"methods": [
				<method_name> : {
					"language" : <supported-language>
					"code" : <code>
					"encoding" : ("none" | "base64")
				}
			]
	}
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	PathSep = "/"
	LineSep = "\n"
	Comment = "#"
)

const (
	errorConfigKeyNull        = "Path is null"
	errorConfigKeyNotFound    = "Key (%s) is not found"
	errorConfigKeyTypeInvalid = "Key (%s) type is invalid"
)

type Config struct {
	FileName   string
	rootObject interface{}
}

// NewConfig returns a new Config.
func NewConfig() (*Config, error) {
	config := &Config{}
	return config, nil
}

// NewConfigFromFile returns a new Config from the given file.
func NewConfigFromFile(file string) (*Config, error) {
	config := &Config{}
	err := config.ParseFromFile(file)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// NewConfigFromString returns a new Config from the given string.
func NewConfigFromString(s string) (*Config, error) {
	config := &Config{}
	err := config.ParseFromString(s)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// ParseFromFile parses the given file.
func (config *Config) ParseFromFile(file string) error {
	config.FileName = file

	_, err := os.Stat(file)
	if err != nil {
		return err
	}

	sourceBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return config.ParseFromString(string(sourceBytes))
}

// ParseFromString parses the given string.
func (config *Config) ParseFromString(source string) error {
	lines := strings.Split(source, LineSep)

	// Strip comment and null lines
	var strippedSource bytes.Buffer
	for _, line := range lines {
		if len(line) <= 0 {
			continue
		}
		commentIdx := strings.Index(line, Comment)
		if 0 <= commentIdx {
			continue
		}
		strippedSource.WriteString(line + LineSep)
	}

	return json.Unmarshal(strippedSource.Bytes(), &config.rootObject)
}

// getKeyObjectFromObject returns a object the given key.
func (config *Config) getKeyObjectFromObject(key string, obj interface{}) (interface{}, error) {
	switch obj.(type) {
	case map[string]interface{}:
		jsonDir, _ := obj.(map[string]interface{})
		keyObj, hasKey := jsonDir[key]
		if !hasKey {
			return "", errors.New(fmt.Sprintf(errorConfigKeyNotFound, key))
		}
		switch keyObj.(type) {
		case string:
			return keyObj, nil
		case map[string]interface{}:
			return keyObj, nil
		default:
			return "", errors.New(fmt.Sprintf(errorConfigKeyTypeInvalid, key))
		}
	}
	return "", nil
}

// GetKey returns a object the given paths.
func (config *Config) getPathObjectFromObject(paths []string, rootObj interface{}) (interface{}, error) {
	obj := rootObj
	for _, path := range paths {
		keyObj, err := config.getKeyObjectFromObject(path, obj)
		if err != nil {
			return nil, err
		}
		obj = keyObj
	}

	return obj, nil
}

// GetStringByPaths returns a key value by the given paths.
func (config *Config) GetKeyStringByPaths(paths []string) (string, error) {
	keyValue := ""

	keyObj, err := config.getPathObjectFromObject(paths, config.rootObject)
	if err != nil {
		return "", err
	}

	switch keyObj.(type) {
	case string:
		keyValue, _ = keyObj.(string)
	default:
		return "", errors.New(fmt.Sprintf(errorConfigKeyTypeInvalid, (PathSep + strings.Join(paths, PathSep))))
	}

	return keyValue, nil
}

// GetStringByXPath returns a key value by the given XPath.
func (config *Config) GetKeyStringByXPath(path string) (string, error) {
	paths := strings.Split(path, PathSep)
	if len(paths) <= 0 {
		return "", errors.New(errorConfigKeyNull)
	}
	return config.GetKeyStringByPaths(paths)
}
