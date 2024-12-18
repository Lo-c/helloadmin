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
	GetById(ctx context.Context, id int64) (*Model, error)
	Create(ctx context.Context, acme *Model) error
	Update(ctx context.Context, id int64, acme *Model) error
	Delete(ctx context.Context, id int64) error
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
	if err := query.Find(&acmes).Error; err != nil {
		return count, nil, err
	}
	return count, &acmes, nil
}

func (r *acmeRepository) GetById(ctx context.Context, id int64) (*Model, error) {
	var acme Model
	if err := r.DB(ctx).Where("id = ?", id).First(&acme).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ecode.ErrAcmeIdNotFound
		}
		return nil, err
	}
	return &acme, nil
}

func (r *acmeRepository) Create(ctx context.Context, acme *Model) error {
	if err := r.DB(ctx).Create(acme).Error; err != nil {
		return err
	}
	return nil
}

func (r *acmeRepository) Update(ctx context.Context, id int64, acme *Model) error {
	if err := r.DB(ctx).Model(acme).
		Where("id = ?", id).
		Updates(acme).Error; err != nil {
		return err
	}
	return nil
}

func (r *acmeRepository) Delete(ctx context.Context, id int64) error {
	if err := r.DB(ctx).Delete(&Model{}, id).Error; err != nil {
		return err
	}
	return nil
}
