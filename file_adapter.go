package log

type fileAdapter struct {
	dir     string
	filters Filters
}

func (f *fileAdapter) Log(e Event) error {
	// TODO:
	return nil
}

func NewFilterAdapter(dir string, filters Filters) Adapter {
	return &fileAdapter{
		dir:     dir,
		filters: filters,
	}
}
