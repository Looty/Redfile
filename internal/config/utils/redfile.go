package config

var f Redfile

type Task struct {
	Description string            `yaml:"description,omitempty"`
	Cmds        []string          `yaml:"cmds"`
	Env         map[string]string `yaml:"env"`
}

type Redfile struct {
	Version float32           `yaml:"version"`
	Env     map[string]string `yaml:"env"`
	Tasks   map[string]Task   `yaml:"tasks"`
}
