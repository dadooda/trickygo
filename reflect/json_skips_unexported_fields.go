
package trickygo

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JSON decoder can't overwrite unexported fields and thus silently skips them, retaining previously set
// values, if any.
//
// In the example below, `openPerson` has all fields exported and JSON decoder updates them. Another struct
// `secretPerson`, has `age` unexported and this field is silently skipped by JSON encoder.
func JsonSkipsUnexportedFields() {
	type openPerson struct {
		Name string
		Age int
	}

	oprs := openPerson{"Alice", 10}
	fmt.Printf("openPerson before:%v\n", oprs)
	err := json.NewDecoder(strings.NewReader(`{"name": "Bob", "age": 35}`)).Decode(&oprs)
	if err != nil { panic(err) }
	fmt.Printf("openPerson after:%v\n", oprs)

	type secretPerson struct {
		Name string
		age int				// Unexported.
	}

	sprs := secretPerson{"Alice", 10}
	fmt.Printf("secretPerson before:%v\n", sprs)
	err = json.NewDecoder(strings.NewReader(`{"name": "Bob", "age": 35}`)).Decode(&sprs)
	if err != nil { panic(err) }
	fmt.Printf("secretPerson after:%v\n", sprs)
}