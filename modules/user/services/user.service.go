package services

import (
	"fmt"
	"main/entity"
	"main/pkg/validators"
)

type UserService struct {
	Repo Repository
}

func New(repo Repository) UserService {
	return UserService{repo}
}

type Repository interface {
	IsUniquePhoneNumber(phone string) (bool, error)
	RegisterUser(user entity.UserEntity) (entity.UserEntity, error)
}

type RegisterRequest struct {
	Name  string
	Phone string
}

type RegisterResponse struct {
	User entity.UserEntity
}

func (us UserService) Register(req RegisterRequest) (RegisterResponse, error) {
	//TODO - Otp verification
	//* validation here
	if !validators.IsPhoneNumber(req.Phone) {
		return RegisterResponse{}, fmt.Errorf("Invalid phone-number")
	}

	if isUnique, err := us.Repo.IsUniquePhoneNumber(req.Phone); err != nil {
		return RegisterResponse{}, fmt.Errorf("Unexpected error : %w", err)
	} else if !isUnique {
		return RegisterResponse{}, fmt.Errorf("Phone number registered before")
	}

	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("Name must be atLeast 4 character")
	}

	//*Create-user
	newUser := entity.UserEntity{
		Id:    0,
		Phone: req.Phone,
		Name:  req.Phone,
	}

	createdUser, err := us.Repo.RegisterUser(newUser)

	if err != nil {
		return RegisterResponse{}, fmt.Errorf("Unexpected error : %w", err)
	}

	return RegisterResponse{
		User: createdUser,
	}, nil

}
