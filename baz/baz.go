package baz

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type JSTime time.Time

type Record struct {
	ProcessID string `json:"process_id"`
	DeletedAt JSTime `json:"deleted_at"`
}

func (t JSTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	if tt.IsZero() {
		return []byte("null"), nil
	}
	v := strconv.Itoa(int(tt.UnixMilli()))

	return []byte(v), nil
}

func (t *JSTime) UnmarshalJSON(data []byte) error {
	var jsonNumber json.Number
	err := json.Unmarshal(data, &jsonNumber)
	if err != nil {
		return err
	}
	unix, err := jsonNumber.Int64()
	if err != nil {
		return err
	}

	*t = JSTime(time.Unix(0, unix))
	return nil
}

func Run() {
	r := Record{
		ProcessID: "0001",
		DeletedAt: JSTime(time.Now()),
	}

	b, _ := json.Marshal(r)
	fmt.Println(string(b))

	v := `{"process_id":"001","deleted_at":1639485878560}`
	var a *Record
	if err := json.Unmarshal([]byte(v), &a); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", time.Time(r.DeletedAt).Format(time.RFC3339Nano))
}
