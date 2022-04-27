package errctx

type Fields map[string]any

func AppendTo(err error, fields Fields) error {
	// Реализуй меня.
	return nil
}

func From(err error) Fields {
	// Реализуй меня.
	return nil
}
