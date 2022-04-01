package action

import (
	"fmt"
	"gitnavigator/lib/config"
	"gitnavigator/lib/types"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

//func GitLog(d string) {
//	cmd := exec.Command("git", "-C", d, "--nopager", "log", "-1")
//	bs, err := cmd.CombinedOutput()
//
//}

func OpenGUI() {
	cmd := exec.Command("open", "http://localhost"+config.Config.Addr)
	bs, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf(err.Error())
	} else {
		log.Printf("%s", string(bs))
	}
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
	var toexe []string
	for _, v := range config.Config.GlobalCmds {
		if v.Id == r.Id {
			toexe = make([]string, len(v.Cmd))
			copy(toexe, v.Cmd)

			cmd := exec.Command(toexe[0], toexe[1:]...)
			return cmd.CombinedOutput()
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
			cmd := exec.Command(toexe[0], toexe[1:]...)
			return cmd.CombinedOutput()
		}
	}
	return []byte("Cmd Not Found"), nil
}
