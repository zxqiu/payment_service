package actions

type ActionMetadata interface {
	GetActionID() string
}

type Action interface {
	Execute(metadata ActionMetadata) error
}
