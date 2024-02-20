package types

type FlagDetail struct {
	Name string
	Type string
}

type FlagValue struct {
	Value           interface{}
	DefaultNilValue interface{}
}
