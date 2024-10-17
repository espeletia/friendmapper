package usecases

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"hradec/internal/domain"
	"hradec/internal/ports"
	"regexp"

	"go.uber.org/zap"
)

func NewUserUsecase(usi ports.UserDatabaseStore,
	salt string,
) *UserUsecase {
	return &UserUsecase{
		store:      usi,
		salt:       salt,
		emailRegex: *regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`),
	}
}

type UserUsecase struct {
	store          ports.UserDatabaseStore
	salt           string
	validatorURL   string
	emailRegex     regexp.Regexp
	skipValidation bool
}

func (uu *UserUsecase) CreateUser(ctx context.Context, user domain.UserData, password string) (*domain.User, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infof("Creating user with email: %s", user.Email)
	hashedPassword := uu.HashPassword(password)
	if !uu.emailRegex.MatchString(user.Email) {
		return nil, domain.ErrInvalidEmail
	}
	user.Hash = hashedPassword
	stored, err := uu.store.CreateUser(ctx, user, uu.skipValidation)
	if err != nil {
		sugar.Infof("Error whilst creating user")
		return nil, err
	}
	return stored, nil
}

func (uu *UserUsecase) GetUserById(ctx context.Context, Id int64) (*domain.User, error) {
	return uu.store.GetUserById(ctx, Id)
}

func (uu *UserUsecase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return uu.store.GetUserByEmail(ctx, email)
}

func (uu *UserUsecase) HashPassword(password string) string {
	passwordBytes := []byte(password)
	hash := sha512.New()
	saltBytes := []byte(uu.salt)
	passwordBytes = append(passwordBytes, saltBytes...)
	hash.Write(passwordBytes)
	hashedPasswordBytes := hash.Sum(nil)
	hashedPasswordHex := hex.EncodeToString(hashedPasswordBytes)
	return hashedPasswordHex
}
