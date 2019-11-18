package log

const DefaultScopeName = "default"

type Options struct {
	Project  string              `json:"project" default:"noname"`
	Env      string              `json:"env" default:"noname"`
	Hostname string              `json:"hostname"`
	Topics   map[string][]string `json:"topics" default:"{\"default\":[\"-debug\"]}"`
	Console  ConsoleOptions      `json:"console" default:"{\"enabled\":true,\"dir\":\"\",\"topics\":{\"default\":[\"-debug\"]}}"`
	File     FileOptions         `json:"file" default:"{\"enabled\":true,\"topics\":{\"default\":[\"-debug\"]}}"`
}

type ConsoleOptions struct {
	Enabled bool                `json:"enabled"`
	Topics  map[string][]string `json:"topics"`
}

type FileOptions struct {
	Enabled bool                `json:"enabled"`
	Dir     string              `json:"dir"`
	Topics  map[string][]string `json:"topics"`
}
