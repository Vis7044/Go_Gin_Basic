package services

import (
	"github.com/Vis7044/GinCrud2/models"
	"github.com/Vis7044/GinCrud2/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestService struct {
	repo *repository.TestRepository
}

func NewTestService(r *repository.TestRepository) *TestService {
	return &TestService{
		repo: r,
	}
}

func (s *TestService) Create(test models.Test) (*models.Test, error) {
	test.Id = primitive.NewObjectID()
	_, err := s.repo.Create(test)
	if err != nil {
		return nil, err
	}
	return &test, nil
}

func (s *TestService) GetAll(limit, skip int) (*[]models.Test, error) {
	if(limit<=0) {
		limit=15
	}
	if(skip <0) {
		skip = 0
	}
	result, err := s.repo.GetAll(limit, skip)
	if err != nil {
		return nil , err
	}
	return result, nil
}

func (s *TestService) GetOne(id primitive.ObjectID) (*models.Test, error) {
	result, err := s.repo.GetOne(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *TestService) UpdateOne( id primitive.ObjectID, test models.Test) (*int64, error) {
	result , err := s.repo.UpdateOne(id, test)
	if err != nil {
		return nil, err

	}
	return result, nil
}

func (s *TestService) DeleteOne(id primitive.ObjectID) (*int64, error) {
	result, err := s.repo.DeleteTest(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}