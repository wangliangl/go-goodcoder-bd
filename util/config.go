package util

import (
	"gopkg.in/yaml.v2"
	"icode.baidu.com/baidu/go-lib/log"
	"os"
)

type Config struct {
	Spider SpiderConf `yaml:"spider"`
}

type SpiderConf struct {
	UrlListFile     string `yaml:"urlListFile"`
	OutputDirectory string `yaml:"outputDirectory"`
	MaxDepth        int    `yaml:"maxDepth"`
	CrawlInterval   int    `yaml:"crawlInterval"`
	CrawlTimeout    int    `yaml:"crawlTimeout"`
	TargetUrl       string `yaml:"targetUrl"`
	ThreadCount     int    `yaml:"threadCount"`
}

func LoadConfig(path string) (*Config, error) {
	conf := &Config{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		err = yaml.NewDecoder(f).Decode(conf)
		if err != nil {
			_ = log.Logger.Error("decode config err:%s", err.Error())
			return nil, err
		}
	}
	return conf, nil
}
