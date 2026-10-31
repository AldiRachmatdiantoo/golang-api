package services

import "github.com/AldiRachmatdiantoo/golang-api.git/internal/models"

type userRepository interface {
	FindAll()
}
type Service struct {
	repo userRepository
	user []models.User
}

func NewService(rp userRepository) *Service {
	return &Service{repo: rp}
}
func (s *Service) GetUsersData() {

}
