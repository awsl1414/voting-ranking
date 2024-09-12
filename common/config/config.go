/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-18 17:06:43
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-09-10 16:44:40
 * @FilePath: /voting-ranking/common/config/config.go
 * @Description:
 *
 */

package config

import (
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

// 总配置文件
type config struct {
	Server        server        `yaml:"server"`
	Db            db            `yaml:"db"`
	Log           log           `yaml:"log"`
	Redis         redis         `yaml:"redis"`
	ImageSettings imageSettings `yaml:"imageSettings"`
}

// 项目端口配置
type server struct {
	Address string `yaml:"address"`
	Model   string `yaml:"model"`
}

// 数据库配置
type db struct {
	Dialects string `yaml:"dialects"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

// 日志配置
type log struct {
	Path  string `yaml:"path"`
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
}

// redis配置
type redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}

// 静态文件配置
type imageSettings struct {
	UploadDir string `yaml:"uploadDir"`
	ImageHost string `yaml:"imageHost"`
}

var Config *config

func init() {
	yamlFile, err := os.ReadFile("./config.yaml")
	// 有错直接down
	if err != nil {
		panic(err)
	}

	// 绑定值
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
	// 使用环境变量覆盖配置值
	if envHost := os.Getenv("DB_HOST"); envHost != "" {
		Config.Db.Host = envHost
	}
	if envPort := os.Getenv("DB_PORT"); envPort != "" {
		if port, err := strconv.Atoi(envPort); err == nil {
			Config.Db.Port = port
		}
	}
	if envUser := os.Getenv("DB_USER"); envUser != "" {
		Config.Db.Username = envUser
	}
	if envPass := os.Getenv("DB_PASSWORD"); envPass != "" {
		Config.Db.Password = envPass
	}
	if envRHost := os.Getenv("REDIS_HOST"); envRHost != "" {
		Config.Redis.Address = envRHost
	}
}
