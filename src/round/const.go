// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package round

const (
	system_trigger_timer_precision       = 0.1
	rpc_http_endpoint                    = "/rpc"
	rpc_http_content_type                = "application/json"
	node_status_unknown                  = 0
	node_status_stop                     = 1
	node_status_activating               = 2
	node_status_active                   = 3
	node_status_terminating              = 4
	node_status_promice                  = 5
	system_method_post_job               = "post_job"
	system_method_set_method             = "set_method"
	system_method_set_alias              = "set_alias"
	system_method_set_route              = "set_route"
	system_method_set_timer              = "set_timer"
	system_method_add_node               = "add_node"
	system_method_remove_method          = "remove_method"
	system_method_remove_alias           = "remove_alias"
	system_method_remove_route           = "remove_route"
	system_method_remove_timer           = "remove_timer"
	system_method_remove_node            = "remove_node"
	system_method_echo                   = "_echo"
	system_method_get_network_state      = "get_network_state"
	system_method_get_cluster_state      = "get_cluster_state"
	system_method_set_registry           = "set_registry"
	system_method_get_registry           = "get_registry"
	system_method_remove_registry        = "remove_registry"
	system_method_get_node_state         = "get_node_state"
	system_method_get_node_config        = "get_node_config"
	system_method_param_name             = "name"
	system_method_param_language         = "language"
	system_method_param_code             = "code"
	system_method_param_encode           = "encode"
	system_method_param_base64           = "base64"
	system_method_param_src              = "src"
	system_method_param_dest             = "dest"
	system_method_param_type             = "type"
	system_method_param_cond             = "cond"
	system_method_param_addr             = "addr"
	system_method_param_port             = "port"
	system_method_param_hash             = "hash"
	system_method_param_version          = "version"
	system_method_param_nodes            = "nodes"
	system_method_param_cluster          = "cluster"
	system_method_param_clusters         = "clusters"
	system_method_param_state            = "state"
	system_method_param_method           = "method"
	system_method_param_defaults         = "defaults"
	system_method_param_start_time       = "start_time"
	system_method_param_stop_time        = "stop_time"
	system_method_param_duration         = "duration"
	system_method_param_loop             = "loop"
	system_method_add_node_cmd           = "/usr/local/bin/roundd"
	system_method_remove_node_delay_msec = 5000
	system_method_dest_all_node          = "*"
	system_method_dest_any_node          = "?"
	script_post_method                   = "post_method"
	script_print_method                  = "print"
	script_search_method                 = "search"
	script_search_method_wait_msec       = 5000
	message_searching                    = "Searching"
	message_done                         = "Done"
)
