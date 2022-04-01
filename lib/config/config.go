package config

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

type CmdCfg struct {
	Id    string   `json:"id"`
	Label string   `json:"label"`
	Cmd   []string `json:"cmd"`
}

type cfg struct {
	Addr       string   `yaml:"addr"`
	Root       string   `yaml:"root"`
	Favorites  []string `yaml:"favorites"`
	GlobalCmds []CmdCfg `yaml:"global_cmds"`
	RepoCmds   []CmdCfg `yaml:"repo_cmds"`
}

var Config = &cfg{
	Addr:       ":19999",
	Root:       "~/projects",
	Favorites:  []string{},
	GlobalCmds: []CmdCfg{},
	RepoCmds:   []CmdCfg{},
}

func enc() ([]byte, error) {
	return yaml.Marshal(Config)
}
func dec(bs []byte) error {
	return yaml.Unmarshal(bs, Config)
}

//Load will load config from ~/.gitnavigator.yaml.
func Load(fname string) error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	if fname == "" {

		fname = path.Join(usr.HomeDir, ".gitnavigator.yaml")
	}
	bs, err := os.ReadFile(fname)
	if err != nil {
		return Save()
	}
	if err != nil {
		return err
	}
	err = dec(bs)
	if err == nil {
		Config.Root = strings.ReplaceAll(Config.Root, "~", usr.HomeDir)
	}

	for i := range Config.GlobalCmds {
		Config.GlobalCmds[i].Id = fmt.Sprintf("g-%v", i)
	}

	for i := range Config.RepoCmds {
		Config.RepoCmds[i].Id = fmt.Sprintf("r-%v", i)

	}

	return err
}

//Save will save actual config to ~/.gitnavigator.yaml.
func Save() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	fname := path.Join(usr.HomeDir, ".gitnavigator.yaml")
	bs, err := enc()
	if err != nil {
		return err
	}
	err = os.WriteFile(fname, bs, 0600)

	return err
}
