package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/katsukiniwa/practical-go-programming/bar"
)

type ip struct {
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

func ipFunc() {
	f, err := os.Open("ip.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var resp ip
	if err := json.NewDecoder(f).Decode(&resp); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func main() {
	fmt.Print("Hello World")
	ipFunc()
	bar.Sample()
	bar.DecodeIp()
}
