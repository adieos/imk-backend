package service

import (
	"context"
	"os"
	"sync"

	"github.com/adieos/imk-backend/dto"
	"github.com/adieos/imk-backend/entity"
	"github.com/adieos/imk-backend/helpers"
	"github.com/adieos/imk-backend/repository"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error)
		GetUserById(ctx context.Context, userId string) (dto.UserResponse, error)

		Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error)
	}

	userService struct {
		userRepo   repository.UserRepository
		jwtService JWTService
	}
)

func NewUserService(userRepo repository.UserRepository,
	jwtService JWTService) UserService {
	return &userService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

var (
	mu sync.Mutex
)

var (
	LOCAL_URL          = os.Getenv("APP_URL")
	VERIFY_EMAIL_ROUTE = "register/verify_email"
	RESET_EMAIL_ROUTE  = "reset"
)

func (s *userService) RegisterUser(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	_, flag, _ := s.userRepo.CheckEmail(ctx, nil, req.Email)
	if flag {
		return dto.UserResponse{}, dto.ErrEmailAlreadyExists
	}

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	userReg, err := s.userRepo.RegisterUser(ctx, nil, user)
	if err != nil {
		return dto.UserResponse{}, dto.ErrCreateUser
	}

	return dto.UserResponse{
		Name:  userReg.Name,
		Email: userReg.Email,
	}, nil
}

func (s *userService) GetUserById(ctx context.Context, userId string) (dto.UserResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserById
	}

	return dto.UserResponse{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *userService) Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	check, flag, err := s.userRepo.CheckEmail(ctx, nil, req.Email)
	if err != nil || !flag {
		return dto.UserLoginResponse{}, dto.ErrInvalidCredentials
	}

	checkPassword, err := helpers.CheckPassword(check.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		return dto.UserLoginResponse{}, dto.ErrInvalidCredentials
	}

	user, err := s.userRepo.GetUserByEmail(ctx, nil, check.Email)
	if err != nil {
		return dto.UserLoginResponse{}, dto.ErrGetUserByEmail
	}

	token := s.jwtService.GenerateTokenUser(user.ID.String(), "USER")

	return dto.UserLoginResponse{
		Token: token,
	}, nil
}
