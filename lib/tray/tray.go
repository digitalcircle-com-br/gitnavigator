package tray

import (
	"github.com/getlantern/systray"
	"gitnavigator/lib/action"
	"runtime"
)

//var root = ""
//var repoMenu *systray.MenuItem
//var menuMap = make(map[string]*systray.MenuItem)
//var favsMap = make(map[string]*systray.MenuItem)

func Run() {
	systray.Run(func() {
		SetTitle("Git Navigator")
		if runtime.GOOS == "windows" {
			systray.SetIcon(ico_win)
		} else {
			systray.SetIcon(ico_notwin)
		}

		open := systray.AddMenuItem("Open", "Open GUI")
		refresh := systray.AddMenuItem("Refresh", "Reloads Config")
		quit := systray.AddMenuItem("Quit", "Quit the whole app")
		go func() {
			for {
				select {
				case <-open.ClickedCh:
					action.OpenGUI()
				case <-refresh.ClickedCh:
					action.ConfigReload()
				case <-quit.ClickedCh:
					systray.Quit()
				}

			}
		}()
	}, func() {

	})
}

var menus = make(map[string]*systray.MenuItem)

func SetTitle(s string) {
	systray.SetTitle(s)
}
func Stop() {
	systray.Quit()
}

//func SetRoot(s string) {
//	root = s
//}

//func GetRoot() string {
//	return root
//}

//func AddRepo(s string) {
//	if repoMenu == nil {
//
//	}
//	p := s
//	var parent *systray.MenuItem
//	var ok = false
//	dirstack := make([]string, 0)
//	for {
//		p = path.Dir(p)
//		parent, ok = menuMap[p]
//		if !ok {
//			dirstack = append(dirstack, p)
//			if p == root {
//				parent = repoMenu
//				break
//			} else {
//				continue
//			}
//		} else {
//			break
//		}
//	}
//
//	for i := len(dirstack) - 1; i >= 0; i-- {
//		log.Printf("Adding item: %s", path.Base(dirstack[i]))
//		m := parent.AddSubMenuItem(path.Base(dirstack[i]), "")
//		menuMap[dirstack[i]] = m
//		parent = m
//	}
//	amenu := parent.AddSubMenuItem(path.Base(s), "")
//	menus[s] = amenu
//	openFolder := amenu.AddSubMenuItem("Open Folder", "")
//	openVS := amenu.AddSubMenuItem("Open VSCode", "")
//	openGoland := amenu.AddSubMenuItem("Open Goland", "")
//	terminal := amenu.AddSubMenuItem("Terminal", "")
//	web := amenu.AddSubMenuItem("Web", "")
//	log := amenu.AddSubMenuItem("Git - Logs", "")
//	gitstatus := amenu.AddSubMenuItem("Git - Status", "")
//	amenu.AddSubMenuItem("Add To Favorites", "")
//
//	go func() {
//		for {
//			select {
//			case <-openFolder.ClickedCh:
//				action.OpenFolder(s)
//			case <-openVS.ClickedCh:
//				action.OpenCode(s)
//			case <-openGoland.ClickedCh:
//				action.OpenGoland(s)
//			case <-web.ClickedCh:
//				action.OpenNavigator(s)
//			case <-terminal.ClickedCh:
//				action.OpenTerminal(s)
//			case <-log.ClickedCh:
//				action.OpenLog(s)
//			case <-gitstatus.ClickedCh:
//				action.OpenStatus(s)
//			}
//		}
//	}()
//}

//func AddFav(s string) {
//	favMenu := systray.AddMenuItem(s, "")
//	favsMap[s] = favMenu
//}
//
//func RemFav(s string) {
//	fm, ok := favsMap[s]
//	if !ok {
//		return
//	}
//}
