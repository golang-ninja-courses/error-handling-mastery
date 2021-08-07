package errctx

type Fields map[string]interface{}

func AppendTo(err error, fields Fields) error {
	// Реализуй меня.
	return nil
}

func From(err error) Fields {
	// Реализуй меня.
	return nil
}
