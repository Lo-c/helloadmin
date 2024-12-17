package acme

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"helloadmin/internal/ecode"
	"helloadmin/internal/repository"
)

type Repository interface {
	Find(ctx context.Context, req *FindRequest) (int64, *[]Model, error)
	GetAcmePath(ctx context.Context, id int64) (*Model, error)
}

func NewRepository(r *repository.Repository) Repository {
	return &acmeRepository{
		Repository: r,
	}
}

type acmeRepository struct {
	*repository.Repository
}

func (r *acmeRepository) Find(ctx context.Context, req *FindRequest) (int64, *[]Model, error) {
	var count int64
	var acmes []Model
	query := r.DB(ctx)
	if req.Id != "" {
		query = query.Where("id = ?", req.Id)
	}
	query.Model(Model{}).Count(&count)
	if err := query.Order("sort DESC").Find(&acmes).Error; err != nil {
		return count, nil, err
	}
	return count, &acmes, nil
}

func (r *acmeRepository) GetAcmePath(ctx context.Context, id int64) (*Model, error) {
	var acme Model
	if err := r.DB(ctx).Where("id = ?", id).First(&acme).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ecode.ErrAcmePathNotFound
		}
		return nil, err
	}
	return &acme, nil
}
