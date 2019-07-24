
package trickygo

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ExampleJsonDecoderSkipsUnexportedFields() {
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

	// Output:
	// openPerson before:{Alice 10}
	// openPerson after:{Bob 35}
	// secretPerson before:{Alice 10}
	// secretPerson after:{Bob 10}
}