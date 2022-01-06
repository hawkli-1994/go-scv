package factory

type GitlabPushHookEvent struct {
	event Event
}

func (e *GitlabPushHookEvent) GetBranch() string {
	return ""
}
