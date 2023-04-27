package common

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func Viper(path ...string) {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(CfgEnv); configEnv == "" {
				config = DefaultCfg
				fmt.Printf("您正在使用config的默认值, 配置路径为%v\n", DefaultCfg)
			} else {
				config = configEnv
				fmt.Printf("您正在使用GVA_CONFIG环境变量, 配置路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值, 配置路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func UseViper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&Cfg); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&Cfg); err != nil {
		fmt.Println(err)
	}
	setupViper(v)
}
