package tmpl

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"unsafe"
)

func TestParseAndExecuteTemplate_Success(t *testing.T) {
	tmpl := `
<html>
	<body>
		{{ with .Name -}}
		<h1>Hello {{ . }}!</h1>
		{{- end }}
	</body>
</html>`
	b := bytes.NewBuffer(nil)

	err := ParseAndExecuteTemplate(b, "greeting", tmpl, struct{ Name string }{Name: "Anthony"})
	require.NoError(t, err)
	assert.Equal(t, `
<html>
	<body>
		<h1>Hello Anthony!</h1>
	</body>
</html>`, b.String())
}

func TestParseAndExecuteTemplate_InvalidTemplate(t *testing.T) {
	tmpl := `
<html>
	<body>
		{{{ with .Name -}}
		<h1>Hello {{ . }}!</h1>
		{{- end }}
	</body>
</html>`
	b := bytes.NewBuffer(nil)

	err := ParseAndExecuteTemplate(b, "greeting", tmpl, struct{ Name string }{Name: "Anthony"})
	require.Error(t, err)
	assert.ErrorIs(t, err, errParseTemplate)
}

func TestParseAndExecuteTemplate_ExecutingError(t *testing.T) {
	tmpl := `
<html>
	<body>
		{{ with .Name -}}
		<h1>Hello {{ . }}!</h1>
		{{- end }}
	</body>
</html>`
	b := bytes.NewBuffer(nil)

	err := ParseAndExecuteTemplate(b, "greeting", tmpl, struct{ Name unsafe.Pointer }{})
	require.Error(t, err)
	assert.ErrorIs(t, err, errExecuteTemplate)
}
