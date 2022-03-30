package config

import (
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/klog/v2"
	"strconv"
)

var keyMap map[KeyName]string

type Config struct {
	Server Server
	Logger Logger
}

type Server struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Logger struct {
	LogPath  string `yaml:"path"`
	LogName  string `yaml:"name"`
	LogDebug string `yaml:"debug"`
}

func init() {
	var config Config
	yamlFile, err := ioutil.ReadFile("./conf/config.yaml")
	if err != nil {
		klog.Fatal(err)
		return
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		klog.Fatal(err)
		return
	}
	keyMap = make(map[KeyName]string)
	keyMap[ServerName] = config.Server.Name
	keyMap[ServerHost] = config.Server.Host
	keyMap[ServerPort] = config.Server.Port
	keyMap[LogPath] = config.Logger.LogPath
	keyMap[LogName] = config.Logger.LogName
	keyMap[LogDebug] = config.Logger.LogDebug
}

func GetString(KeyName KeyName) string {
	return keyMap[KeyName]
}

func GetInt(keyName KeyName) int {
	intStr := keyMap[keyName]
	if intStr == "" {
		klog.Fatal("GetInt not read config ===>" + keyName)
		return -1
	}
	v, err := strconv.Atoi(intStr)
	if intStr == "" {
		klog.Fatal(err)
		return -1
	}
	return v
}


