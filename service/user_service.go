package service

import (
	"bytes"
	"context"
	"html/template"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/adieos/imk-backend/constants"
	"github.com/adieos/imk-backend/dto"
	"github.com/adieos/imk-backend/entity"
	"github.com/adieos/imk-backend/helpers"
	"github.com/adieos/imk-backend/repository"
	"github.com/adieos/imk-backend/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error)
		GetUserById(ctx context.Context, userId string) (dto.UserResponse, error)

		SendVerificationEmail(ctx context.Context, req dto.SendVerificationEmailRequest) error
		VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) (dto.VerifyEmailResponse, error)
		UpdateUser(ctx context.Context, req dto.UserUpdateRequest, userId string) (dto.UserUpdateResponse, error)
		DeleteUser(ctx context.Context, userId string) error
		Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error)
		ResetPassword(ctx context.Context, email, newPassword string) error
		ForgetPassword(ctx context.Context, req dto.ForgetPasswordRequest) error
		MakeForgetPasswordEmail(receiverEmail string) (map[string]string, error)
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

	provinceID, err := uuid.Parse(req.Province)
	if err != nil {
		return dto.UserResponse{}, dto.ErrParseUUID
	}

	user := entity.User{
		Name:       req.Name,
		Email:      req.Email,
		Password:   req.Password,
		Instansi:   req.Institution,
		NoTelp:     req.TelpNumber,
		InfoFrom:   req.InfoFrom,
		Jenjang:    req.EducationalLevel,
		Role:       constants.ENUM_ROLE_USER,
		IsVerified: true,
		ProvinceID: provinceID,
	}

	userReg, err := s.userRepo.RegisterUser(ctx, nil, user)
	if err != nil {
		return dto.UserResponse{}, dto.ErrCreateUser
	}

	return dto.UserResponse{
		Name:     userReg.Name,
		Email:    userReg.Email,
		Instansi: userReg.Instansi,
		NoTelp:   userReg.NoTelp,
		InfoFrom: userReg.InfoFrom,
		Jenjang:  userReg.Jenjang,
	}, nil
}

func makeVerificationEmail(receiverEmail string) (map[string]string, error) {
	expired := time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")
	plainText := receiverEmail + "_" + expired
	token, err := utils.AESEncrypt(plainText)
	if err != nil {
		return nil, err
	}

	appUrl := os.Getenv("APP_URL")

	verifyLink := appUrl + "/" + VERIFY_EMAIL_ROUTE + "?token=" + token

	readHtml, err := os.ReadFile("utils/email-template/verification_email.html")
	if err != nil {
		return nil, err
	}

	data := struct {
		Email  string
		Verify string
	}{
		Email:  receiverEmail,
		Verify: verifyLink,
	}

	tmpl, err := template.New("custom").Parse(string(readHtml))
	if err != nil {
		return nil, err
	}

	var strMail bytes.Buffer
	if err := tmpl.Execute(&strMail, data); err != nil {
		return nil, err
	}

	draftEmail := map[string]string{
		"subject": "ILITS 2025 - Verification Email",
		"body":    strMail.String(),
	}

	return draftEmail, nil
}

func (s *userService) SendVerificationEmail(ctx context.Context, req dto.SendVerificationEmailRequest) error {
	user, err := s.userRepo.GetUserByEmail(ctx, nil, req.Email)
	if err != nil {
		return dto.ErrEmailNotFound
	}

	draftEmail, err := makeVerificationEmail(user.Email)
	if err != nil {
		return err
	}

	err = utils.SendMail(user.Email, draftEmail["subject"], draftEmail["body"])
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) (dto.VerifyEmailResponse, error) {
	decryptedToken, err := utils.AESDecrypt(req.Token)
	if err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrTokenInvalid
	}

	tokenParts := strings.Split(decryptedToken, "_")
	if len(tokenParts) < 2 {
		return dto.VerifyEmailResponse{}, dto.ErrTokenInvalid
	}

	email := tokenParts[0]
	expirationDate := tokenParts[1]
	expirationTime, err := time.Parse("2006-01-02 15:04:05", expirationDate)
	if err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrTokenInvalid
	}
	// email, expired, err := s.jwtService.GetUserEmailByToken(req.Token)
	// if err != nil {
	// 	return dto.VerifyEmailResponse{}, dto.ErrTokenInvalid
	// }

	now := time.Now()

	if expirationTime.Before(now) {
		return dto.VerifyEmailResponse{
			Email:      email,
			IsVerified: false,
		}, dto.ErrTokenExpired
	}

	user, err := s.userRepo.GetUserByEmail(ctx, nil, email)
	if err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrUserNotFound
	}

	if user.IsVerified {
		return dto.VerifyEmailResponse{}, dto.ErrAccountAlreadyVerified
	}

	updatedUser, err := s.userRepo.UpdateUser(ctx, nil, entity.User{
		ID:         user.ID,
		IsVerified: true,
	})
	if err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrUpdateUser
	}

	return dto.VerifyEmailResponse{
		Email:      email,
		IsVerified: updatedUser.IsVerified,
	}, nil
}

