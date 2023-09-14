package usecases

type Command interface {
	Validate() error
	Execute() error
}

type Query[T any] interface {
	Validate() error
	Execute() (T, error)
}
