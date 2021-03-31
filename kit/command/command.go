package command

type Type string

type Command interface {
	Type() Type
}
