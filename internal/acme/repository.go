package acme

import (
	//"context"
	//"errors"
	//"gorm.io/gorm"
	//"helloadmin/internal/ecode"
	"helloadmin/internal/repository"
)

type Repository interface {
}

func NewRepository(r *repository.Repository) Repository {
	return &acmeRepository{
		Repository: r,
	}
}

type acmeRepository struct {
	*repository.Repository
}
