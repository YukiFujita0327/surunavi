package _const

type LoginResultType int
const (
	Success LoginResultType = iota
	NoUser
	MissPassword
)
