package query

type Type string

type Command interface {
	Type() Type
}
