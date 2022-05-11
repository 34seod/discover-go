package serialize

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/34seod/discover-go/enum"
)

func Example_marshalJSON() {
	t := enum.Task{
		Title:    "Laundry",
		Status:   enum.DONE,
		Daedline: enum.NewDaedline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
	// Output:
	// {"title":"Laundry","status":"DONE","daedline":"2015-08-16T15:43:00Z"}
}

func Example_unmarshalJSON() {
	b := []byte(`{"title":"Buy Milk","status":"DONE","daedline":"2015-08-16T15:43:00Z"}`)
	t := enum.Task{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t.Title)
	fmt.Println(t.Status)
	fmt.Println(t.Daedline.UTC())
	// Output:
	// Buy Milk
	// 2
	// 2015-08-16 15:43:00 +0000 UTC
}

func Example_mapMarshalJSON() {
	type any interface{}
	b, _ := json.Marshal(map[string]any{
		"Name": "John",
		"Age":  16,
	})
	fmt.Println(string(b))
	// Output:
	// {"Age":16,"Name":"John"}
}

type Fields struct {
	VisibleField   string
	InvisibleField string
}
type InFields struct {
	Fields
	InvisibleField string `json:",omitempty"`
	Additional     string
}

func ExampleOmitFields() {
	f := &Fields{"a", "b"}
	b, _ := json.Marshal(InFields{Fields: *f, Additional: "c"})
	fmt.Println(string(b))
	x := InFields{
		Fields:         *f,
		InvisibleField: "aa",
		Additional:     "c",
	}
	fmt.Println(x.InvisibleField)
	x.VisibleField = "eee"
	x.InvisibleField = "ffff"
	x.Fields.InvisibleField = "dddd"
	fmt.Println(f.InvisibleField)
	fmt.Println(x.InvisibleField)
	fmt.Println(x.VisibleField)
	fmt.Println(x.Fields.InvisibleField)
	// Output:
	// {"VisibleField":"a","Additional":"c"}
	// dddd
}
