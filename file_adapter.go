package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type fileAdapter struct {
	dir        string
	filter     *Filter
	files      map[string]*os.File
	filesMutex *sync.RWMutex
}

func NewFilterAdapter(dir string, filter *Filter) Adapter {
	f := &fileAdapter{
		dir:        dir,
		filter:     filter,
		files:      map[string]*os.File{},
		filesMutex: &sync.RWMutex{},
	}
	return f
}

func (fa *fileAdapter) retrieveFile(filename string) (f *os.File, err error) {
	fa.filesMutex.RLock()
	f = fa.files[filename]
	fa.filesMutex.RUnlock()

	if f != nil {
		return
	}

	fa.filesMutex.Lock()
	defer fa.filesMutex.Unlock()

	if f = fa.files[filename]; f != nil {
		return
	}
	f, err = os.OpenFile(filepath.Join(fa.dir, filename), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if f != nil {
		fa.files[filename] = f
	}
	return
}

func (fa *fileAdapter) Log(e Event) (err error) {
	if !fa.filter.IsTopicEnabled(e.Topic) {
		return
	}

	content := &bytes.Buffer{}
	content.WriteString(e.Timestamp.Format("2006-01-02T15:04:05.000-0700"))
	content.WriteByte(' ')
	if len(e.Labels) == 0 {
		content.Write([]byte("{}"))
	} else {
		if s, err := json.Marshal(e.Labels); err != nil {
			return err
		} else {
			content.Write(s)
		}
	}
	if len(e.Message) > 0 {
		content.WriteByte(' ')
		content.WriteString(strings.TrimSpace(e.Message))
	}
	content.WriteByte('\n')

	filename := fmt.Sprintf(
		"%s:%s:%s:%s.%04d-%02d-%02d.log",
		e.Project, e.Env, e.Hostname, e.Topic,
		e.Timestamp.Year(), e.Timestamp.Month(), e.Timestamp.Day(),
	)
	var f *os.File
	if f, err = fa.retrieveFile(filename); err != nil {
		return
	}

	// assuming atomic write
	if _, err = f.Write(content.Bytes()); err != nil {
		return
	}
	return
}

func (fa *fileAdapter) Close() error {
	oldFiles := fa.files
	fa.files = map[string]*os.File{}
	for _, a := range oldFiles {
		_ = a.Close()
	}
	return nil
}
