package api

import (
	"encoding/json"
	"gitnavigator/lib/action"
	"gitnavigator/lib/config"
	"gitnavigator/lib/repomgr"
	"gitnavigator/lib/static"
	"gitnavigator/lib/types"
	"log"
	"net/http"
	"os/exec"
	"path"

	"github.com/gorilla/mux"
)

func dec(i interface{}, r *http.Request) {
	json.NewDecoder(r.Body).Decode(i)
	r.Body.Close()

}

func enc(i interface{}, w http.ResponseWriter) {
	w.Header().Add("content-type", "application/json")
	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		log.Printf("Error encoding: %s", err.Error())
	}

}

func Start() {
	r := mux.NewRouter()

	r.Methods(http.MethodGet).Path("/git/log").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		gitpath := request.URL.Query().Get("path")
		cmd := exec.Command("git", "-C", path.Join(gitpath), `--no-pager`, "log")
		bs, err := cmd.CombinedOutput()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			writer.Write([]byte("\n"))
			writer.Write([]byte(bs))
			return
		}
		writer.Write(bs)
	})

	r.Methods(http.MethodGet).Path("/git/status").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		gitpath := request.URL.Query().Get("path")
		cmd := exec.Command("git", "-C", path.Join(gitpath), `status`)
		bs, err := cmd.CombinedOutput()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			writer.Write([]byte("\n"))
			writer.Write([]byte(bs))
			return
		}
		writer.Write(bs)
	})

	r.Methods(http.MethodGet).Path("/api/v1/dirs").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		dirs := repomgr.DirsArr()
		enc(dirs, writer)
	})

	r.Methods(http.MethodGet).Path("/api/v1/dirs/discover").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		go repomgr.DiscoverRepos()
		enc("ok", writer)
	})

	r.Methods(http.MethodGet).Path("/api/v1/cmds/global").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		enc(config.Config.GlobalCmds, writer)
	})

	r.Methods(http.MethodGet).Path("/api/v1/cmds/repo").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		enc(config.Config.RepoCmds, writer)
	})

	r.Methods(http.MethodPost).Path("/api/v1/cmd/global").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		in := &types.ReqCmd{}
		dec(in, request)
		bs, err := action.ExecGlobalCmd(in)
		out := &types.ResCmd{}
		out.Out = string(bs)
		if err != nil {
			out.Err = err.Error()
		}
		enc(out, writer)
	})

	r.Methods(http.MethodPost).Path("/api/v1/cmd/repo").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		in := &types.ReqCmd{}
		dec(in, request)
		bs, err := action.ExecRepoCmd(in)
		out := &types.ResCmd{}
		out.Out = string(bs)
		if err != nil {
			out.Err = err.Error()
		}
		enc(out, writer)
	})

	r.PathPrefix("/").HandlerFunc(static.Serve)

	go func() {
		log.Printf("Server will listen to: %s", config.Config.Addr)
		log.Fatal(http.ListenAndServe(config.Config.Addr, r))
	}()
}
