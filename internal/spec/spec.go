// Package spec describes the Yesly API operations the CLI exposes. The table
// itself is generated from the OpenAPI specification into zz_generated.go.
package spec

// Kind is how a flag's value is read and sent.
type Kind int

const (
	// String is sent as a JSON string.
	String Kind = iota
	// Integer is sent as a JSON number without decimals.
	Integer
	// Number is sent as a JSON number.
	Number
	// Bool is sent as true or false.
	Bool
	// StringList is a comma-separated or repeated flag, sent as an array of strings.
	StringList
	// Map is key=value pairs, sent as an object of strings.
	Map
	// JSON is any JSON value, given inline or as @file.
	JSON
)

// Field is one flag: a query parameter or a body property. Body properties
// inside objects use dotted names, such as contact.name.
type Field struct {
	Name        string
	Kind        Kind
	Required    bool
	Enum        []string
	Default     string
	Description string
}

// Param is a path parameter, given as a positional argument.
type Param struct {
	Name        string
	Description string
}

// Operation is one API call.
type Operation struct {
	Resource    string
	Method      string
	HTTP        string
	Path        string
	Summary     string
	Description string
	PathParams  []Param
	Query       []Field
	Body        []Field
	// HasBody is true when the operation takes a JSON body, even an empty one.
	HasBody bool
	// Idempotent operations get an Idempotency-Key.
	Idempotent bool
	// List operations return a page and take starting_after.
	List bool
	// Deprecated operations keep working but should not be used in new code.
	Deprecated bool
}
