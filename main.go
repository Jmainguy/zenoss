package zenoss

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func CreateAlarm(url, user, password, summary, device, severity string) (UUID string, success bool) {
	payload := fmt.Sprintf(`{
                "action": "EventsRouter",
                "method": "add_event",
                "data": [{
                    "summary": "%s",
                    "device": "%s",
                    "component": "",
                    "severity": "%s",
                    "evclasskey": "",
                    "evclass": "/App"
                }],
                "type": "rpc",
                "tid": 1
}`, summary, device, severity)

	p := strings.NewReader(payload)
	req, err := http.NewRequest("POST", url, p)
	if err != nil {
		fmt.Println(err)
		UUID = ""
		success = false
		return
	}
	if err != nil {
		fmt.Println(err)
		UUID = ""
		success = false
		return
	}
	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		UUID = ""
		success = false
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		UUID = ""
		success = false
		return
	}
	response := CAResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		UUID = ""
		success = false
		return
	}
	UUID = response.UUID
	success = response.Result.Success
	return
}
