package log

import (
	"os"
)

var (
	activeProject     = "noname"
	activeEnv         = "noname"
	activeHostname, _ = os.Hostname()
	activeAdapter     = SimpleAdapter()

	MainScope = NewScope("main")
)

func Setup(opts Options) {
}
