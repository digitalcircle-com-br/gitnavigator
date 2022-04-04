package app

import (
	"flag"
	"gitnavigator/lib/action"
	"gitnavigator/lib/api"
	"gitnavigator/lib/config"
	"gitnavigator/lib/repomgr"
	"gitnavigator/lib/tray"
	"log"
)

func parseCmd() {
	flag.String("root", "", "Root Dir(for dev)")
}

func Run() {
	hErr(func() error {
		action.Setup()
		cfg := flag.String("c", "", "Optional config file - defaults to [~/.gitnavigator.yaml]")
		flag.Parse()
		err := config.Load(*cfg)
		if err != nil {
			return err
		}
		api.Start()
		go repomgr.DiscoverRepos()
		tray.Run()

		return nil
	})
}
func hErr(h func() error) {
	err := h()
	if err != nil {
		log.Printf(err.Error())
	}
}

func OpenGUI() {
	action.OpenGUI(nil)
}

func AddFavorite(d string) {
	hErr(func() error {
		config.Config.Favorites = append(config.Config.Favorites, d)
		return config.Save()
	})
}
func RemoveFavorite(d string) {
	hErr(func() error {
		nfav := make([]string, len(config.Config.Favorites), 0)
		for _, v := range config.Config.Favorites {
			if v != d {
				nfav = append(nfav, v)
			}
		}
		config.Config.Favorites = nfav
		return config.Save()
	})
}
