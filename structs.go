package zenoss

type CAResponse struct {
	UUID   string `json:"uuid"`
	Action string `json:"action"`
	Result struct {
		Msg     string `json:"msg"`
		Success bool   `json:"success"`
	} `json:"result"`
	Tid    int    `json:"tid"`
	Type   string `json:"type"`
	Method string `json:"method"`
}
