package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

type Server struct {
	Addr string `mapstructure:"addr"`
}

type UserCenter struct {
	PrivateUrl string `mapstructure:"private-url"`
	PublicUrl  string `mapstructure:"public-url"`
}

type BlockchainAPI struct {
	Url string `mapstructure:"url"`
}

type Redis struct {
	Uri      string `mapstructure:"uri"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
}

type Mysql struct {
	DSN string `mapstructure:"dsn"`
}

type Mongo struct {
	Uri string `mapstructure:"uri"`
}

type DB struct {
	Redis Redis `mapstructure:"redis"`
	Mysql Mysql `mapstructure:"mysql"`
	Mongo Mongo `mapstructure:"mongo"`
}

type Config struct {
	Server        Server        `mapstructure:"server"`
	UserCenter    UserCenter    `mapstructure:"user-center"`
	BlockchainAPI BlockchainAPI `mapstructure:"blockchain-api"`
	DB            DB            `mapstructure:"db"`
}

var GlobalConfig *Config

func Init() {
	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	buf, err := json.MarshalIndent(&c, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	GlobalConfig = &c
}
