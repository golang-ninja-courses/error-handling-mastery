package errors

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapper(t *testing.T) {
	ifaceObj := (*Wrapper)(nil)

	assertNotImplements(t, ifaceObj, simpleError(0))
	assertNotImplements(t, ifaceObj, withUnwrapOnly(0))
	assertNotImplements(t, ifaceObj, withCauseOnly(0))
	assertNotImplements(t, ifaceObj, withUnwrapError(0))
	assertNotImplements(t, ifaceObj, withCauseError(0))
	assertNotImplements(t, ifaceObj, withBothNotError(0))
	assertImplements(t, ifaceObj, withBothError(0))
}

func assertImplements(t *testing.T, ifaceObj interface{}, obj interface{}) {
	t.Helper()
	assert.Implements(t, ifaceObj, obj)
}

func assertNotImplements(t *testing.T, ifaceObj interface{}, obj interface{}) {
	t.Helper()
	ifaceType := reflect.TypeOf(ifaceObj).Elem()
	assert.False(t, reflect.TypeOf(obj).Implements(ifaceType))
}

type simpleError int                //
func (e simpleError) Error() string { return "" }

type withUnwrapOnly int                //
func (e withUnwrapOnly) Unwrap() error { return nil }

type withCauseOnly int               //
func (e withCauseOnly) Cause() error { return nil }

type withUnwrapError int                //
func (e withUnwrapError) Error() string { return "" }
func (e withUnwrapError) Unwrap() error { return nil }

type withCauseError int                //
func (e withCauseError) Error() string { return "" }
func (e withCauseError) Cause() error  { return nil }

type withBothNotError int                //
func (e withBothNotError) Unwrap() error { return nil }
func (e withBothNotError) Cause() error  { return nil }

type withBothError int                //
func (e withBothError) Error() string { return "" }
func (e withBothError) Unwrap() error { return nil }
func (e withBothError) Cause() error  { return nil }
