package factory

type EventInterface interface {
	GetBranch() string
}

type Event struct {
	PayloadSource string
	PayloadMap    map[string]interface{}
}
