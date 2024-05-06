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
	"github.com/openkrafter/anytore-backend/model"
)

func Login(ctx context.Context, email string, password string) (string, *model.SafeUser, error) {
	user, err := GetUserByEmail(ctx, email)
	if err != nil {
		logger.Logger.Error("Failed to get user by email", err)
		return "", &model.SafeUser{}, err
	}

	err = auth.PassHasher.CompareHashAndPassword(user.Password, password)
	if err != nil {
		logger.Logger.Error("Failed to compare password", err)
		return "", &model.SafeUser{}, err
	}

	token, err := generateToken(user.Id)
	if err != nil {
		logger.Logger.Error("Failed to generate token", err)
		return "", &model.SafeUser{}, err
	}

	safeUser := &model.SafeUser{
		Name:  user.Name,
		Email: user.Email,
	}
	return token, safeUser, nil
}

func ValidateToken(token string) (int, error) {
	keyString := os.Getenv("TOKEN_SECRET")
	key, err := hex.DecodeString(keyString)
	if err != nil {
		logger.Logger.Error("Failed to decode key", err)
		return 0, err
	}

	parsedToken, err := jwt.ParseEncrypted(
		token,
		[]jose.KeyAlgorithm{jose.DIRECT},
		[]jose.ContentEncryption{jose.A256GCM})
	if err != nil {
		logger.Logger.Error("Failed to parse encrypted token", err)
		return 0, err
	}

	claims := jwt.Claims{}
	if err := parsedToken.Claims(key, &claims); err != nil {
		logger.Logger.Error("Failed to get claims from token", err)
		return 0, err
	}

	userId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		logger.Logger.Error("Failed to convert userId string to int", err)
		return 0, err
	}

	return userId, nil
}

func generateToken(userId int) (string, error) {
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
