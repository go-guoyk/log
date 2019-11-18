package autoconf

import (
	"github.com/novakit/conf"
	"github.com/novakit/log"
)

func init() {
	opts := log.Options{}
	conf.RegisterLoader(&conf.Loader{
		Name:   "log",
		Target: &opts,
		Loaded: func() {
			log.Setup(opts)
		},
	})
}
