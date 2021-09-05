package actions

type ActionMetadata interface{}

type Action interface {
	Execute(metadata ActionMetadata) error
}
