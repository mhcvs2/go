package config

import (
	"github.com/Sirupsen/logrus"
	"io/ioutil"
	"github.com/ghodss/yaml"
)

var Config []Task

type Task struct {
	Name string `json:"name"`
	Src GitInfo `json:"src"`
	Dest GitInfo `json:"dest"`
	Clean bool `json:"clean"`
}

type GitInfo struct {
	Dir string `json:"dir"`
}

func Load(paths []string) {
	var config = []byte{}
	var err error
	for _, path := range paths {
		if path != ""{
			config, err = ioutil.ReadFile(path)
			if err != nil{
				continue
			}else {
				logrus.Info("Load: "+path)
				break
			}
		}
	}
	if len(config) == 0 {
		logrus.Fatal("Fail to load config")
	}

	err = yaml.Unmarshal(config, &Config)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	if len(Config) == 0 {
		logrus.Fatal("Nothing to do")
	}
	for _, task := range Config {
		logrus.Debug(task.Name)
	}
}
