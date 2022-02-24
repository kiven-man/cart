package common

import (
	"strconv"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
)

// GetConsulConfig 设置配置中心
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		// optionally specify consul address; default to localhost:8500
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		// optionally specify prefix; defaults to /micro/config
		consul.WithPrefix(prefix),
		// optionally strip the provided prefix from the keys, defaults to false
		consul.StripPrefix(true),
	)
	//	配置初始化
	config, err := config.NewConfig()
	if err != nil {
		return config, nil
	}
	//加载配置
	err = config.Load(consulSource)
	return config, err
}
