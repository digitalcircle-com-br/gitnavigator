class APIProxy {
    endpoint = "/api"

    onerror = (err: any) => {
        console.error(err)
    }

    async invokeThrow(method: string, url: string, body: any): Promise<any> {
        let nurl = this.endpoint + url
        var myHeaders = new Headers();
        var init: RequestInit = {
            method: method,
            headers: myHeaders,
            mode: 'cors',
            credentials: 'include',
            body: JSON.stringify(body),
            cache: 'no-cache'
        };

        if (method == "GET" || method == "HEAD") {
            delete init.body
        }

        let res = await fetch(nurl, init)
        if (res.status >= 400) {
            let msg = await res.text()
            if (msg == "" || msg == null) {
                msg = "Erro ao invocar url:" +
                    method + " " + nurl + " => " +
                    res.status + ":" +
                    res.statusText
            }
            return Promise.reject(msg)
        }
        return await res.text()
    }

    async invoke(method: string, url: string, body: any) {
        let nurl = this.endpoint + url
        var myHeaders = new Headers();
        var init: RequestInit = {
            method: method,
            headers: myHeaders,
            mode: 'cors',
            credentials: 'include',
            body: JSON.stringify(body),
            cache: 'default'
        };

        if (method == "GET" || method == "HEAD") {
            delete init.body
        }

        let res = await fetch(nurl, init)
        if (res.status >= 400) {
            let msg = await res.text()
            if (msg == "" || msg == null) {
                msg = "Erro ao invocar url:" +
                    method + " " + url + " => " +
                    res.status + ":" +
                    res.statusText
            }
            this.onerror(msg)
        }

        return await res.text()
    }

    async invokePlain(method: string, url: string, body: string) {
        let nurl = this.endpoint + url
        var myHeaders = new Headers();
        var init: RequestInit = {
            method: method,
            headers: myHeaders,
            mode: 'cors',
            body: body,
            cache: 'default'
        };

        if (method == "GET" || method == "HEAD") {
            delete init.body
        }

        let res = await fetch(nurl, init)
        if (res.status >= 400) {
            let msg = await res.text()
            if (msg == "" || msg == null) {
                msg = "Erro ao invocar url:" +
                    method + " " + url + " => " +
                    res.status + ":" +
                    res.statusText
            }
            this.onerror(msg)
        }

        return await res.text()
    }

    async json(method: string, url: string, body: any) {
        let txt = await this.invoke(method, url, body)
        if (!txt || txt == "") {
            txt = "{}"
        }
        return JSON.parse(txt)
    }

    async jsonGet(url: string) {
        let txt = await this.invoke("GET", url, null)
        if (!txt || txt == "") {
            txt = "{}"
        }
        return JSON.parse(txt)
    }

    async jsonPost(url: string, body: any) {
        let txt = await this.invoke("POST", url, body)
        if (!txt || txt == "") {
            txt = "{}"
        }
        return JSON.parse(txt)
    }

    async jsonPut(method: string, url: string, body: any) {
        let txt = await this.invoke("PUT", url, body)
        if (!txt || txt == "") {
            txt = "{}"
        }
        return JSON.parse(txt)
    }

    async jsonDelete(method: string, url: string) {
        let txt = await this.invoke("DELETE", url, null)
        if (!txt || txt == "") {
            txt = "{}"
        }
        return JSON.parse(txt)
    }
}

class API {
    proxy = new APIProxy()

    async GitRepos(): Promise<Repo[]> {
        return await this.proxy.jsonGet("/v1/dirs")
    }

    async GitDiscover(): Promise<string> {
        return await this.proxy.jsonGet("/v1/dirs")
    }

    async CmdsGlobal(): Promise<Cmd[]> {
        return await this.proxy.jsonGet("/v1/cmds/global")
    }

    async CmdsRepo(): Promise<Cmd[]> {
        return await this.proxy.jsonGet("/v1/cmds/repo")
    }

    async CmdExecGlobal(r: ReqCmd): Promise<ResCmd> {
        return await this.proxy.jsonPost("/v1/cmd/global", r)
    }
    
    async CmdExecRepo(r: ReqCmd): Promise<ResCmd> {
        return await this.proxy.jsonPost("/v1/cmd/repo", r)
    }

    log(i: any) {
        console.log(i)
    }
}

export interface Repo {
    dir: string
    status: string
    err: string
    size: number
    files: number
    branch: string
    pending: boolean
}

export interface Cmd {
    id: string
    label: string
}

export interface ReqCmd {
    id: string
    repo: string
    params: string[]
}

export interface ResCmd {
    out: string
    err: string
}


let api = new API()
export default api