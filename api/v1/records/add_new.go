package records_api

import "github.com/gin-gonic/gin"

// Add new record godoc
// @Summary Add a new record (doctor)
// @Description Add a new record for a patient
// @Tags records
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /api/v1/records [post]
func AddNewRecordHandler(c *gin.Context) {

}
