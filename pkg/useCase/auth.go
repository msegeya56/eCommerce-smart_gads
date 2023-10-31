package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/msegeya56/eCommerce-smart_gads/pkg/domain"
	"github.com/msegeya56/eCommerce-smart_gads/pkg/repository/interfaces"
	service "github.com/msegeya56/eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/msegeya56/eCommerce-smart_gads/pkg/verify"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	AuthRepository interfaces.AuthRepository
	userRepository interfaces.UserRepository
}

func NewAuthUseCase(repo interfaces.AuthRepository, userRepo interfaces.UserRepository) service.AuthService {
	return &AuthUseCase{AuthRepository: repo,
		userRepository: userRepo,
	}
}

func (u *AuthUseCase) SignUp(ctx context.Context, user domain.Users) error {
	// Check if user already exist
	DBUser, err := u.userRepository.FindUser(ctx, user)

	if err != nil {
		return err
	}

	if DBUser.ID == 0 {
		// Hash user password
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			fmt.Println("Hashing failed")
			return err
		}
		user.Password = string(hashedPass)

		// Save user if not exist
		err = u.AuthRepository.SaveUser(ctx, user)
		if err != nil {
			return errors.New("failed to save user")
		}

	} else {
		return fmt.Errorf("%v user already exists", DBUser.UserName)
	}

	return nil
}

func (u *AuthUseCase) Login(ctx context.Context, user domain.Users) (domain.Users, error) {
	// Find user in db
	DBUser, err := u.userRepository.FindUser(ctx, user)
	if err != nil {
		return user, err
	} else if DBUser.ID == 0 {
		return user, errors.New("user not exist")
	}
	// Check if the user blocked by admin
	if DBUser.BlockStatus {
		return user, errors.New("user blocked by admin")
	}

	if _, err := verify.TwilioSendOTP("+91" + DBUser.Phone); err != nil {
		// response := response.ErrorResponse(500, "failed to send otp", err.Error(), nil)
		// c.JSON(http.StatusInternalServerError, response)
		return user, fmt.Errorf("failed to send otp %v",
			err)
	}
	// check password with hashed pass
	if bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(user.Password)) != nil {
		return user, errors.New("password incorrect")
	}

	return DBUser, nil

}
func (u *AuthUseCase) OTPLogin(ctx context.Context, user domain.Users) (domain.Users, error) {
	// Find user in db
	DBUser, err := u.userRepository.FindUser(ctx, user)
	if err != nil {
		return user, err
	} else if DBUser.ID == 0 {
		return user, errors.New("user not exist")
	}
	return DBUser, nil
}
