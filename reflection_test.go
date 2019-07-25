
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
		IsMale bool
	}

	fmt.Println("# Basic type/value retrieval")

	prs := person{"Alice", 10, false}
	elem := reflect.ValueOf(&prs).Elem()
	ty := elem.Type()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		fieldT := ty.Field(i)
		fmt.Printf("field %d: name:%#v Type().Name():%#v CanSet:%#v String():%#v\n", i, fieldT.Name, field.Type().Name(), field.CanSet(), field.String())
	}

	fmt.Println("# Type-specific value retrieval")

	// NOTE: Lookup is case-sensitive!
	fName := elem.FieldByName("Name")
	fAge := elem.FieldByName("age")
	fIsMale := elem.FieldByName("IsMale")
	fmt.Printf("fName.String():%#v fAge.Int():%#v fIsMale.Bool():%#v", fName.String(), fAge.Int(), fIsMale.Bool())

	// Output:
	// # Basic type/value retrieval
	// field 0: name:"Name" Type().Name():"string" CanSet:true String():"Alice"
	// field 1: name:"age" Type().Name():"int" CanSet:false String():"<int Value>"
	// field 2: name:"IsMale" Type().Name():"bool" CanSet:true String():"<bool Value>"
	// # Type-specific value retrieval
	// fName.String():"Alice" fAge.Int():10 fIsMale.Bool():false
}