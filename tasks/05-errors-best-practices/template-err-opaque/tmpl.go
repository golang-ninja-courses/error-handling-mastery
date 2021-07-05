package tmpl

// IsFunctionNotDefinedError говорит, является ли err ошибкой неопределённой в шаблоне функции.
func IsFunctionNotDefinedError(err error) bool {
	return false
}

// IsExecUnexportedFieldError говорит, является ли err template.ExecError,
// а именно ошибкой использования неэспортируемого поля структуры.
func IsExecUnexportedFieldError(err error) bool {
	return false
}
