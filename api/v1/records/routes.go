package records_api

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.POST("", AddNewRecordHandler)
	r.POST("/:record_id/attachments", AddRecordAttachmentsHandler)
	r.GET("", GetRecordListHandler)
	r.GET("/:recordID", GetRecordInfoHandler)
	r.GET("/:recordID/attachments", GetRecordAttachmentsHandler)
}
