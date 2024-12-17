package acme

import (
	"context"
	"helloadmin/internal/repository"
)

type Repository interface {
	Find(ctx context.Context, req *FindRequest) (int64, *[]Model, error)
	UpdateAcmePath(ctx context.Context, id int64, acme *Model) error
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

func (r *acmeRepository) UpdateAcmePath(ctx context.Context, id int64, acme *Model) error {
	if err := r.DB(ctx).Model(acme).
		Select("id", "acme_path", "email").
		Where("id = ?", id).
		Updates(acme).Error; err != nil {
		return err
	}
	return nil
}
