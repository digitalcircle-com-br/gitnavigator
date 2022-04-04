package action

import (
	"fmt"
	"gitnavigator/lib/config"
	"gitnavigator/lib/types"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gen2brain/beeep"
	"github.com/skratchdot/open-golang/open"
)

func OpenGUI() {
	open.Run("http://localhost" + config.Config.Addr)
}

func OpenFolder(d string) {
	cmd := exec.Command("open", d)
	bs, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf(err.Error())
	} else {
		log.Printf("%s", string(bs))
	}
}

func OpenCode(d string) {
	cmd := exec.Command("code", d)
	bs, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf(err.Error())
	} else {
		log.Printf("%s", string(bs))
	}
}
func OpenGoland(d string) {
	cmd := exec.Command("open", "-na", "GoLand.app", d)
	bs, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf(err.Error())
	} else {
		log.Printf("%s", string(bs))
	}
}

func OpenNavigator(d string) {
	bs, err := os.ReadFile(path.Join(d, ".git", "config"))
	if err != nil {
		log.Printf("Error: %s ", err.Error())
	}
	str := string(bs)
	rx := regexp.MustCompile(`url = (?P<URL>.*)\n`)

	url := rx.FindStringSubmatch(str)[1]

	urlparts := strings.Split(url, "@")

	url = urlparts[1]

	url = strings.Replace(url, ":", "/", 1)

	url = strings.Replace(url, ".git", "", 1)

	url = fmt.Sprintf("https://%s", url)

	cmd := exec.Command("open", url)

	cmd.Run()
}

func OpenLog(d string) {
	cmd := exec.Command("open", "http://localhost:9999/git/log?path="+d)

	cmd.Run()
}
func OpenStatus(d string) {
	cmd := exec.Command("open", "http://localhost:9999/git/status?path="+d)

	cmd.Run()
}

func OpenTerminal(d string) {
	cmd := exec.Command("open", "-a", "terminal", d)
	bs, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf(err.Error())
	} else {
		log.Printf("%s", string(bs))
	}
}

func ExecGlobalCmd(r *types.ReqCmd) ([]byte, error) {
	for _, v := range config.Config.GlobalCmds {
		if v.Id == r.Id {
			return Cmd(v.Cmd)
		}
	}
	return []byte("Cmd Not Found"), nil
}

func ExecRepoCmd(r *types.ReqCmd) ([]byte, error) {
	var toexe []string
	for _, v := range config.Config.RepoCmds {
		if v.Id == r.Id {
			toexe = make([]string, len(v.Cmd))
			copy(toexe, v.Cmd)
			for i, v := range toexe {
				if v == "$repo" {
					toexe[i] = r.Repo
				}
			}
			return Cmd(toexe)
		}
	}
	return []byte("Cmd Not Found"), nil
}

func Notify(s string, p ...interface{}) {
	beeep.Alert("Git Navigator", fmt.Sprintf(s, p...), "")
}

func Cmd(s []string) ([]byte, error) {
	cmd := exec.Command(s[0], s[1:]...)
	return cmd.CombinedOutput()
}

func GitBranch(s string) ([]byte, error) {
	return Cmd([]string{"git", "-C", s, "rev-parse", "--abbrev-ref", "HEAD"})
}

func GitStatusShort(s string) ([]byte, error) {
	return Cmd([]string{"git", "-C", s, "status", "--short"})
}

func Du(s string) int64 {
	total := int64(0)
	filepath.Walk(s, func(path string, info fs.FileInfo, err error) error {
		switch {
		case err != nil:
			return err
		case info.IsDir():
			return nil
		default:
			total = total + info.Size()
			return nil
		}

	})
	return total
}

func NoFiles(s string) int64 {
	total := int64(0)
	filepath.Walk(s, func(path string, info fs.FileInfo, err error) error {
		switch {
		case err != nil:
			return err
		case info.IsDir():
			return nil
		default:
			total = total + 1
			return nil
		}

	})
	return total
}

func ConfigReload() {
	fname, err := config.Reload()
	if err != nil {
		Notify(err.Error())
	} else {
		Notify("Reloaded: %s", fname)
	}
}
