package main

import (
	"flag"
	"fmt"
	"gateway/configs"
	"gateway/internal/server"
	"github.com/spf13/viper"
	"path"
)

var (
	WebConfigPath = "configs/conf.yml"
	version       = "__BUILD_VERSION_"
	execDir       string
	provider      string
	st, v, V      bool
)

func main() {
	flag.StringVar(&execDir, "d", ".", "项目目录")
	flag.StringVar(&provider, "p", "consul", "项目配置提供者")
	flag.BoolVar(&v, "v", false, "查看版本号")
	flag.BoolVar(&V, "V", false, "查看版本号")
	flag.BoolVar(&st, "s", false, "项目状态")
	flag.Parse()
	if v || V {
		fmt.Println(version)
	}
	// init配置
	sc, err := initConfigFromConsul()
	must(err)
	server.NewServer(sc).Run()
}

func initConfigFromConsul() (*configs.ServerConfig, error) {
	sc := &configs.ServerConfig{}
	v := viper.New()
	v.SetConfigFile(path.Join(execDir, WebConfigPath))
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := v.Unmarshal(sc); err != nil {
		return nil, err
	}

	err := v.AddRemoteProvider(provider,
		fmt.Sprintf("%s:%d", sc.Consul.Host, sc.Consul.Port),
		sc.ConfigPath)
	if err != nil {
		return nil, err
	}

	v.SetConfigType("YAML")

	if err = v.ReadRemoteConfig(); err != nil {
		return nil, err
	}

	if err = v.Unmarshal(sc); err != nil {
		return nil, err
	}
	return sc, nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
