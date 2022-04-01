package types

type ReqCmd struct {
	Id     string   `json:"id"`
	Repo   string   `json:"repo"`
	Params []string `json:"params"`
}

type ResCmd struct {
	Out string `json:"out"`
	Err string `json:"err"`
}
