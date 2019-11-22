package log

type Options struct {
	Project  string          `json:"project" default:"noname"`
	Env      string          `json:"env" default:"noname"`
	Hostname string          `json:"hostname"`
	Topics   []string        `json:"topics" default:"[\"-debug\"]"`
	Console  *ConsoleOptions `json:"console" default:"{\"enabled\":true,\"topics\":[\"-debug\"]}"`
	File     *FileOptions    `json:"file" default:"{\"enabled\":true,\"dir\":\"log\",\"topics\":[\"-debug\"]}"`
}

type ConsoleOptions struct {
	Enabled bool     `json:"enabled"`
	Topics  []string `json:"topics"`
}

type FileOptions struct {
	Enabled bool     `json:"enabled"`
	Dir     string   `json:"dir"`
	Topics  []string `json:"topics"`
}
