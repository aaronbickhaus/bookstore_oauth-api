package access_token

import "github.com/aaronbickhaus/bookstore_oauth-api/src/utils/errors"

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpriationTime(AccessToken)  *errors.RestErr
}

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpriationTime(AccessToken)  *errors.RestErr
}

type service struct {
	repository Repository
}

func (s *service) Create(at AccessToken) *errors.RestErr {
	err := s.repository.Create(at)
	if err != nil {
		return  err
	}
	return  nil
}

func (s *service) UpdateExpriationTime(at AccessToken) *errors.RestErr {
	err := s.repository.UpdateExpriationTime(at)
	if err != nil {
		return  err
	}
	return  nil
}

func NewService(repo Repository) Service {
	return &service{
		repository:  repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
  accessToken, err := s.repository.GetById(accessTokenId)
  if err != nil {
  	return nil, err
  }
  return accessToken, nil
}