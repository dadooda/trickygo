
package trickygo

// JSON decoder can't overwrite unexported fields and thus silently skips them, retaining previously set
// values, if any.
//
// In the example below, `openPerson` has all fields exported and JSON decoder updates them. Another struct,
// `secretPerson`, has `age` unexported and this field is silently skipped by JSON decoder, although JSON data
// contains a valid `age` field.
func JsonDecoderSkipsUnexportedFields() {}

// Go reflection types and methods are often confusing.
//
// In the example below, some common cases are covered in form of a quick cheatsheet.
func ReflectionCheatsheet() {}