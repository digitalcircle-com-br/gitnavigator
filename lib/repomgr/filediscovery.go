package repomgr

import (
	"gitnavigator/lib/config"
	"io/fs"
	"log"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

type GitDir struct {
	Id     string `json:"id"`
	Dir    string `json:"dir"`
	Status string `json:"status"`
	Err    string `json:"err"`
}

//type Node struct {
//	Dir      string
//	Children map[string]*Node
//}
//
//func NewNode(s string) *Node{
//	ret:=&Node{
//		Dir:      s,
//		Children: make(map[string]*Node),
//	}
//	return ret
//}
var dirs = make(map[string]*GitDir)

//var root = &Node{
//	Children: make(map[]),
//}

func RepoStatus(d string) (string, error) {
	cmd := exec.Command("git", "-C", d, "status")
	bs, err := cmd.CombinedOutput()
	return string(bs), err
}

func Get(d string) *GitDir {
	return dirs[d]
}

func List() []string {
	ret := make([]string, len(dirs))
	i := 0
	for k := range dirs {
		ret[i] = k
		i++
	}
	sort.Strings(ret)
	return ret
}

func Dirs() map[string]*GitDir {
	return dirs
}

func DirsArr() []*GitDir {
	ret := make([]*GitDir, len(dirs))
	i := 0
	for _, v := range dirs {
		ret[i] = v
		i++
	}
	return ret
}

func DiscoverRepos() {
	log.Printf("Running DiscoverRepos on: %s", config.Config.Root)
	r := config.Config.Root

	go filepath.Walk(r, func(p string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && strings.HasSuffix(p, ".git") {
			p = path.Dir(p)
			_, ok := dirs[p]
			if !ok {
				gd := &GitDir{
					Id:     p,
					Dir:    p,
					Status: "",
					Err:    "",
				}
				st, err := RepoStatus(p)
				if err != nil {
					gd.Err = err.Error()
				}
				gd.Status = st
				dirs[p] = gd
				log.Printf("Found repo: %s", p)
				//ret <- p
			}
		}
		return nil
	})

}
