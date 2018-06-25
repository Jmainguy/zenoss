package zenoss

import (
	"fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

func check(e error) {
    if e != nil {
        fmt.Println(e)
    }
}

func createAlarm(url, user, password, summary, device string) (response string) {
    payload := fmt.Sprintf(`{
                "action": "EventsRouter",
                "method": "add_event",
                "data": [{
                    "summary": "%s",
                    "device": %s,
                    "component": "golang =)",
                    "severity": "Critical",
                    "evclasskey": "",
                    "evclass": "/App"
                }],
                "type": "rpc",
                "tid": 1
}`, summary, device)

    p := strings.NewReader(payload)
	req, err := http.NewRequest("POST", url, p)
	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
    bodyBytes, err := ioutil.ReadAll(resp.Body)
    check(err)
    response = string(bodyBytes)
    return response
}
