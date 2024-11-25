package configs

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"sync"
)

var f = flag.String("f", "etc/config.yaml", "config file")

var sy sync.Once

func Once() {
	sy.Do(func() {
		flag.Parse()
		var c = new(Config)
		conf.MustLoad(*f, c)
		Init(c)

		config = c
	})
}
