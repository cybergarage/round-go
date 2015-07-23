// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package round

const (
	SystemTriggerTimerPrecision     = 0.1
	UpnpDeviceType                  = "urn:cybergarage-org:device:round:1"
	UpnpService                     = "urn:cybergarage-org:service:round:1"
	RpcHttpEndpoint                 = "/rpc"
	RpcHttpContentType              = "application/json"
	NodeStatusUnknown               = 0
	NodeStatusStop                  = 1
	NodeStatusActivating            = 2
	NodeStatusActive                = 3
	NodeStatusTerminating           = 4
	NodeStatusPromice               = 5
	SystemMethodPostJob             = "post_job"
	SystemMethodSetMethod           = "set_method"
	SystemMethodSetAlias            = "set_alias"
	SystemMethodSetRoute            = "set_route"
	SystemMethodSetTimer            = "set_timer"
	SystemMethodAddNode             = "add_node"
	SystemMethodRemoveMethod        = "remove_method"
	SystemMethodRemoveAlias         = "remove_alias"
	SystemMethodRemoveRoute         = "remove_route"
	SystemMethodRemoveTimer         = "remove_timer"
	SystemMethodRemoveNode          = "remove_node"
	SystemMethodEcho                = "_echo"
	SystemMethodGetNetworkState     = "get_network_state"
	SystemMethodGetClusterState     = "get_cluster_state"
	SystemMethodSetRegistry         = "set_registry"
	SystemMethodGetRegistry         = "get_registry"
	SystemMethodRemoveRegistry      = "remove_registry"
	SystemMethodGetNodeState        = "get_node_state"
	SystemMethodGetNodeConfig       = "get_node_config"
	SystemMethodParamName           = "name"
	SystemMethodParamLanguage       = "language"
	SystemMethodParamCode           = "code"
	SystemMethodParamEncode         = "encode"
	SystemMethodParamBase64         = "base64"
	SystemMethodParamSrc            = "src"
	SystemMethodParamDest           = "dest"
	SystemMethodParamType           = "type"
	SystemMethodParamCond           = "cond"
	SystemMethodParamAddr           = "addr"
	SystemMethodParamPort           = "port"
	SystemMethodParamHash           = "hash"
	SystemMethodParamVersion        = "version"
	SystemMethodParamNodes          = "nodes"
	SystemMethodParamCluster        = "cluster"
	SystemMethodParamClusters       = "clusters"
	SystemMethodParamState          = "state"
	SystemMethodParamMethod         = "method"
	SystemMethodParamDefaults       = "defaults"
	SystemMethodParamStartTime      = "start_time"
	SystemMethodParamStopTime       = "stop_time"
	SystemMethodParamDuration       = "duration"
	SystemMethodParamLoop           = "loop"
	SystemMethodAddNodeCmd          = "/usr/local/bin/roundd"
	SystemMethodRemoveNodeDelayMsec = 5000
	SystemMethodDestAllNode         = "*"
	SystemMethodDestAnyNode         = "?"
	ScriptPostMethod                = "post_method"
	ScriptPrintMethod               = "print"
	ScriptSearchMethod              = "search"
	ScriptSearchMethodWaitMsec      = 5000
	MessageSearching                = "Searching"
	MessageDone                     = "Done"
)
