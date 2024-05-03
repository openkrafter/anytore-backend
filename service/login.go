package service

import (
	"context"
	"encoding/hex"
	"os"
	"strconv"
	"time"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/google/uuid"
	"github.com/openkrafter/anytore-backend/auth"
	"github.com/openkrafter/anytore-backend/logger"
)

func Login(ctx context.Context, email string, password string) (string, error) {
	user, err := GetUserByEmail(ctx, email)
	if err != nil {
		logger.Logger.Error("Failed to get user by email", err)
		return "", err
	}

	err = auth.PassHasher.CompareHashAndPassword(user.Password, password)
	if err != nil {
		logger.Logger.Error("Failed to compare password", err)
		return "", err
	}

	token, err := generateToken(1)
	if err != nil {
		logger.Logger.Error("Failed to generate token", err)
		return "", err
	}

	return token, nil
}

func generateToken(userId int) (string, error) {
	// key := []byte(os.Getenv("TOKEN_SECRET"))
	keyString := os.Getenv("TOKEN_SECRET")
	key, err := hex.DecodeString(keyString)
	if err != nil {
		logger.Logger.Error("Failed to decode key", err)
		return "", err
	}

	encrypter, err := jose.NewEncrypter(
		jose.A256GCM,
		jose.Recipient{Algorithm: jose.DIRECT, Key: key},
		(&jose.EncrypterOptions{}).WithType("JWT"))
	if err != nil {
		logger.Logger.Error("Failed to create encrypter", err)
		return "", err
	}

	userIdString := strconv.Itoa(userId)
	claims := jwt.Claims{
		Subject:   userIdString,
		Issuer:    "anytore",
		Audience:  []string{"anytore"},
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Expiry:    jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		ID:        uuid.New().String(),
	}

	encryptedToken, err := jwt.Encrypted(encrypter).Claims(claims).Serialize()
	if err != nil {
		logger.Logger.Error("Failed to generate token", err)
		return "", err
	}

	return encryptedToken, nil
}
