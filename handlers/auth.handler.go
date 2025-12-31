package handlers

import (
	"crypto/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/malachi190/watchcircle/helper"
	"github.com/malachi190/watchcircle/logger"
	"github.com/malachi190/watchcircle/models"
	"github.com/malachi190/watchcircle/repository"
	"github.com/malachi190/watchcircle/types"
	"github.com/malachi190/watchcircle/validator"
	"github.com/tendermint/crypto/bcrypt"
)

type AuthHandler struct {
	Repo *repository.Repo
}

func NewAuthHandler(repo *repository.Repo) *AuthHandler {
	return &AuthHandler{
		Repo: repo,
	}
}

// Register
func (h *AuthHandler) Register(ctx *gin.Context) {
	var reqBody types.RegisterPayload

	if errors, ok := validator.Body(ctx, &reqBody); !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": false,
			"errors": errors,
		})
		return
	}

	// generate otp code
	code, err := helper.GenerateOtpCode()

	if err != nil {
		logger.Error.Printf("OTP code generation failed: %v\n", err)
		ctx.JSON(500, gin.H{
			"status":  false,
			"message": "Something went wrong, please try again",
		})
		return
	}

	// store in redis cache
	key := reqBody.Email

	if err := h.Repo.Redis.Set(ctx, key, code, 10*time.Minute).Err(); err != nil {
		logger.Error.Printf("Failed to store OTP code in redis: %v\n", err)
		ctx.JSON(500, gin.H{
			"status":  false,
			"message": "Something went wrong, please try again",
		})
		return
	}

	// Hash password
	salt := make([]byte, 16)

	if _, err := rand.Read(salt); err != nil {
		logger.Error.Printf("Error creating random bytes: %v\n", err)
		ctx.JSON(500, gin.H{
			"status":  false,
			"message": "Something went wrong, please try again",
		})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword(salt, []byte(reqBody.Password), 10)

	if err != nil {
		logger.Error.Printf("Error hashing password: %v\n", err)
		ctx.JSON(500, gin.H{
			"status":  false,
			"message": "Something went wrong, please try again",
		})
		return
	}

	// store user in db
	user := &models.UserModel{
		Email:       reqBody.Email,
		DisplayName: reqBody.DisplayName,
		Username:    reqBody.Username,
		Password:    string(hashed),
	}

	if err := h.Repo.Service.Auth.CreateUser(ctx, user); err != nil {
		logger.Error.Printf("Error while creating account: %v\n", err)
		ctx.JSON(500, gin.H{
			"status":  false,
			"message": "Failed to create account.",
		})
		return
	}

	// send otp mail
	type mailData struct {
		DisplayName string
		OTP         string
	}

	_, err = h.Repo.Mailer.SendTemplate("welcome", mailData{DisplayName: reqBody.DisplayName, OTP: code}, reqBody.Email, "Verify Email")

	if err != nil {
		logger.Error.Printf("Error while sending otp code: %v\n", err)
	}

	// return
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Account created successfully",
		"data":    user,
	})
}

// VerifyEmail verifies a users email address
func (h *AuthHandler) VerifyEmail(ctx *gin.Context) {
	var reqBody types.VerifyEmailPayload

	if errors, ok := validator.Body(ctx, &reqBody); !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": false,
			"errors": errors,
		})
		return
	}

	// get code from redis store
	value, err := h.Repo.Redis.Get(ctx, reqBody.Email).Result()

	if err != nil {
		logger.Error.Printf("error while retrieving otp: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid email address provided",
		})
		return
	}

	// check if code has expired
	ttl := h.Repo.Redis.TTL(ctx, reqBody.Email).Val()

	if ttl == -2 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "OTP has expired, please request a new one.",
		})
		return
	}

	if value != reqBody.OTP {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid OTP provided",
		})
		return
	}

	now := time.Now().UTC()

	usr := &models.UserModel{
		Email:           reqBody.Email,
		EmailVerifiedAt: &now,
	}

	// set user as verified
	if err := h.Repo.Service.Auth.UpdateUserVerification(ctx, usr); err != nil {
		logger.Error.Printf("error updating user model: %v", err)
		ctx.JSON(500, gin.H{
			"status":  false,
			"message": "Something went wrong, please try again",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Email verified successfully.",
	})
}

