package types

type CmdCfg struct {
	Id    string   `json:"id"`
	Label string   `json:"label"`
	Cmd   []string `json:"cmd"`
	Os    string   `json:"os"`
}

type Config struct {
	Addr       string   `yaml:"addr"`
	Root       string   `yaml:"root"`
	Favorites  []string `yaml:"favorites"`
	GlobalCmds []CmdCfg `yaml:"global_cmds"`
	RepoCmds   []CmdCfg `yaml:"repo_cmds"`
}
