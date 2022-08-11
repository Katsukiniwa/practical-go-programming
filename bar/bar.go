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

type user struct {
	UserID    string   `json:"use_id"`
	UserName  string   `json:"user_name"`
	Languages []string `json:"languages"`
}

func PrintUser() {
	u := user{
		UserID:    "001",
		UserName:  "gopher",
		Languages: []string{}, // 空スライスをセットするとnilではなく空配列がセットされる
	}

	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}

type FormInput struct {
	Name        string `json:"name"`
	CompanyName string `json:"company_name,omitempty"`
}

func PrintFormInput() {
	in := FormInput{Name: "山田太郎"}

	b, _ := json.Marshal(in)
	fmt.Println(string(b))
}

type Bottle struct {
	Name  string `json:"name"`
	Price int    `json:"price,omitempty"`
	KCal  *int   `json:"kcal,omitempty"`
}

func PrintBottle() {
	b := Bottle{
		Name:  "ミネラルウォーター",
		Price: 0,
		KCal:  Int(0),
	}

	out, _ := json.Marshal(b)
	fmt.Println(string(out))
}

func Int(v int) *int {
	return &v
}
