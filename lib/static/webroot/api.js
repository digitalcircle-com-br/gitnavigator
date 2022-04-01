class APIProxy {
    constructor() {
        this.endpoint = "/api";
        this.onerror = (err) => {
            console.error(err);
        };
    }
    async invokeThrow(method, url, body) {
        let nurl = this.endpoint + url;
        var myHeaders = new Headers();
        var init = {
            method: method,
            headers: myHeaders,
            mode: 'cors',
            credentials: 'include',
            body: JSON.stringify(body),
            cache: 'no-cache'
        };
        if (method == "GET" || method == "HEAD") {
            delete init.body;
        }
        let res = await fetch(nurl, init);
        if (res.status >= 400) {
            let msg = await res.text();
            if (msg == "" || msg == null) {
                msg = "Erro ao invocar url:" +
                    method + " " + nurl + " => " +
                    res.status + ":" +
                    res.statusText;
            }
            return Promise.reject(msg);
        }
        return await res.text();
    }
    async invoke(method, url, body) {
        let nurl = this.endpoint + url;
        var myHeaders = new Headers();
        var init = {
            method: method,
            headers: myHeaders,
            mode: 'cors',
            credentials: 'include',
            body: JSON.stringify(body),
            cache: 'default'
        };
        if (method == "GET" || method == "HEAD") {
            delete init.body;
        }
        let res = await fetch(nurl, init);
        if (res.status >= 400) {
            let msg = await res.text();
            if (msg == "" || msg == null) {
                msg = "Erro ao invocar url:" +
                    method + " " + url + " => " +
                    res.status + ":" +
                    res.statusText;
            }
            this.onerror(msg);
        }
        return await res.text();
    }
    async invokePlain(method, url, body) {
        let nurl = this.endpoint + url;
        var myHeaders = new Headers();
        var init = {
            method: method,
            headers: myHeaders,
            mode: 'cors',
            body: body,
            cache: 'default'
        };
        if (method == "GET" || method == "HEAD") {
            delete init.body;
        }
        let res = await fetch(nurl, init);
        if (res.status >= 400) {
            let msg = await res.text();
            if (msg == "" || msg == null) {
                msg = "Erro ao invocar url:" +
                    method + " " + url + " => " +
                    res.status + ":" +
                    res.statusText;
            }
            this.onerror(msg);
        }
        return await res.text();
    }
    async json(method, url, body) {
        let txt = await this.invoke(method, url, body);
        if (!txt || txt == "") {
            txt = "{}";
        }
        return JSON.parse(txt);
    }
    async jsonGet(url) {
        let txt = await this.invoke("GET", url, null);
        if (!txt || txt == "") {
            txt = "{}";
        }
        return JSON.parse(txt);
    }
    async jsonPost(url, body) {
        let txt = await this.invoke("POST", url, body);
        if (!txt || txt == "") {
            txt = "{}";
        }
        return JSON.parse(txt);
    }
    async jsonPut(method, url, body) {
        let txt = await this.invoke("PUT", url, body);
        if (!txt || txt == "") {
            txt = "{}";
        }
        return JSON.parse(txt);
    }
    async jsonDelete(method, url) {
        let txt = await this.invoke("DELETE", url, null);
        if (!txt || txt == "") {
            txt = "{}";
        }
        return JSON.parse(txt);
    }
}
class API {
    constructor() {
        this.proxy = new APIProxy();
    }
    async GitRepos() {
        return await this.proxy.jsonGet("/v1/dirs");
    }
    async GitDiscover() {
        return await this.proxy.jsonGet("/v1/dirs");
    }
    async CmdsGlobal() {
        return await this.proxy.jsonGet("/v1/cmds/global");
    }
    async CmdsRepo() {
        return await this.proxy.jsonGet("/v1/cmds/repo");
    }
    async CmdExecGlobal(r) {
        return await this.proxy.jsonPost("/v1/cmd/global", r);
    }
    async CmdExecRepo(r) {
        return await this.proxy.jsonPost("/v1/cmd/repo", r);
    }
    log(i) {
        console.log(i);
    }
}
let api = new API();
export default api;
