package log

import "os"

var (
	activeProject     = "noname"
	activeEnv         = "noname"
	activeHostname, _ = os.Hostname()
	activeFilters     = Filters{}
	activeAdapters    []Adapter
)

func Setup(opts Options) {
	if len(opts.Project) != 0 {
		activeProject = opts.Project
	}
	if len(opts.Env) != 0 {
		activeEnv = opts.Env
	}
	if len(opts.Hostname) != 0 {
		activeHostname = opts.Hostname
	}
	activeFilters = NewFilters(opts.Topics)
	var adapters []Adapter
	if opts.Console.Enabled {
		adapters = append(adapters, NewConsoleAdapter(os.Stdout, NewFilters(opts.Console.Topics)))
	}
	if opts.File.Enabled {
		adapters = append(adapters, NewFilterAdapter(opts.File.Dir, NewFilters(opts.File.Topics)))
	}
	activeAdapters = adapters
}
