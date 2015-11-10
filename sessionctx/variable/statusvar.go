// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package variable

import (
	"strings"

	"github.com/juju/errors"
)

// StatusVars is global status vars map.
var StatusVars map[string]*SysVar

// ErrUnknownStatusVar is used when can't find the status variable name.
var ErrUnknownStatusVar = errors.New("Unknown status var")

// GetStatusVar returns status var infomation for name.
func GetStatusVar(name string) *SysVar {
	name = strings.ToLower(name)
	return StatusVars[name]
}

func init() {
	StatusVars = make(map[string]*SysVar)
	for _, v := range defaultStatusVars {
		StatusVars[v.Name] = v
	}
}

var defaultStatusVars = []*SysVar{
	{ScopeGlobal, "aborted_clients", "0"},
	{ScopeGlobal, "binlog_cache_disk_use", "0"},
	{ScopeGlobal, "binlog_cache_use", "0"},
	{ScopeGlobal, "binlog_stmt_cache_disk_use", "0"},
	{ScopeGlobal, "binlog_stmt_cache_use", "0"},
	{ScopeGlobal | ScopeSession, "bytes_received", "2262"},
	{ScopeGlobal | ScopeSession, "bytes_sent", "106142"},
	{ScopeGlobal | ScopeSession, "com_admin_commands", "0"},
	{ScopeGlobal | ScopeSession, "com_assign_to_keycache", "0"},
	{ScopeGlobal | ScopeSession, "com_alter_db", "0"},
	{ScopeGlobal | ScopeSession, "com_alter_db_upgrade", "0"},
	{ScopeGlobal | ScopeSession, "com_alter_event", "0"},
	{ScopeGlobal | ScopeSession, "com_alter_function", "0"},
	{ScopeGlobal | ScopeSession, "com_alter_procedure", "0"},
	{ScopeGlobal | ScopeSession, "com_alter_server", "0"},
	{ScopeGlobal | ScopeSession, "com_alter_table", "0"},
	{ScopeGlobal | ScopeSession, "com_alter_tablespace", "0"},
	{ScopeGlobal | ScopeSession, "com_alter_user", "0"},
	{ScopeGlobal | ScopeSession, "com_analyze", "0"},
	{ScopeGlobal | ScopeSession, "com_begin", "0"},
	{ScopeGlobal | ScopeSession, "com_binlog", "0"},
	{ScopeGlobal | ScopeSession, "com_call_procedure", "0"},
	{ScopeGlobal | ScopeSession, "com_change_db", "2"},
	{ScopeGlobal | ScopeSession, "com_change_master", "0"},
	{ScopeGlobal | ScopeSession, "com_change_repl_filter", "0"},
	{ScopeGlobal | ScopeSession, "com_check", "0"},
	{ScopeGlobal | ScopeSession, "com_checksum", "0"},
	{ScopeGlobal | ScopeSession, "com_commit", "0"},
	{ScopeGlobal | ScopeSession, "com_create_db", "1"},
	{ScopeGlobal | ScopeSession, "com_create_event", "0"},
	{ScopeGlobal | ScopeSession, "com_create_function", "0"},
	{ScopeGlobal | ScopeSession, "com_create_index", "0"},
	{ScopeGlobal | ScopeSession, "com_create_procedure", "0"},
	{ScopeGlobal | ScopeSession, "com_create_server", "0"},
	{ScopeGlobal | ScopeSession, "com_create_table", "1"},
	{ScopeGlobal | ScopeSession, "com_create_trigger", "0"},
	{ScopeGlobal | ScopeSession, "com_create_udf", "0"},
	{ScopeGlobal | ScopeSession, "com_create_user", "0"},
	{ScopeGlobal | ScopeSession, "com_create_view", "0"},
	{ScopeGlobal | ScopeSession, "com_dealloc_sql", "0"},
	{ScopeGlobal | ScopeSession, "com_delete", "0"},
	{ScopeGlobal | ScopeSession, "com_delete_multi", "0"},
	{ScopeGlobal | ScopeSession, "com_do", "0"},
	{ScopeGlobal | ScopeSession, "com_drop_db", "0"},
	{ScopeGlobal | ScopeSession, "com_drop_event", "0"},
	{ScopeGlobal | ScopeSession, "com_drop_function", "0"},
	{ScopeGlobal | ScopeSession, "com_drop_index", "0"},
	{ScopeGlobal | ScopeSession, "com_drop_procedure", "0"},
	{ScopeGlobal | ScopeSession, "com_drop_server", "0"},
	{ScopeGlobal | ScopeSession, "com_drop_table", "0"},
	{ScopeGlobal | ScopeSession, "com_drop_trigger", "0"},
	{ScopeGlobal | ScopeSession, "com_drop_user", "0"},
	{ScopeGlobal | ScopeSession, "com_drop_view", "0"},
	{ScopeGlobal | ScopeSession, "com_empty_query", "0"},
	{ScopeGlobal | ScopeSession, "com_execute_sql", "0"},
	{ScopeGlobal | ScopeSession, "com_explain_other", "0"},
	{ScopeGlobal | ScopeSession, "com_flush", "0"},
	{ScopeGlobal | ScopeSession, "com_get_diagnostics", "0"},
	{ScopeGlobal | ScopeSession, "com_grant", "0"},
	{ScopeGlobal | ScopeSession, "com_ha_close", "0"},
	{ScopeGlobal | ScopeSession, "com_ha_open", "0"},
	{ScopeGlobal | ScopeSession, "com_ha_read", "0"},
	{ScopeGlobal | ScopeSession, "com_help", "0"},
	{ScopeGlobal | ScopeSession, "com_insert", "0"},
	{ScopeGlobal | ScopeSession, "com_insert_select", "0"},
	{ScopeGlobal | ScopeSession, "com_install_plugin", "0"},
	{ScopeGlobal | ScopeSession, "com_kill", "0"},
	{ScopeGlobal | ScopeSession, "com_load", "0"},
	{ScopeGlobal | ScopeSession, "com_lock_tables", "0"},
	{ScopeGlobal | ScopeSession, "com_optimize", "0"},
	{ScopeGlobal | ScopeSession, "com_preload_keys", "0"},
	{ScopeGlobal | ScopeSession, "com_prepare_sql", "0"},
	{ScopeGlobal | ScopeSession, "com_purge", "0"},
	{ScopeGlobal | ScopeSession, "com_purge_before_date", "0"},
	{ScopeGlobal | ScopeSession, "com_release_savepoint", "0"},
	{ScopeGlobal | ScopeSession, "com_rename_table", "0"},
	{ScopeGlobal | ScopeSession, "com_rename_user", "0"},
	{ScopeGlobal | ScopeSession, "com_repair", "0"},
	{ScopeGlobal | ScopeSession, "com_replace", "0"},
	{ScopeGlobal | ScopeSession, "com_replace_select", "0"},
	{ScopeGlobal | ScopeSession, "com_reset", "0"},
	{ScopeGlobal | ScopeSession, "com_resignal", "0"},
	{ScopeGlobal | ScopeSession, "com_revoke", "0"},
	{ScopeGlobal | ScopeSession, "com_revoke_all", "0"},
	{ScopeGlobal | ScopeSession, "com_rollback", "0"},
	{ScopeGlobal | ScopeSession, "com_rollback_to_savepoint", "0"},
	{ScopeGlobal | ScopeSession, "com_savepoint", "0"},
	{ScopeGlobal | ScopeSession, "com_select", "9"},
	{ScopeGlobal | ScopeSession, "com_set_option", "0"},
	{ScopeGlobal | ScopeSession, "com_signal", "0"},
	{ScopeGlobal | ScopeSession, "com_show_binlog_events", "0"},
	{ScopeGlobal | ScopeSession, "com_show_binlogs", "0"},
	{ScopeGlobal | ScopeSession, "com_show_charsets", "0"},
	{ScopeGlobal | ScopeSession, "com_show_collations", "0"},
	{ScopeGlobal | ScopeSession, "com_show_create_db", "0"},
	{ScopeGlobal | ScopeSession, "com_show_create_event", "0"},
	{ScopeGlobal | ScopeSession, "com_show_create_func", "0"},
	{ScopeGlobal | ScopeSession, "com_show_create_proc", "0"},
	{ScopeGlobal | ScopeSession, "com_show_create_table", "0"},
	{ScopeGlobal | ScopeSession, "com_show_create_trigger", "0"},
	{ScopeGlobal | ScopeSession, "com_show_databases", "1"},
	{ScopeGlobal | ScopeSession, "com_show_engine_logs", "0"},
	{ScopeGlobal | ScopeSession, "com_show_engine_mutex", "0"},
	{ScopeGlobal | ScopeSession, "com_show_engine_status", "0"},
	{ScopeGlobal | ScopeSession, "com_show_events", "0"},
	{ScopeGlobal | ScopeSession, "com_show_errors", "0"},
	{ScopeGlobal | ScopeSession, "com_show_fields", "0"},
	{ScopeGlobal | ScopeSession, "com_show_function_code", "0"},
	{ScopeGlobal | ScopeSession, "com_show_function_status", "0"},
	{ScopeGlobal | ScopeSession, "com_show_grants", "0"},
	{ScopeGlobal | ScopeSession, "com_show_keys", "0"},
	{ScopeGlobal | ScopeSession, "com_show_master_status", "0"},
	{ScopeGlobal | ScopeSession, "com_show_open_tables", "0"},
	{ScopeGlobal | ScopeSession, "com_show_plugins", "0"},
	{ScopeGlobal | ScopeSession, "com_show_privileges", "0"},
	{ScopeGlobal | ScopeSession, "com_show_procedure_code", "0"},
	{ScopeGlobal | ScopeSession, "com_show_procedure_status", "0"},
	{ScopeGlobal | ScopeSession, "com_show_processlist", "0"},
	{ScopeGlobal | ScopeSession, "com_show_profile", "0"},
	{ScopeGlobal | ScopeSession, "com_show_profiles", "0"},
	{ScopeGlobal | ScopeSession, "com_show_relaylog_events", "0"},
	{ScopeGlobal | ScopeSession, "com_show_slave_hosts", "0"},
	{ScopeGlobal | ScopeSession, "com_show_slave_status", "0"},
	{ScopeGlobal | ScopeSession, "com_show_status", "25"},
	{ScopeGlobal | ScopeSession, "com_show_storage_engines", "0"},
	{ScopeGlobal | ScopeSession, "com_show_table_status", "0"},
	{ScopeGlobal | ScopeSession, "com_show_tables", "2"},
	{ScopeGlobal | ScopeSession, "com_show_triggers", "0"},
	{ScopeGlobal | ScopeSession, "com_show_variables", "4"},
	{ScopeGlobal | ScopeSession, "com_show_warnings", "2"},
	{ScopeGlobal | ScopeSession, "com_show_create_user", "0"},
	{ScopeGlobal | ScopeSession, "com_shutdown", "0"},
	{ScopeGlobal | ScopeSession, "com_slave_start", "0"},
	{ScopeGlobal | ScopeSession, "com_slave_stop", "0"},
	{ScopeGlobal | ScopeSession, "com_group_replication_start", "0"},
	{ScopeGlobal | ScopeSession, "com_group_replication_stop", "0"},
	{ScopeGlobal | ScopeSession, "com_stmt_execute", "0"},
	{ScopeGlobal | ScopeSession, "com_stmt_close", "0"},
	{ScopeGlobal | ScopeSession, "com_stmt_fetch", "0"},
	{ScopeGlobal | ScopeSession, "com_stmt_prepare", "0"},
	{ScopeGlobal | ScopeSession, "com_stmt_reset", "0"},
	{ScopeGlobal | ScopeSession, "com_stmt_send_long_data", "0"},
	{ScopeGlobal | ScopeSession, "com_truncate", "0"},
	{ScopeGlobal | ScopeSession, "com_uninstall_plugin", "0"},
	{ScopeGlobal | ScopeSession, "com_unlock_tables", "0"},
	{ScopeGlobal | ScopeSession, "com_update", "0"},
	{ScopeGlobal | ScopeSession, "com_update_multi", "0"},
	{ScopeGlobal | ScopeSession, "com_xa_commit", "0"},
	{ScopeGlobal | ScopeSession, "com_xa_end", "0"},
	{ScopeGlobal | ScopeSession, "com_xa_prepare", "0"},
	{ScopeGlobal | ScopeSession, "com_xa_recover", "0"},
	{ScopeGlobal | ScopeSession, "com_xa_rollback", "0"},
	{ScopeGlobal | ScopeSession, "com_xa_start", "0"},
	{ScopeGlobal | ScopeSession, "com_stmt_reprepare", "0"},
	{ScopeSession, "compression", "off"},
	{ScopeGlobal, "connection_errors_accept", "0"},
	{ScopeGlobal, "connection_errors_internal", "0"},
	{ScopeGlobal, "connection_errors_max_connections", "0"},
	{ScopeGlobal, "connection_errors_peer_address", "0"},
	{ScopeGlobal, "connection_errors_select", "0"},
	{ScopeGlobal, "connection_errors_tcpwrap", "0"},
	{ScopeGlobal, "connections", "3"},
	{ScopeGlobal | ScopeSession, "created_tmp_disk_tables", "0"},
	{ScopeGlobal, "created_tmp_files", "6"},
	{ScopeGlobal | ScopeSession, "created_tmp_tables", "3"},
	{ScopeGlobal, "delayed_errors", "0"},
	{ScopeGlobal, "delayed_insert_threads", "0"},
	{ScopeGlobal, "delayed_writes", "0"},
	{ScopeGlobal, "flush_commands", "1"},
	{ScopeGlobal | ScopeSession, "handler_commit", "5"},
	{ScopeGlobal | ScopeSession, "handler_delete", "0"},
	{ScopeGlobal | ScopeSession, "handler_discover", "0"},
	{ScopeGlobal | ScopeSession, "handler_external_lock", "265"},
	{ScopeGlobal | ScopeSession, "handler_mrr_init", "0"},
	{ScopeGlobal | ScopeSession, "handler_prepare", "0"},
	{ScopeGlobal | ScopeSession, "handler_read_first", "8"},
	{ScopeGlobal | ScopeSession, "handler_read_key", "6"},
	{ScopeGlobal | ScopeSession, "handler_read_last", "0"},
	{ScopeGlobal | ScopeSession, "handler_read_next", "1"},
	{ScopeGlobal | ScopeSession, "handler_read_prev", "0"},
	{ScopeGlobal | ScopeSession, "handler_read_rnd", "0"},
	{ScopeGlobal | ScopeSession, "handler_read_rnd_next", "7642"},
	{ScopeGlobal | ScopeSession, "handler_rollback", "0"},
	{ScopeGlobal | ScopeSession, "handler_savepoint", "0"},
	{ScopeGlobal | ScopeSession, "handler_savepoint_rollback", "0"},
	{ScopeGlobal | ScopeSession, "handler_update", "0"},
	{ScopeGlobal | ScopeSession, "handler_write", "6"},
	{ScopeGlobal, "innodb_buffer_pool_dump_status", "notstarted"},
	{ScopeGlobal, "innodb_buffer_pool_load_status", "Bufferpool(s)loadcom"},
	{ScopeGlobal, "innodb_buffer_pool_resize_status", "notstarted"},
	{ScopeGlobal, "innodb_buffer_pool_pages_data", "440"},
	{ScopeGlobal, "innodb_buffer_pool_bytes_data", "7208960"},
	{ScopeGlobal, "innodb_buffer_pool_pages_dirty", "0"},
	{ScopeGlobal, "innodb_buffer_pool_bytes_dirty", "0"},
	{ScopeGlobal, "innodb_buffer_pool_pages_flushed", "54"},
	{ScopeGlobal, "innodb_buffer_pool_pages_free", "7751"},
	{ScopeGlobal, "innodb_buffer_pool_pages_misc", "0"},
	{ScopeGlobal, "innodb_buffer_pool_pages_total", "8191"},
	{ScopeGlobal, "innodb_buffer_pool_read_ahead_rnd", "0"},
	{ScopeGlobal, "innodb_buffer_pool_read_ahead", "0"},
	{ScopeGlobal, "innodb_buffer_pool_read_ahead_evicted", "0"},
	{ScopeGlobal, "innodb_buffer_pool_read_requests", "1466"},
	{ScopeGlobal, "innodb_buffer_pool_reads", "402"},
	{ScopeGlobal, "innodb_buffer_pool_wait_free", "0"},
	{ScopeGlobal, "innodb_buffer_pool_write_requests", "400"},
	{ScopeGlobal, "innodb_data_fsyncs", "24"},
	{ScopeGlobal, "innodb_data_pending_fsyncs", "0"},
	{ScopeGlobal, "innodb_data_pending_reads", "0"},
	{ScopeGlobal, "innodb_data_pending_writes", "0"},
	{ScopeGlobal, "innodb_data_read", "6656512"},
	{ScopeGlobal, "innodb_data_reads", "426"},
	{ScopeGlobal, "innodb_data_writes", "84"},
	{ScopeGlobal, "innodb_data_written", "1232384"},
	{ScopeGlobal, "innodb_dblwr_pages_written", "20"},
	{ScopeGlobal, "innodb_dblwr_writes", "2"},
	{ScopeGlobal, "innodb_log_waits", "0"},
	{ScopeGlobal, "innodb_log_write_requests", "14"},
	{ScopeGlobal, "innodb_log_writes", "8"},
	{ScopeGlobal, "innodb_os_log_fsyncs", "12"},
	{ScopeGlobal, "innodb_os_log_pending_fsyncs", "0"},
	{ScopeGlobal, "innodb_os_log_pending_writes", "0"},
	{ScopeGlobal, "innodb_os_log_written", "17920"},
	{ScopeGlobal, "innodb_page_size", "16384"},
	{ScopeGlobal, "innodb_pages_created", "39"},
	{ScopeGlobal, "innodb_pages_read", "401"},
	{ScopeGlobal, "innodb_pages_written", "54"},
	{ScopeGlobal, "innodb_row_lock_current_waits", "0"},
	{ScopeGlobal, "innodb_row_lock_time", "0"},
	{ScopeGlobal, "innodb_row_lock_time_avg", "0"},
	{ScopeGlobal, "innodb_row_lock_time_max", "0"},
	{ScopeGlobal, "innodb_row_lock_waits", "0"},
	{ScopeGlobal, "innodb_rows_deleted", "0"},
	{ScopeGlobal, "innodb_rows_inserted", "0"},
	{ScopeGlobal, "innodb_rows_read", "8"},
	{ScopeGlobal, "innodb_rows_updated", "0"},
	{ScopeGlobal, "innodb_num_open_files", "17"},
	{ScopeGlobal, "innodb_truncated_status_writes", "0"},
	{ScopeGlobal, "innodb_available_undo_logs", "128"},
	{ScopeSession, "last_query_cost", "104.799000"},
	{ScopeSession, "last_query_partial_plans", "1"},
	{ScopeGlobal, "key_blocks_not_flushed", "0"},
	{ScopeGlobal, "key_blocks_unused", "6695"},
	{ScopeGlobal, "key_blocks_used", "3"},
	{ScopeGlobal, "key_read_requests", "6"},
	{ScopeGlobal, "key_reads", "3"},
	{ScopeGlobal, "key_write_requests", "0"},
	{ScopeGlobal, "key_writes", "0"},
	{ScopeGlobal | ScopeSession, "locked_connects", "0"},
	{ScopeGlobal | ScopeSession, "max_execution_time_exceeded", "0"},
	{ScopeGlobal | ScopeSession, "max_execution_time_set", "0"},
	{ScopeGlobal | ScopeSession, "max_execution_time_set_failed", "0"},
	{ScopeGlobal, "max_used_connections", "1"},
	{ScopeGlobal, "max_used_connections_time", "2015-11-0902:49:42"},
	{ScopeGlobal, "not_flushed_delayed_rows", "0"},
	{ScopeGlobal, "ongoing_anonymous_transaction_count", "0"},
	{ScopeGlobal, "open_files", "14"},
	{ScopeGlobal, "open_streams", "0"},
	{ScopeGlobal, "open_table_definitions", "105"},
	{ScopeGlobal | ScopeSession, "Open_tables", "101"},
	{ScopeGlobal, "opened_files", "142"},
	{ScopeGlobal | ScopeSession, "opened_table_definitions", "106"},
	{ScopeGlobal | ScopeSession, "opened_tables", "108"},
	{ScopeGlobal, "performance_schema_accounts_lost", "0"},
	{ScopeGlobal, "performance_schema_cond_classes_lost", "0"},
	{ScopeGlobal, "performance_schema_cond_instances_lost", "0"},
	{ScopeGlobal, "performance_schema_digest_lost", "0"},
	{ScopeGlobal, "performance_schema_file_classes_lost", "0"},
	{ScopeGlobal, "performance_schema_file_handles_lost", "0"},
	{ScopeGlobal, "performance_schema_file_instances_lost", "0"},
	{ScopeGlobal, "performance_schema_hosts_lost", "0"},
	{ScopeGlobal, "performance_schema_index_stat_lost", "0"},
	{ScopeGlobal, "performance_schema_locker_lost", "0"},
	{ScopeGlobal, "performance_schema_memory_classes_lost", "0"},
	{ScopeGlobal, "performance_schema_metadata_lock_lost", "0"},
	{ScopeGlobal, "performance_schema_mutex_classes_lost", "0"},
	{ScopeGlobal, "performance_schema_mutex_instances_lost", "0"},
	{ScopeGlobal, "performance_schema_nested_statement_lost", "0"},
	{ScopeGlobal, "performance_schema_prepared_statements_lost", "0"},
	{ScopeGlobal, "performance_schema_program_lost", "0"},
	{ScopeGlobal, "performance_schema_rwlock_classes_lost", "0"},
	{ScopeGlobal, "performance_schema_rwlock_instances_lost", "0"},
	{ScopeGlobal, "performance_schema_session_connect_attrs_lost", "0"},
	{ScopeGlobal, "performance_schema_socket_classes_lost", "0"},
	{ScopeGlobal, "performance_schema_socket_instances_lost", "0"},
	{ScopeGlobal, "performance_schema_stage_classes_lost", "0"},
	{ScopeGlobal, "performance_schema_statement_classes_lost", "0"},
	{ScopeGlobal, "performance_schema_table_handles_lost", "0"},
	{ScopeGlobal, "performance_schema_table_instances_lost", "0"},
	{ScopeGlobal, "performance_schema_table_lock_stat_lost", "0"},
	{ScopeGlobal, "performance_schema_thread_classes_lost", "0"},
	{ScopeGlobal, "performance_schema_thread_instances_lost", "0"},
	{ScopeGlobal, "performance_schema_users_lost", "0"},
	{ScopeGlobal, "prepared_stmt_count", "0"},
	{ScopeGlobal, "qcache_free_blocks", "1"},
	{ScopeGlobal, "qcache_free_memory", "1031832"},
	{ScopeGlobal, "qcache_hits", "0"},
	{ScopeGlobal, "qcache_inserts", "0"},
	{ScopeGlobal, "qcache_lowmem_prunes", "0"},
	{ScopeGlobal, "qcache_not_cached", "5"},
	{ScopeGlobal, "qcache_queries_in_cache", "0"},
	{ScopeGlobal, "qcache_total_blocks", "1"},
	{ScopeGlobal | ScopeSession, "queries", "56"},
	{ScopeGlobal | ScopeSession, "questions", "54"},
	{ScopeGlobal | ScopeSession, "select_full_join", "0"},
	{ScopeGlobal | ScopeSession, "select_full_range_join", "0"},
	{ScopeGlobal | ScopeSession, "select_range", "0"},
	{ScopeGlobal | ScopeSession, "select_range_check", "0"},
	{ScopeGlobal | ScopeSession, "select_scan", "24"},
	{ScopeGlobal, "slave_open_temp_tables", "0"},
	{ScopeGlobal | ScopeSession, "slow_launch_threads", "0"},
	{ScopeGlobal | ScopeSession, "slow_queries", "0"},
	{ScopeGlobal | ScopeSession, "sort_merge_passes", "0"},
	{ScopeGlobal | ScopeSession, "sort_range", "0"},
	{ScopeGlobal | ScopeSession, "sort_rows", "0"},
	{ScopeGlobal | ScopeSession, "sort_scan", "0"},
	{ScopeGlobal, "ssl_accept_renegotiates", "0"},
	{ScopeGlobal, "ssl_accepts", "0"},
	{ScopeGlobal, "ssl_callback_cache_hits", "0"},
	{ScopeGlobal | ScopeSession, "ssl_cipher", ""},
	{ScopeGlobal | ScopeSession, "ssl_cipher_list", ""},
	{ScopeGlobal, "ssl_client_connects", "0"},
	{ScopeGlobal, "ssl_connect_renegotiates", "0"},
	{ScopeGlobal, "ssl_ctx_verify_depth", "0"},
	{ScopeGlobal, "ssl_ctx_verify_mode", "0"},
	{ScopeGlobal | ScopeSession, "Ssl_default_timeout", "0"},
	{ScopeGlobal, "ssl_finished_accepts", "0"},
	{ScopeGlobal, "ssl_finished_connects", "0"},
	{ScopeGlobal | ScopeSession, "ssl_server_not_after", ""},
	{ScopeGlobal | ScopeSession, "ssl_server_not_before", ""},
	{ScopeGlobal, "ssl_session_cache_hits", "0"},
	{ScopeGlobal, "ssl_session_cache_misses", "0"},
	{ScopeGlobal, "ssl_session_cache_mode", "NONE"},
	{ScopeGlobal, "ssl_session_cache_overflows", "0"},
	{ScopeGlobal, "ssl_session_cache_size", "0"},
	{ScopeGlobal, "ssl_session_cache_timeouts", "0"},
	{ScopeGlobal | ScopeSession, "ssl_sessions_reused", "0"},
	{ScopeGlobal, "ssl_used_session_cache_entries", "0"},
	{ScopeGlobal | ScopeSession, "ssl_verify_depth", "0"},
	{ScopeGlobal | ScopeSession, "ssl_verify_mode", "0"},
	{ScopeGlobal | ScopeSession, "ssl_version", ""},
	{ScopeGlobal, "table_locks_immediate", "123"},
	{ScopeGlobal, "table_locks_waited", "0"},
	{ScopeGlobal | ScopeSession, "table_open_cache_hits", "28"},
	{ScopeGlobal | ScopeSession, "table_open_cache_misses", "108"},
	{ScopeGlobal | ScopeSession, "table_open_cache_overflows", "0"},
	{ScopeGlobal, "tc_log_max_pages_used", "0"},
	{ScopeGlobal, "tc_log_page_size", "0"},
	{ScopeGlobal, "tc_log_page_waits", "0"},
	{ScopeGlobal, "threads_cached", "0"},
	{ScopeGlobal, "threads_connected", "1"},
	{ScopeGlobal, "threads_created", "1"},
	{ScopeGlobal, "threads_running", "1"},
	{ScopeGlobal, "uptime", "15998"},
	{ScopeGlobal, "uptime_since_flush_status", "15998"},
}
