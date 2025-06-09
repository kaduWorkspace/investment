package repository

type Connection interface {
}

type Repository[T any] interface {
    Save(fields T) (int, error)
    Update(fields T) error
    Get(filters T) (T, error)
    Delete(fields T) error
}
