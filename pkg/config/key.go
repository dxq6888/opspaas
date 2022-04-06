package config

type KeyName string

const (
	ServerName KeyName = "server_name"
	ServerHost KeyName = "server_host"
	ServerPort KeyName = "server_port"
	LogPath    KeyName = "log_path"
	LogName    KeyName = "log_name"
	LogDebug   KeyName = "log_debug"
	MyHost     KeyName = "my_host"
	MyUser     KeyName = "my_user"
	MyPasswd   KeyName = "my_passwd"
	MyPort     KeyName = "my_port"
	MyDb       KeyName = "my_db"
)

type Namespaces struct {
	Name string `json:"name"`
}
