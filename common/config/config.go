/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-18 17:06:43
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-19 18:27:33
 * @FilePath: /voting-ranking/common/config/config.go
 * @Description:
 *
 */

package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// 总配置文件
type config struct {
	Server server `yaml:"server"`
	Db     db     `yaml:"db"`
	Log    log    `yaml:"log"`
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
}
