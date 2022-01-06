package factory

func NewEvent(payload interface{}) (event EventInterface) {
	if payload == "gitlab" {
		return &GitlabPushHookEvent{}
	}
	return 
}