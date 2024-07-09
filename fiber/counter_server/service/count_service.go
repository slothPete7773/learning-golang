package service

import "counter_server/domain"

type countService struct {
	countRepo domain.Repository
}

func NewCountService(countRepo domain.Repository) *countService {
	return &countService{countRepo: countRepo}
}

func (c *countService) Increase() (int32, error) {
	return c.countRepo.Increase()
}

func (c *countService) Decrease() (int32, error) {
	return c.countRepo.Decrease()
}

func (c *countService) Get() (int32, error) {
	return c.countRepo.Get()
}

func (c *countService) Reset() error {
	return c.countRepo.Reset()
}
