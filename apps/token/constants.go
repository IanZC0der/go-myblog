package token

const (
	TOKEN_COOKIE_NAME        = "access_token"
	TOKEN_GIN_KEY_IN_CONTEXT = "access_token"
)

type QueryTokenType int

const (
	QUERY_BY_ID = iota
	QUERY_BY_ACCESS_TOKEN
)
