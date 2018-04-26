package mysql_statsd

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"io"
	"github.com/smira/go-statsd"
)


type Option func(c *Collector)

func OptSetDatabase(dbPath string) Option {
	return func(c *Collector) {
		db , err:= xorm.NewEngine("mysql", dbPath)
		if err != nil {
			panic(err)
		}
		c.db = db
	}
}

func OptSetMetrics(metrics []string) Option {
	return func(c *Collector) {
		c.setMetrics(metrics)
	}
}


func OptSetLogger(w io.Writer) Option {
	return func(c *Collector) {
		c.logger = log.New(w, "statsd-mysql",  log.Ldate | log.Ltime | log.Llongfile)
	}
}

func OptSetStatsd(client *statsd.Client) Option {
	return func(c *Collector) {
		c.statsd = client
	}
}

func OptSetInterval(interval int64) Option {
	return func(c *Collector) {
		c.Interval = interval
	}
}
