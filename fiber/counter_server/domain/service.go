package domain

type Service interface {
	Increase() (int32, error)
	Decrease() (int32, error)
	Get() (int32, error)
	Reset() error
}

type Repository interface {
	Increase() (int32, error)
	Decrease() (int32, error)
	Get() (int32, error)
	Reset() error
}
