package bar

import (
	"encoding/json"
	"fmt"
	"log"
)

func Sample() {
	fmt.Print("ip")
}

type ipWithPrivateURL struct {
	Origin string `json:"origin"`
	url    string `json:"url"`
}

func DecodeIp() {
	s := `{"origin":"255.255.255.255","url": "https://httpbin.org/get"}`
	var resp ipWithPrivateURL
	if err := json.Unmarshal([]byte(s), &resp); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}