func (s *userService) GetUserById(ctx context.Context, userId string) (dto.UserResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserById
	}

	var ocId, ocStatus, wsnId, wsnStatus *string

	return dto.UserResponse{
		Name:        user.Name,
		Email:       user.Email,
		Instansi:    user.Instansi,
		NoTelp:      user.NoTelp,
		InfoFrom:    user.InfoFrom,
		Jenjang:     user.Jenjang,
		OCId:        ocId,
		OCStatus:    ocStatus,
		WSNId:       wsnId,
		WSNStatus:   wsnStatus,
		ILJMajor:    user.ILJMajor,
		ILJSubmajor: user.ILJSubmajor,
		Role:        user.Role,
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, req dto.UserUpdateRequest, userId string) (dto.UserUpdateResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.UserUpdateResponse{}, dto.ErrUserNotFound
	}

	provinceID, err := uuid.Parse(req.Asal_Provinsi)
	if err != nil {
		return dto.UserUpdateResponse{}, dto.ErrParseUUID
	}

	data := entity.User{
		ID:         user.ID,
		Name:       req.Name,
		Email:      user.Email,
		Instansi:   req.Instansi,
		NoTelp:     req.No_Telp,
		InfoFrom:   req.Info_From,
		ProvinceID: provinceID,
		Jenjang:    req.Jenjang,
	}

	userUpdate, err := s.userRepo.UpdateUser(ctx, nil, data)
	if err != nil {
		return dto.UserUpdateResponse{}, dto.ErrUpdateUser
	}

	return dto.UserUpdateResponse{
		Name:      userUpdate.Name,
		Email:     userUpdate.Email,
		Instansi:  userUpdate.Instansi,
		No_Telp:   userUpdate.NoTelp,
		Info_From: userUpdate.InfoFrom,
		Jenjang:   userUpdate.Jenjang,
	}, nil
}

func (s *userService) DeleteUser(ctx context.Context, userId string) error {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.ErrUserNotFound
	}

	err = s.userRepo.DeleteUser(ctx, nil, user.ID.String())
	if err != nil {
		return dto.ErrDeleteUser
	}

	return nil
}

func (s *userService) Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	check, flag, err := s.userRepo.CheckEmail(ctx, nil, req.Email)
	if err != nil || !flag {
		return dto.UserLoginResponse{}, dto.ErrInvalidCredentials
	}

	if !check.IsVerified {
		return dto.UserLoginResponse{}, dto.ErrAccountNotVerified
	}

	checkPassword, err := helpers.CheckPassword(check.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		return dto.UserLoginResponse{}, dto.ErrInvalidCredentials
	}

	user, err := s.userRepo.GetUserByEmail(ctx, nil, check.Email)
	if err != nil {
		return dto.UserLoginResponse{}, dto.ErrGetUserByEmail
	}

	token := s.jwtService.GenerateTokenUser(user.ID.String(), user.Role)

	return dto.UserLoginResponse{
		Token: token,
		Role:  check.Role,
	}, nil
}

func (s *userService) ResetPassword(ctx context.Context, token string, newPassword string) error {
	decryptedToken, err := utils.AESDecrypt(token)
	if err != nil {
		return dto.ErrTokenInvalid
	}

	tokenParts := strings.Split(decryptedToken, "_")
	if len(tokenParts) < 2 {
		return dto.ErrTokenInvalid
	}

	email := tokenParts[0]
	expirationDate := tokenParts[1]
	expirationTime, err := time.Parse("2006-01-02 15:04:05", expirationDate)

	if err != nil {
		return dto.ErrTokenInvalid
	}

	// time.Now().After(expirationTime)
	// timeDifference := math.Abs(time.Since(expirationTime).Seconds())
	// fmt.Println(timeDifference)
	// fmt.Println(expirationTime)

	if time.Now().After(expirationTime) {
		return dto.ErrTokenExpired
	}
	hashedPassword, err := helpers.HashPassword(newPassword)

	if err != nil {
		return dto.ErrHashPasswordFailed
	}

	err = s.userRepo.ResetPassword(ctx, email, hashedPassword)
	if err != nil {
		return dto.ErrUpdateUser
	}

	return nil
}

func (s *userService) MakeForgetPasswordEmail(receiverEmail string) (map[string]string, error) {
	expired := time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")
	plainText := receiverEmail + "_" + expired
	token, err := utils.AESEncrypt(plainText)
	if err != nil {
		return nil, err
	}

	local_url := os.Getenv("APP_URL")
	verifyLink := local_url + "/" + RESET_EMAIL_ROUTE + "?token=" + token

	readHtml, err := os.ReadFile("utils/email-template/verification_email.html")
	if err != nil {
		return nil, err
	}

	data := struct {
		Email  string
		Verify string
	}{
		Email:  receiverEmail,
		Verify: verifyLink,
	}

	tmpl, err := template.New("custom").Parse(string(readHtml))
	if err != nil {
		return nil, err
	}

	var strMail bytes.Buffer
	if err := tmpl.Execute(&strMail, data); err != nil {
		return nil, err
	}

	draftEmail := map[string]string{
		"subject": "Reset Password",
		"body":    strMail.String(),
	}

	return draftEmail, nil
}

func (s *userService) ForgetPassword(ctx context.Context, req dto.ForgetPasswordRequest) error {
	user, err := s.userRepo.GetUserByEmail(ctx, nil, req.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return dto.ErrUserNotFound
		}
		return err
	}
	draftEmail, err := s.MakeForgetPasswordEmail(user.Email)
	if err != nil {
		return err
	}
	err = utils.SendMail(user.Email, draftEmail["subject"], draftEmail["body"])
	if err != nil {
		return err
	}
	return nil
}
