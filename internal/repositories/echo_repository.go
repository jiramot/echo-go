package repository

import (
    domain "jiramot/echo-go/internal/core/domain"
)

type repository struct {

}

func NewEchoRepository() *repository {
    return &repository{}
}

func (repo *repository) Echo(message string) (domain.Echo, error){
    return domain.Echo{}, nil
}