// ResendOtp resends otp code to the user
func (h *AuthHandler) ResendOtp(ctx *gin.Context) {
	var reqBody types.ResendEmailPayload

	// generate code
	code, err := helper.GenerateOtpCode()

	if err != nil {
		logger.Error.Printf("OTP code generation failed: %v\n", err)
		ctx.JSON(500, gin.H{
			"status":  false,
			"message": "Something went wrong, please try again",
		})
		return
	}

	// store in redis cache
	key := reqBody.Email

	if err := h.Repo.Redis.Set(ctx, key, code, 10*time.Minute).Err(); err != nil {
		logger.Error.Printf("Failed to store OTP code in redis: %v\n", err)
		ctx.JSON(500, gin.H{
			"status":  false,
			"message": "Something went wrong, please try again",
		})
		return
	}

	// send email
	type mailData struct {
		OTP string
	}

	_, err = h.Repo.Mailer.SendTemplate("resend", mailData{OTP: code}, reqBody.Email, "Verify Email")

	if err != nil {
		logger.Error.Printf("Error while sending otp code: %v\n", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "OTP code resent successfully.",
	})
}

// Login, creates a new session for the user
func (h *AuthHandler) Login(ctx *gin.Context) {
	var reqBody types.LoginPayload

	if errors, ok := validator.Body(ctx, &reqBody); !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": false,
			"errors": errors,
		})
		return
	}

	// check users credentials
	user := &models.UserModel{
		Email: reqBody.Email,
	}

	if err := h.Repo.Service.Auth.GetUser(ctx, user); err != nil {
		logger.Error.Printf("database error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid email address provided",
		})
		return
	}

	// check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid email or password",
		})
		return
	}

	// generate jwt access token
	accessToken, err1 := helper.GenerateAccessToken(user, h.Repo.Config)
	refresToken, err2 := helper.GenerateRefreshToken(user, h.Repo.Config)

	if err1 != nil || err2 != nil {
		logger.Error.Printf("jwt token errors: %v:%v", err1, err2)
		ctx.JSON(500, gin.H{
			"status":  false,
			"message": "Something went wrong, please try again",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":        true,
		"message":       "Login successful",
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refresToken,
	})
}

// FetchRefreshToken, generates and returns a new refresh token
func (h *AuthHandler) FetchRefreshToken(ctx *gin.Context) {
	// get refresh token from request header
	token := ctx.Request.Header.Get("X-Refresh-Token")

	decoded, err := helper.DecodeJwt(token, []byte(h.Repo.Config.Jwt.RefreshSecret))

	if err != nil {
		logger.Error.Printf("error decoding jwt: %v", err)
		ctx.JSON(500, gin.H{
			"status":  false,
			"message": "Something went wrong, please try again",
		})
		return
	}

	userId, _ := decoded.GetSubject()

	id, _ := uuid.Parse(userId)

	user := &models.UserModel{
		Id: id,
	}

	if err := h.Repo.Service.Auth.GetUserById(ctx, user); err != nil {
		logger.Error.Printf("error getting user by id: %v", err)
		ctx.JSON(500, gin.H{
			"status":  false,
			"message": "Something went wrong, please try again",
		})
		return
	}

	// generate a new refresh token
	newToken, err := helper.GenerateRefreshToken(user, h.Repo.Config)

	if err != nil {
		logger.Error.Printf("jwt token error: %v", err)
		ctx.JSON(500, gin.H{
			"status":  false,
			"message": "Something went wrong, please try again",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":        true,
		"refresh_token": newToken,
	})
}
