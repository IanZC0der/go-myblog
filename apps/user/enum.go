package user

type Role int

const (
	ROLE_MEMBER Role = iota
	ROLE_ADMIN
)

type QueryBy int

const (
	QUERY_BY_ID QueryBy = iota
	QUERY_BY_USERNAME
)
