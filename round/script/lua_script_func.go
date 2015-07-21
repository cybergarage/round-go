// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package script

// #cgo CFLAGS: -I/usr/local/include -DROUND_SUPPORT_LUA -I../../../_obj
// #cgo LDFLAGS: -L/usr/local/lib -lround -llua
// #include <roundc/round.h>
// #include "_cgo_export.h"
//int round_lua_setregistry(lua_State* L)
//{
//  const char *key = luaL_checkstring(L, 1);
//  const char *val = luaL_checkstring(L, 2);
//
//  bool isSuccess = (LocalNodeSetRegistry((char *)key, (char *)val) != 1) ? true : false;
//  lua_pushboolean(L, isSuccess);
//
//  return 1;
//}
//
//int round_lua_getregistry(lua_State* L)
//{
//  bool isSuccess = true;
//
//  const char *key = luaL_checkstring(L, 1);
//  const char *result = LocalNodeGetRegistry((char *)key);
//
//  lua_pushboolean(L, isSuccess);
//  lua_pushstring(L, result);
//
//  return 2;
//}
import "C"
