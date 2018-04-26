package main

import (
	"flag"
	"gopkg.in/yaml.v2"
	"os"
	"io/ioutil"
	ms "mysql_statsd"
	 "github.com/smira/go-statsd"
	"fmt"
)

func main() {
	conf := flag.String("c", "conf_sample.yaml", "config file path")
	flag.Parse()
	file, err := os.Open(*conf)
	if err != nil  {
		panic(err)
	}
	b, err := ioutil.ReadAll(file)
	if err != nil  {
		panic(err)
	}
	config := new(ms.Config)
	err = yaml.Unmarshal(b, config)
	if err != nil  {
		panic(err)
	}
	fmt.Printf("config: %+v \n", config)
	config.SetDefault()
	statsdClient := statsd.NewClient(config.Statsd)

	collector := new(ms.Collector)
	collector.Setting(
		ms.OptSetDatabase(config.Mysql),
		ms.OptSetLogger(os.Stdout),
		ms.OptSetMetrics(config.Metrics),
		ms.OptSetStatsd(statsdClient),
		ms.OptSetInterval(config.Interval),
	)
	collector.Start()



}
