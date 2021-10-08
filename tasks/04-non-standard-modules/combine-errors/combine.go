package errors

// Combine "прицепляет" ошибки other к err так, что они начинают фигурировать при выводе
// её на экран через спецификатор `%+v`. Если err является nil, то первостепенной ошибкой
// становится первая из ошибок other.
func Combine(err error, other ...error) error {
	// Реализуй меня.
	return nil
}
