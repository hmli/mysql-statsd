package mysql_statsd

type Config struct {
	Mysql string
	Statsd string
	Interval int64
	Metrics []string
}

func (c *Config) SetDefault() {
	if c.Interval == 0 {
		c.Interval = 60
	}
}