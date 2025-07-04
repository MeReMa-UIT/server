package account_api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	registration_services "github.com/merema-uit/server/services/registration"
	"github.com/merema-uit/server/utils"
)

// Initiate registration godoc
// @Summary Initiate registration (admin, receptionist)
// @Description Check whether the citizen ID is already registered
// @Tags accounts
// @Accept json
// @Produce json
// @Param user body models.InitRegistrationRequest true "Initiate registration data"
// @Security BearerAuth
// @Success 200 {object} models.AccountRegistrationResponse
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /accounts/register [post]
func InitRegistrationHandler(ctx *gin.Context) {
	var req models.InitRegistrationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	response, err := registration_services.InitRegistration(context.Background(), req, authHeader)

	if err != nil {
		switch err {
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Can't initialize user registration", "error", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Register account godoc
// @Summary Register new account (admin, receptionist)
// @Description Create a new account
// @Tags accounts
// @Accept json
// @Produce json
// @Param user body models.AccountRegistrationRequest true "User registration data"
// @Security BearerAuth
// @Success 201 {object} models.AccountRegistrationResponse
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 409
// @Failure 500
// @Router /accounts/register/create [post]
func RegisterAccountHandler(ctx *gin.Context) {
	var req models.AccountRegistrationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	authHeader := ctx.GetHeader("Authorization")

	response, err := registration_services.RegisterAccount(context.Background(), req, authHeader)

	if err != nil {
		switch err {
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrCitizenIDExists:
			ctx.JSON(http.StatusConflict, gin.H{"error": "Citizen ID already exists"})
		case errs.ErrEmailOrPhoneAlreadyUsed:
			ctx.JSON(http.StatusConflict, gin.H{"error": "Email or phone number already used"})
		default:
			utils.Logger.Error("Can't register new account", "error", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

// Register patient godoc
// @Summary Register new patient (receptionist)
// @Description Create a new patient account
// @Tags accounts
// @Accept json
// @Produce json
// @Param user body models.PatientRegistrationRequest true "Patient registration data"
// @Security BearerAuth
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /accounts/register/patients [post]
func RegisterPatientHandler(ctx *gin.Context) {
	var req models.PatientRegistrationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	authHeader := ctx.GetHeader("Authorization")

	err := registration_services.RegisterPatient(context.Background(), req, authHeader)

	if err != nil {
		switch err {
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Can't register new patient", "error", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Patient registered successfully"})
}

// Register staff godoc
// @Summary Register new staff (admin)
// @Description Create a new staff account
// @Tags accounts
// @Accept json
// @Produce json
// @Param user body models.StaffRegistrationRequest true "Staff registration data"
// @Security BearerAuth
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 409
// @Failure 500
// @Router /accounts/register/staffs [post]
func RegisterStaffHandler(ctx *gin.Context) {
	var req models.StaffRegistrationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	authHeader := ctx.GetHeader("Authorization")

	err := registration_services.RegisterStaff(context.Background(), req, authHeader)

	if err != nil {
		switch err {
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrAccountAlreadyLinked:
			ctx.JSON(http.StatusConflict, gin.H{"error": "Account already linked with other staff"})
		default:
			utils.Logger.Error("Can't register new staff", "error", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Staff registered successfully"})
}
