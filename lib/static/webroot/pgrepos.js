import { ReplaceCenter, ReplaceTitle, SetGlobalRefresh } from "./index.js";
import api from "./api.js";
export async function Show() {
    let gcmds = await api.CmdsGlobal();
    api.log(gcmds);
    let rcmds = await api.CmdsRepo();
    api.log(rcmds);
    let gTbCfg = {
        view: "toolbar",
        height: 40,
        elements: []
    };
    let rTbCfg = {
        view: "toolbar",
        height: 40,
        elements: []
    };
    if (gcmds) {
        for (let c of gcmds) {
            let btnbuilder = (c) => {
                return {
                    view: "button", label: c.label, click: async () => {
                        let res = await api.CmdExecGlobal({ id: c.id, params: [], repo: "" });
                        if (res.out && res.out != "") {
                            webix.alert(res.out);
                        }
                        if (res.err && res.err != "") {
                            webix.alert(res.err);
                        }
                    }
                };
            };
            let bt = btnbuilder(c);
            gTbCfg.elements?.push(bt);
        }
    }
    if (rcmds) {
        for (let c of rcmds) {
            let btnbuilder = (c) => {
                return {
                    view: "button", label: c.label, click: async () => {
                        let res = await api.CmdExecRepo({ id: c.id, params: [], repo: dt.getSelectedId(false, true) });
                        if (res.out && res.out != "") {
                            webix.alert(res.out);
                        }
                        if (res.err && res.err != "") {
                            webix.alert(res.err);
                        }
                    }
                };
            };
            let bt = btnbuilder(c);
            rTbCfg.elements?.push(bt);
        }
    }
    let tb = {
        view: "datatable",
        resizeColumn: { size: 10 },
        resizeRow: false,
        select: "row",
        //autoConfig: true
        columns: [
            { id: "dir", fillspace: 5, sort: "string" },
            { id: "size", fillspace: 1, sort: "int" },
            { id: "files", fillspace: 1, sort: "int" },
            { id: "branch", fillspace: 1, sort: "string" },
            { id: "pending", fillspace: 1, sort: "string" },
        ]
    };
    let txtFilterCfg = {
        view: "text",
        label: "Filter",
        on: {
            onEnter: () => {
                let txt = txtFilter.getValue();
                if (txt && txt.length > 1) {
                    dt.filter((o) => {
                        return o.dir.indexOf(txt) > -1;
                    });
                }
                else {
                    dt.filter("");
                }
            }
        }
    };
    let layout = {
        rows: []
    };
    //@ts-ignore
    // if (gTbCfg.elements?.length > 0) {
    //     //@ts-ignore
    //     layout.rows.push(gTbCfg)
    // }
    //@ts-ignore
    // if (rTbCfg.elements?.length > 0) {
    //     //@ts-ignore
    //     layout.rows.push(rTbCfg)
    // }
    //@ts-ignore
    layout.rows.push(txtFilterCfg, tb);
    ReplaceTitle("Repos");
    ReplaceCenter(layout);
    let dt = $$(tb.id);
    let data = await api.GitRepos();
    let txtFilter = $$(txtFilterCfg.id);
    dt.parse(data, "json");
    SetGlobalRefresh(async () => {
        let data = await api.GitRepos();
        let txtFilter = $$(txtFilterCfg.id);
        dt.parse(data, "json");
        webix.message(`Data reloaded - found ${data.length} repos`);
    });
    //@ts-ignore
    if (rTbCfg.elements.length > 1) {
        let ctxMenuCfg = {
            view: "context", id: "cm",
            body: {
                view: "toolbar",
                rows: rTbCfg.elements
            },
            master: $$(tb.id) // component object
        };
        webix.ui(ctxMenuCfg);
    }
}
