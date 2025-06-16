package statistic_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	statistic_services "github.com/merema-uit/server/services/statistic"
	"github.com/merema-uit/server/utils"
)

// Compile Record Statistic by Time godoc
// @Summary Compile Record Statistic by time (admin)
// @Description Compile record statistics based on time
// @Tags statistic
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param RecordStatisticRequest body models.RecordStatisticRequest true "Record Statistic Request"
// @Success 200 {array} models.AmountOfRecordsByTime
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /statistic/records [post]
func CompileRecordStatisticByTimeHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	var req models.RecordStatisticRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	res, err := statistic_services.CompileRecordStatistic(c, authHeader, req, "time")
	if err != nil {
		switch err {
		case errs.ErrInvalidCompileType, errs.ErrInvalidTimeUnit, errs.ErrInvalidTimestamp:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errs.ErrInvalidToken, errs.ErrExpiredToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Can't compile record statistic", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, res)
}

// Compile Record Statistic by Doctor godoc
// @Summary Compile Record Statistic by Doctor (admin)
// @Description Compile record statistics based on doctor and time
// @Tags statistic
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param RecordStatisticRequest body models.RecordStatisticRequest true "Record Statistic Request"
// @Success 200 {array} models.AmountOfRecordsByDoctor
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /statistic/records/doctor [post]
func CompileRecordStatisticByDoctorHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	var req models.RecordStatisticRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	res, err := statistic_services.CompileRecordStatistic(c, authHeader, req, "doctor")
	if err != nil {
		switch err {
		case errs.ErrInvalidCompileType, errs.ErrInvalidTimeUnit, errs.ErrInvalidTimestamp:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errs.ErrInvalidToken, errs.ErrExpiredToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Can't compile record statistic", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, res)
}

// Compile Record Statistic by Diagnosis godoc
// @Summary Compile Record Statistic by Diagnosis (admin)
// @Description Compile record statistics based on diagnosis and time
// @Tags statistic
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param RecordStatisticRequest body models.RecordStatisticRequest true "Record Statistic Request"
// @Success 200 {array} models.AmountOfRecordsByDiagnosis
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /statistic/records/diagnosis [post]
func CompileRecordStatisticByDiagnosisHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	var req models.RecordStatisticRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	res, err := statistic_services.CompileRecordStatistic(c, authHeader, req, "diagnosis")
	if err != nil {
		switch err {
		case errs.ErrInvalidCompileType, errs.ErrInvalidTimeUnit, errs.ErrInvalidTimestamp:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errs.ErrInvalidToken, errs.ErrExpiredToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Can't compile record statistic", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, res)
}
