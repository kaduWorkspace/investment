package repository

type Repository[T any] interface {
    Save(fields interface{})
    Update(fields interface{})
    Get(filters interface{}) T
    Delete(fields interface{})
}
