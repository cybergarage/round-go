// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package method

import (
	"github.com/cybergarage/round-go/round/core"
)

func GetSystemStaticMethods() []core.Method {
	return []core.Method{}
}

func GetSystemDynamicMethods() []core.Method {
	return []core.Method{
		NewSetRegistryMethod(),
		NewGetRegistryMethod(),
	}
}
