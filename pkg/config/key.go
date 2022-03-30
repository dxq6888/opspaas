package config

type KeyName string

const (
	ServerName KeyName = "server_name"
	ServerHost KeyName = "server_host"
	ServerPort KeyName = "server_port"
	LogPath    KeyName = "log_path"
	LogName    KeyName = "log_name"
	LogDebug   KeyName = "log_debug"
)
