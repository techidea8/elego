package initialize

import (
	"fmt"
	"turingapp/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//环境变量
func Viper(filepath []string) *viper.Viper {
	config := "./app.yaml"
	if len(filepath) > 0 {
		config = filepath[0]
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
		if err := v.Unmarshal(&global.Conf); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.Conf); err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%+v", global.Conf)
	return v
}

func InitConfig(path ...string) *viper.Viper {
	global.Viper = Viper(path)
	return global.Viper
}
