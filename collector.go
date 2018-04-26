package mysql_statsd

import (
	"github.com/go-xorm/xorm"
	"log"
	"github.com/smira/go-statsd"
	"strconv"
	"strings"
	"time"
)

var statusSQL = "SHOW GLOBAL STATUS"
var innodbSQL = "SHOW ENGINE INNODB STATUS"
var defaultPrefix = "mysql_perf"

type Collector struct {
	metrics []string
	Interval int64
	db *xorm.Engine
	logger *log.Logger
	statsd *statsd.Client
}

// collector
func (c *Collector) CollectStatus() (err error){
	rows, err := c.db.SQL(statusSQL).Rows(new(StatusData))
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		data := new(StatusData)
		rows.Scan(data)
		//c.logger.Printf("statsd collected: %+v", data)
		c.Record(data)
	}
	return nil
	//c.db.SQL(statusSQL).Rows()
}

func (c *Collector) Setting(opt ...Option) {
	for _, o := range opt {
		o(c)
	}
}

func (c *Collector) Record(status *StatusData) {
	if contains(c.metrics, status.Name) {
		c.record(status)
	}
}

func (c *Collector) record(status *StatusData) {
	//c.logger.Printf("status: %+v", status)
	value64, _ := strconv.ParseInt(status.Value, 10, 0)
	c.statsd.Gauge(status.Name, value64)
}

// setMetrics 过滤掉非数字的metrics
// 如果有非数字的， 直接panic 然后指明哪个有错
func (c *Collector) setMetrics(metrics []string) {
	c.metrics = make([]string, 0, len(metrics))
	for _, metric := range metrics {
		metric = strings.TrimSpace(metric)
		if !contains(filter_list, metric) {
			c.metrics = append(c.metrics, metric)
		}
	}
}

func (c *Collector) Start() {
	// TODO grace exit
	c.logger.Println(c.Interval)
	ticker := time.NewTicker(time.Duration(c.Interval) * time.Second)
	for {
		select {
		case <- ticker.C:
			c.CollectStatus()
		}
	}
}

type StatusData struct  {
	Name string `xorm:"Variable_name"`
	Value string `xorm:"Value"`
}

func contains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}


