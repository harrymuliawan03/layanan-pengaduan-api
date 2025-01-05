package authservice

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/configs"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/dto"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/http/requests"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/models"
	authverifrepo "github.com/harrymuliawan03/layanan-pengaduan-api.git/app/repositories/auth_verif_repo"
	userrepo "github.com/harrymuliawan03/layanan-pengaduan-api.git/app/repositories/user_repo"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/schemas"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/utils"
	// "github.com/golang-jwt/jwt/v5"
)

type AuthServiceImpl struct {
	userRepository userrepo.UserRepository
	cnf            *configs.Config
	authverifRepo  authverifrepo.AuthVerifRepository
}

func NewAuthService(ur userrepo.UserRepository, cnf *configs.Config, avr authverifrepo.AuthVerifRepository) AuthService {
	return &AuthServiceImpl{userRepository: ur, cnf: cnf, authverifRepo: avr}
}

// Login implements AuthService.
func (u *AuthServiceImpl) Login(ctx context.Context, req *requests.LoginRequest) (dto.LoginResponse, error) {
	res, err := u.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}
	checkPassword := utils.CheckPasswordHash(req.Password, res.Password)

	if !checkPassword {
		return dto.LoginResponse{}, errors.New("invalid email or password")
	}
	fmt.Println(res)

	// claim := jwt.MapClaims{
	// 	"id":  res.ID,
	// 	"exp": time.Now().Add(time.Duration(time.Hour * 24)).Unix(),
	// }
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// tokenStr, err := token.SignedString([]byte(u.cnf.Jwt.Key))

	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{Valid: true, Name: res.Name, Email: res.Email}, nil
}

// Create implements UserService.
func (u *AuthServiceImpl) Register(ctx context.Context, req *requests.RegisterRequest) error {
	data := models.User{
		Email:     req.Email,
		Name:      req.Name,
		Password:  req.Password,
	}

	err := u.userRepository.Save(ctx, &data)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements UserService.
func (u *AuthServiceImpl) Delete(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// Show implements UserService.
func (u *AuthServiceImpl) Show(ctx context.Context, id uint) (*dto.UserData, error) {
	// Fetch user from the database
	user, err := u.userRepository.FindByID(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	// Prepare response data (convert your user model to DTO)
	userData := &dto.UserData{
		Email:     user.Email,
		// Add other fields as necessary
	}

	return userData, nil
}

// Update implements UserService.
func (u *AuthServiceImpl) Update(ctx context.Context, req requests.UserUpdateRequest) error {
	panic("unimplemented")
}

func (u *AuthServiceImpl) VerifEmail(ctx context.Context, req requests.VerifEmailRequest, code string) error {
	c, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	fmt.Printf("code: %s\n", code)

	var av models.AuthVerif
	av, err := u.authverifRepo.FindByCode(c, code)
	if err != nil {
		return err
	}

	if av.IsUsed {
		return &schemas.ResponseApiError{
			Status:   schemas.ApiErrorBadRequest,
			Message: "Verification link is used",
		}
	}

	av.IsUsed = true
	err = u.authverifRepo.Update(c, &av)
	if err != nil {
		return err
	}
	
	err = u.userRepository.Update(c, models.User{ID: av.UserID})
	if err != nil {
		return err
	}

	err = u.authverifRepo.Delete(c, av.ID)
	if err != nil {
		return err
	}

	return nil
	
}
