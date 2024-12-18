package acme

import "context"

type Service interface {
	GetAcme(ctx context.Context, id int64) (*AcmeData, error)
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

func (s *service) GetAcme(ctx context.Context, id int64) (*AcmeData, error) {
	acme, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &AcmeData{
		Id:       acme.ID,
		AcmePath: acme.ConfigPath,
		Email:    acme.Email,
	}, nil
}

func (s *service) CreateAcme(ctx context.Context, req *CreateRequest) error {
	acme := Model{
		ID:         req.Id,
		ConfigPath: req.AcmePath,
		Email:      req.Email,
	}
	return s.repo.Create(ctx, &acme)
}
