
package trickygo

import (
	"encoding/json"
	"fmt"
	"reflect"
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

func ExampleReflectionCheatsheet() {
	type person struct {
		Name string
		age int
	}

	//
	// Basic type/value retrieval.
	//

	prs := person{"Alice", 10}
	elem := reflect.ValueOf(&prs).Elem()
	ty := elem.Type()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		fieldT := ty.Field(i)
		fmt.Printf("field %d: name:%#v CanSet:%#v\n", i, fieldT.Name, field.CanSet())
	}

	// Output:
	// field 0: name:"Name" CanSet:true
	// field 1: name:"age" CanSet:false
}