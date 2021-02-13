package _const

type LoginResultType int
const (
	Succese LoginResultType = iota
	NoUser
	MissPassword
)
