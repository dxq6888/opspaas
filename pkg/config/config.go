package config

import (
	"fmt"
	yaml2 "gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s.io/klog/v2"
	"strconv"
)

var keyMap map[KeyName]string

type Config struct {
	Server Server
	Logger Logger
	Mysql  Mysql
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

type Mysql struct {
	MyHost   string `yaml:"host"`
	MyUser   string `yaml:"user"`
	MyPasswd string `yaml:"passwd"`
	MyDb     string `yaml:"db"`
	MyPort   string `yaml:"port"`
}

func init() {
	var config Config
	yamlFile, err := ioutil.ReadFile("./conf/config.yaml")
	if err != nil {
		klog.Fatal(err)
		return
	}
	err = yaml2.Unmarshal(yamlFile, &config)
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
	keyMap[MyHost] = config.Mysql.MyHost
	keyMap[MyUser] = config.Mysql.MyUser
	keyMap[MyPasswd] = config.Mysql.MyPasswd
	keyMap[MyDb] = config.Mysql.MyDb
	keyMap[MyPort] = config.Mysql.MyPort
	fmt.Println("config init-------->", keyMap)
}

func GetString(KeyName KeyName) string {
	//fmt.Println("config GetString-------->",keyMap[KeyName])
	return keyMap[KeyName]
}

func GetInt(keyName KeyName) int {
	intStr := keyMap[keyName]
	if intStr == "" {
		//logger := tools.InitLogger()
		//logger.Info("GetInt not read config ===>",zapcore.Field{Interface: keyName})
		fmt.Println("GetInt not read config ===>", keyName)
		return -1
	}
	v, err := strconv.Atoi(intStr)
	if intStr == "" {
		klog.Fatal(err)
		return -1
	}
	return v
}
