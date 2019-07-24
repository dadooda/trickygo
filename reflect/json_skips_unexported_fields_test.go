
package trickygo

func ExampleJsonSkipsUnexportedFields() {
	JsonSkipsUnexportedFields()

	// Output:
	// openPerson before:{Alice 10}
	// openPerson after:{Bob 35}
	// secretPerson before:{Alice 10}
	// secretPerson after:{Bob 10}
}