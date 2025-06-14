package record_api

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.POST("", AddNewRecordHandler)
	r.GET("", GetRecordListHandler)
	r.GET("/:record_id", GetRecordInfoHandler)
	r.PUT("/:record_id", UpdateMedicalRecordHandler)
	r.GET("/:record_id/attachments", GetRecordAttachmentsHandler)
	r.POST("/:record_id/attachments", AddRecordAttachmentsHandler)
	r.DELETE("/:record_id/attachments", DeleteRecordAttachmentHandler)
}
