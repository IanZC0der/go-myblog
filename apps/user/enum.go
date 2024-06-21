package user

type Role int

const (
	ROLE_AUTHOR Role = iota
	ROLE_ADMIN
	ROLE_AUDITOR
)

type QueryBy int

const (
	QUERY_BY_ID QueryBy = iota
	QUERY_BY_USERNAME
)
