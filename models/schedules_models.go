package models

import "time"

type ScheduleBookingRequest struct {
	ExaminationDate time.Time `json:"examination_date"`
	// Type of examination (1: Regular, 2: Service)
	Type int `json:"type" enums:"1,2"`
}

type ScheduleBookingResponse struct {
	ExaminationDate time.Time `json:"examination_date" db:"examination_date"`
	// Type of examination (1: Regular, 2: Service)
	Type                  int    `json:"type" db:"type" enums:"1,2"`
	QueueNumber           int    `json:"queue_number" db:"queue_number"`
	ExpectedReceptionTime string `json:"expected_reception_time" db:"expected_reception_time"`
	// Status of the schedule (1: Waiting, 2: Completed, 3: Cancelled)
	Status int `json:"status" db:"status" enums:"1,2,3"`
}

type ScheduleInfo struct {
	ScheduleID int `json:"schedule_id" db:"schedule_id"`
	ScheduleBookingResponse
}

type GetScheduleListRequest struct {
	// The list of types of examination (1: Regular, 2: Service)
	Type []int `form:"type[]"`
	// The list of statuses of the schedule (1: Waiting, 2: Completed, 3: Cancelled)
	Status []int `form:"status[]"`
}
