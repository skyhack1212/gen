// Generated by: gen
// TypeWriter: gen
// Directive: +gen on typewriter.Tag

package typewriter

// Tags is a slice of type Tag, for use with gen methods below. Use this type where you would use []Tag. (This is required because slices cannot be method receivers.)
type Tags []Tag

// Where returns a new Tags slice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
func (rcv Tags) Where(fn func(Tag) bool) (result Tags) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
