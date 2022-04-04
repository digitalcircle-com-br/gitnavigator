import * as PageRepos from "./pgrepos.js"

let lastcenterid = ""

export function ReplaceTitle(nt: string) {
    webix.ui({id: "lbTitle", view: "label", label: nt, css: {color: "#2d2d2d !important"}}, $$("lbTitle"))
}

export function ReplaceCenter(nl: any) {
    let old = $$(lastcenterid)
    if (old) {
        webix.ui(nl, old)
        lastcenterid = nl.id
    }
}

let globalRefresh = () => {
}

export function SetGlobalRefresh(a: () => void) {
    globalRefresh = a
}

async function main() {
    //await api.init()


    let sidebarConfig = {
        id: null,
        view: "sidebar",
        scroll: "auto",
        data: [
            {id: "repos", value: "Repositories", icon: "fas fa-cubes"}
        ],
        on: {
            onItemClick: (id: string) => {
                if (id != lastclickedmenu) {
                    let h = pageController[id]
                    if (h != null) {
                        h()
                    }
                }
            }
        }

    }
    let lastclickedmenu = ""

    let pageController: { [s: string]: () => void } = {
        repos: PageRepos.Show
    }

    let me: string = <string>(localStorage.getItem("user") ? localStorage.getItem("user") : "N/A")

    let toolbar = {
        view: "toolbar",
        height: 50,
        background: "white",
        borderless: true,
        css: {
            background: "white"
        },
        elements: [
            // {
            //     view: "icon", icon: "fas fa-bars", css: {color: "#2d2d2d !important"}, click: () => {
            //         let sb = <webix.ui.sidebar>$$(<string><unknown>sidebarConfig.id)
            //         sb.toggle()
            //     }
            // },
            {
                template: "<img src='resources/logo_512.png' style='height: 40px;width: 40px'></img>",
                width: 58,
                borderless: true
            },

            {view: "label", label: "Git Navigator ★", width: 120, css: {color: "#2d2d2d !important"}},
            {view: "label", id: "lbTitle"},
            {},
            {
                view: "icon", icon: "fas fa-sync",tip:"Refreshes Repos List", click: () => {
                    globalRefresh()
                }
            },
            // {view: "label", id: "lbMe", label: me, width: me.length * 8},
            // {
            //     view: "icon", icon: "fas fa-user", click: () => {
            //         webix.message("implement me")
            //     }
            // }
        ]
    }
    let tmpl = {template: "", id: "center"}
    let layout = {
        rows: [
            toolbar,
            {cols: [/*sidebarConfig, */ tmpl]}

        ]
    }

    function getAnchor(): string {
        var currentUrl = document.URL,
            urlParts = currentUrl.split('#');

        return (urlParts.length > 1) ? urlParts[1] : "";
    }

    webix.ui(layout)

    lastcenterid = tmpl.id

    let anchor = getAnchor()
    if (pageController[anchor]) {
        pageController[anchor]()
    } else {
        PageRepos.Show()
        // webix.message("implement me")
    }


}

main()

