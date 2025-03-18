package internal

import "flag"

type Config struct {
	attribute string
	interval  int
	workers   int
}

func NewConfig() *Config {
	attribute := flag.String("attribute", "foo", "Log attribute")
	interval := flag.Int("interval", 1000, "Log interval")
	workers := flag.Int("workers", 10, "Log workers")

	flag.Parse()

	cfg := &Config{
		attribute: *attribute,
		interval:  *interval,
		workers:   *workers,
	}

	return cfg
}

func (c *Config) Attribute() string {
	return c.attribute
}

func (c *Config) Interval() int {
	return c.interval
}

func (c *Config) Workers() int {
	return c.workers
}
