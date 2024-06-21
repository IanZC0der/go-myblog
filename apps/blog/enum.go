package blog

type Status int

const (
	DRAFT Status = iota
	PUBLISHED
)

// type Role int

// const (
// 	AUTHOR Role = iota
// 	AUDITOR
// 	ADMIN
// )

type UpdateMode int

const (
	UPDATE_MODE_PUT UpdateMode = iota
	UPDATE_MODE_PATCH
)
