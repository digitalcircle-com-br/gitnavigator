package config

import (
	"fmt"
	"gitnavigator/lib/types"
	"os"
	"os/user"
	"path"
	"runtime"
	"strings"

	"gopkg.in/yaml.v3"
)

var Config = &types.Config{
	Addr:       ":19999",
	Favorites:  []string{},
	GlobalCmds: []types.CmdCfg{},
	RepoCmds:   []types.CmdCfg{},
	Root:       "",
}

func enc() ([]byte, error) {
	return yaml.Marshal(Config)
}
func dec(bs []byte) error {
	return yaml.Unmarshal(bs, Config)
}

func GlobalCmdS() []types.CmdCfg {
	ret := make([]types.CmdCfg, 0)
	thisOS := runtime.GOOS
	for _, v := range Config.GlobalCmds {
		if v.Os == "" || v.Os == thisOS {
			ret = append(ret, v)
		}
	}
	return ret
}

func RepoCmdS() []types.CmdCfg {
	ret := make([]types.CmdCfg, 0)
	thisOS := runtime.GOOS
	for _, v := range Config.RepoCmds {
		if v.Os == "" || v.Os == thisOS {
			ret = append(ret, v)
		}
	}
	return ret
}

var cfgName = ""

//Load will load config from ~/.gitnavigator.yaml.
func Load(fname string) error {
	cfgName = fname
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
func Reload() (string, error) {
	return cfgName, Load(cfgName)
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
