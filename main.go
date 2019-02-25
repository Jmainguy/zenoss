package zenoss

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
    "time"
)

func CreateAlarm(url, user, password, summary, device, component, severity, evclasskey, evclass string) (UUID string, success bool) {
	payload := fmt.Sprintf(`{
                "action": "EventsRouter",
                "method": "add_event",
                "data": [{
                    "summary": "%s",
                    "device": "%s",
                    "component": "%s",
                    "severity": "%s",
                    "evclasskey": "%s",
                    "evclass": "%s"
                }],
                "type": "rpc",
                "tid": 1
}`, summary, device, component, severity, evclasskey, evclass)

	p := strings.NewReader(payload)
	req, err := http.NewRequest("POST", url, p)
	if err != nil {
		fmt.Println(err)
		UUID = ""
		success = false
		return
	}
	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")
    // By Default golangs http client never times out, use a custom one
    // https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
    client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
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
