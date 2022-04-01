package repomgr

import (
	"fmt"
	"gitnavigator/lib/action"
	"gitnavigator/lib/config"
	"io/fs"
	"log"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

type GitDir struct {
	Id      string `json:"id"`
	Dir     string `json:"dir"`
	Status  string `json:"status"`
	Err     string `json:"err"`
	Size    int64  `json:"size"`
	Files   int64  `json:"files"`
	Branch  string `json:"branch"`
	Pending bool   `json:"pending"`
}

var dirs = make(map[string]*GitDir)

func Get(d string) *GitDir {
	return dirs[d]
}

func NErr(s string, p ...interface{}) error {
	return fmt.Errorf(s, p...)
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

var inDisover = false

func DiscoverRepos() bool {
	if inDisover {
		return false
	}
	inDisover = true
	defer func() {
		inDisover = false
	}()
	log.Printf("Running DiscoverRepos on: %s", config.Config.Root)
	r := config.Config.Root

	go func() {
		actDir := ""
		err := filepath.Walk(r, func(p string, info fs.FileInfo, err error) error {
			actDir = p
			if err != nil {
				return err
			}
			if info.IsDir() && strings.HasSuffix(p, ".git") {
				p = path.Dir(p)

				branch, err := action.GitBranch(p)
				if err != nil {
					branch = []byte("No Branch")
				}

				status, err := action.GitStatusShort(p)
				if err != nil {
					return NErr("Cant read git status of %s: %s", p, err.Error())
				}

				pending := len(status) > 1

				_, ok := dirs[p]
				if !ok {
					gd := &GitDir{
						Id:      p,
						Dir:     p,
						Status:  string(status),
						Err:     "",
						Size:    action.Du(p),
						Files:    action.NoFiles(p),
						Branch:  string(branch),
						Pending: pending,
					}

					dirs[p] = gd
					log.Printf("Found repo: %s", p)
					//ret <- p
				}
			}

			return nil
		})
		if err != nil {
			log.Printf("Error checking: %s: %s", actDir, err.Error())
			action.Notify("Error checking: %s: %s", actDir, err.Error())
		}
		action.Notify("Finished reading repos. Found %v repos", len(dirs))
	}()

	return true

}
