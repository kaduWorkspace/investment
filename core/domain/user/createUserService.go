package user

import "kaduhod/fin_v3/core/domain/dto"

type CreateUserServiceI[T dto.Dto] interface {
    Create(usr T) (error)
}
