// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// ActionGroupInterface is a wrapper around the C record GActionGroupInterface.
type ActionGroupInterface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for has_action
	// no type for list_actions
	// no type for get_action_enabled
	// no type for get_action_parameter_type
	// no type for get_action_state_type
	// no type for get_action_state_hint
	// no type for get_action_state
	// no type for change_action_state
	// no type for activate_action
	// no type for action_added
	// no type for action_removed
	// no type for action_enabled_changed
	// no type for action_state_changed
	// no type for query_action
}

// ActionInterface is a wrapper around the C record GActionInterface.
type ActionInterface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for get_name
	// no type for get_parameter_type
	// no type for get_state_type
	// no type for get_state_hint
	// no type for get_enabled
	// no type for get_state
	// no type for change_state
	// no type for activate
}

// ApplicationClass is a wrapper around the C record GApplicationClass.
type ApplicationClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for startup
	// no type for activate
	// no type for open
	// no type for command_line
	// no type for local_command_line
	// no type for before_emit
	// no type for after_emit
	// no type for add_platform_data
	// no type for quit_mainloop
	// no type for run_mainloop
	// no type for shutdown
	// no type for dbus_register
	// no type for dbus_unregister
	// no type for handle_local_options
	// Private : padding
}

// ApplicationCommandLineClass is a wrapper around the C record GApplicationCommandLineClass.
type ApplicationCommandLineClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for print_literal
	// no type for printerr_literal
	// no type for get_stdin
	// Private : padding
}

// PollableInputStreamInterface is a wrapper around the C record GPollableInputStreamInterface.
type PollableInputStreamInterface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for can_poll
	// no type for is_readable
	// no type for create_source
	// no type for read_nonblocking
}

// PollableOutputStreamInterface is a wrapper around the C record GPollableOutputStreamInterface.
type PollableOutputStreamInterface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for can_poll
	// no type for is_writable
	// no type for create_source
	// no type for write_nonblocking
}

// TlsBackendInterface is a wrapper around the C record GTlsBackendInterface.
type TlsBackendInterface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for supports_tls
	// no type for get_certificate_type
	// no type for get_client_connection_type
	// no type for get_server_connection_type
	// no type for get_file_database_type
	// no type for get_default_database
	// no type for supports_dtls
	// no type for get_dtls_client_connection_type
	// no type for get_dtls_server_connection_type
}
