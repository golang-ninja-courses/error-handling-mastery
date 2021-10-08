package jwt

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Token struct {
	Email     string   `json:"email"`
	Subject   string   `json:"subject"`
	Scopes    []string `json:"scopes"`
	ExpiredAt int64    `json:"expired_at"`
}
