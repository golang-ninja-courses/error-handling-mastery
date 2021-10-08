package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// wrapWithPkgErrors - правило, которое предлагает врапить ошибки через github.com/pkg/errors.
func wrapWithPkgErrors(m dsl.Matcher) {
	m.Import("github.com/pkg/errors")

	m.Match(`if err := $x; err != nil { return err }`).
		Report(`err is not wrapped`).
		Suggest(`if err := $x; err != nil { return errors.Wrap($x, "FIXME: wrap the error") }`)

	// Прочие варианты.
	// ...
}
