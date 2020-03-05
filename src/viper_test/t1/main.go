package main

import (
	"github.com/spf13/viper"
	"fmt"
	"os"
)

func t1() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("D:\\github\\go\\src\\viper_test\\t1\\")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 设置大小写GO_WORKER_TEST1都行
	viper.SetEnvPrefix("go_worker")
	viper.BindEnv("test1")

	viper.SetDefault("test3", "test3 default")
	fmt.Println(viper.Get("test1"))
	fmt.Println(viper.Get("test2.name"))
	fmt.Println(viper.Get("test3"))

	os.Setenv("GO_WORKER_TEST1", "13")
	fmt.Println(viper.Get("test1"))
}
//hah
//lala
//test3 default
//13

func bindEnv(name, defaultValue string, v *string) {
	viper.BindEnv(name)
	viper.SetDefault(name, defaultValue)
	*v = viper.Get(name).(string)
}

func t2() {
	var value string
	bindEnv("t1", "dt1", &value)
	fmt.Println(value)
}

func main() {
	t2()
}