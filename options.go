package log

type Options struct {
	Project  string              `json:"project" default:"noname"`
	Env      string              `json:"env" default:"noname"`
	Hostname string              `json:"hostname"`
	Topics   map[string][]string `json:"topics" default:"{\"default\":[\"-debug\"]}"`
	Console  ConsoleOptions      `json:"console"`
	File     FileOptions         `json:"file"`
}

type ConsoleOptions struct {
	Disabled bool                `json:"disabled"`
	Topics   map[string][]string `json:"topics" default:"{\"default\":[\"-debug\"]}"`
}

type FileOptions struct {
	Disabled bool                `json:"disabled"`
	Dir      string              `json:"dir" default:"log"`
	Topics   map[string][]string `json:"topics" default:"{\"default\":[\"-debug\"]}"`
}
