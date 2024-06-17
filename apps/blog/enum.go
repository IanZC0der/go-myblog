package blog

type Status int

const (
	DRAFT Status = iota
	PUBLISHED
)

type UpdateMode int

const (
	UPDATE_MODE_PUT UpdateMode = iota
	UPDATE_MODE_PATCH
)
