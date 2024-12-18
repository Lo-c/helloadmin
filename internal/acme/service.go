package acme

import "context"

type Service interface {
	CreateAcme(ctx context.Context, req *CreateRequest) error
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

type service struct {
	repo Repository
}

func (s *service) CreateAcme(ctx context.Context, req *CreateRequest) error {
	acme := Model{
		ID:         req.Id,
		ConfigPath: req.AcmePath,
		Email:      req.Email,
	}
	return s.repo.Create(ctx, &acme)
}
