
package trickygo

// JSON decoder can't overwrite unexported fields and thus silently skips them, retaining previously set
// values, if any.
//
// In the example below, `openPerson` has all fields exported and JSON decoder updates them. Another struct,
// `secretPerson`, has `age` unexported and this field is silently skipped by JSON decoder.
func JsonDecoderSkipsUnexportedFields() {}
