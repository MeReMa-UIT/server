package statistic_services

import (
	"context"
	"slices"
	"time"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
)

func CompileRecordStatistic(ctx context.Context, authHeader string, req models.RecordStatisticRequest, compileType string) (any, error) {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}

	if claims.Permission != permission.Admin.String() {
		return nil, errs.ErrPermissionDenied
	}

	if ok := slices.Contains([]string{"time", "doctor", "diagnosis"}, compileType); !ok {
		return nil, errs.ErrInvalidCompileType
	}

	var timestampStart, timestampEnd time.Time
	switch req.TimeUnit {
	case "day":
		year, month, day := req.Timestamp.Date()
		timestampStart = time.Date(year, month, day, 0, 0, 0, 0, req.Timestamp.UTC().Location())
		timestampEnd = timestampStart.AddDate(0, 0, 1)
	case "week":
		if req.Timestamp.Weekday() != time.Monday {
			return nil, errs.ErrInvalidTimestamp
		}
		year, month, day := req.Timestamp.Date()
		timestampStart = time.Date(year, month, day, 0, 0, 0, 0, req.Timestamp.UTC().Location())
		timestampEnd = timestampStart.AddDate(0, 0, 7)
	case "month":
		month := req.Timestamp.Month()
		year := req.Timestamp.Year()
		timestampStart = time.Date(year, month, 1, 0, 0, 0, 0, req.Timestamp.UTC().Location())
		timestampEnd = timestampStart.AddDate(0, 1, 0)
	case "year":
		year := req.Timestamp.Year()
		timestampStart = time.Date(year, 1, 1, 0, 0, 0, 0, req.Timestamp.UTC().Location())
		timestampEnd = timestampStart.AddDate(1, 0, 0)
	default:
		return nil, errs.ErrInvalidTimeUnit
	}

	recordList, err := repo.GetRecordListByTime(ctx, timestampStart, timestampEnd)
	if err != nil {
		return nil, err
	}

	switch compileType {
	case "time":
		return compileRecordStatisticByTime(recordList, timestampStart, timestampEnd, req.TimeUnit)
	case "doctor":
		return compileRecordStatisticByDoctor(ctx, recordList, timestampStart, timestampEnd, req.TimeUnit)
	case "diagnosis":
		return compileRecordStatisticByRecordType(ctx, recordList, timestampStart, timestampEnd, req.TimeUnit)
	default:
		return nil, nil
	}

}

func compileRecordStatisticByTime(recordList []models.RecordInfoForStatistic, timestampStart, timestampEnd time.Time, timeUnit string) ([]models.AmountOfRecordsByTime, error) {
	var ret []models.AmountOfRecordsByTime
	switch timeUnit {
	case "day":
		ret = append(ret, models.AmountOfRecordsByTime{
			TimestampStart: timestampStart,
			Amount:         len(recordList),
		})
	case "week":
		ret = make([]models.AmountOfRecordsByTime, 7)
		for i := 0; i < 7; i++ {
			ret[i] = models.AmountOfRecordsByTime{
				TimestampStart: timestampStart.AddDate(0, 0, i),
				Amount:         0,
			}
		}
		for _, record := range recordList {
			dayOfWeek := int(record.CreatedAt.Weekday())
			if dayOfWeek == 0 {
				dayOfWeek = 7
			}
			ret[dayOfWeek-1].Amount++
		}
	case "month":
		j := 0
		for i := timestampStart; i.Before(timestampEnd); i = i.AddDate(0, 0, 1) {
			if i.Weekday() == time.Monday || len(ret) == 0 {
				j++
				ret = append(ret, models.AmountOfRecordsByTime{
					TimestampStart: i,
					Amount:         0,
				})
			}
		}

		// i = (record.CreatedAt.Weekday() + firstWeekdadyOfMonth - 1) // 7
		firstWeekdayOfMonth := timestampStart.Weekday()
		for _, record := range recordList {
			i := (int(record.CreatedAt.Day()) + int(firstWeekdayOfMonth) - 1) / 7
			ret[i].Amount++
		}
	case "year":
		ret = make([]models.AmountOfRecordsByTime, 12)
		for i := 0; i < 12; i++ {
			ret[i] = models.AmountOfRecordsByTime{
				TimestampStart: timestampStart.AddDate(0, i, 0),
				Amount:         0,
			}
		}
		for _, record := range recordList {
			ret[int(record.CreatedAt.Month())-1].Amount++
		}
	}
	return ret, nil
}

func compileRecordStatisticByDoctor(ctx context.Context, recordList []models.RecordInfoForStatistic, timestampStart, timestampEnd time.Time, timeUnit string) ([]models.AmountOfRecordsByDoctor, error) {
	classifiedRecordList := make(map[int64][]models.RecordInfoForStatistic)
	for _, record := range recordList {
		classifiedRecordList[record.DoctorID] = append(classifiedRecordList[record.DoctorID], record)
	}

	var ret []models.AmountOfRecordsByDoctor
	for doctorID, records := range classifiedRecordList {
		list, err := compileRecordStatisticByTime(records, timestampStart, timestampEnd, timeUnit)
		if err != nil {
			return nil, err
		}
		ret = append(ret, models.AmountOfRecordsByDoctor{
			DoctorID:     doctorID,
			AmountByTime: list,
		})
	}

	doctorList, err := repo.GetDoctorList(ctx)
	if err != nil {
		return nil, err
	}

	for _, doctorID := range doctorList {
		if _, ok := classifiedRecordList[doctorID]; !ok {
			ret = append(ret, models.AmountOfRecordsByDoctor{
				DoctorID:     doctorID,
				AmountByTime: nil,
			})
		}
	}

	return ret, nil
}
func compileRecordStatisticByRecordType(ctx context.Context, recordList []models.RecordInfoForStatistic, timestampStart, timestampEnd time.Time, timeUnit string) ([]models.AmountOfRecordsByDiagnosis, error) {
	classifiedRecordList := make(map[string][]models.RecordInfoForStatistic)
	for _, record := range recordList {
		classifiedRecordList[record.PrimaryDiagnosis] = append(classifiedRecordList[record.PrimaryDiagnosis], record)
	}

	var ret []models.AmountOfRecordsByDiagnosis
	for diagnosisID, records := range classifiedRecordList {
		list, err := compileRecordStatisticByTime(records, timestampStart, timestampEnd, timeUnit)
		if err != nil {
			return nil, err
		}
		ret = append(ret, models.AmountOfRecordsByDiagnosis{
			DiagnosisID:  diagnosisID,
			AmountByTime: list,
		})
	}

	diagnosisList, err := repo.GetDiagnosisList(ctx)
	if err != nil {
		return nil, err
	}

	for _, diagnosis := range diagnosisList {
		if _, ok := classifiedRecordList[diagnosis.ICDCode]; !ok {
			ret = append(ret, models.AmountOfRecordsByDiagnosis{
				DiagnosisID:  diagnosis.ICDCode,
				AmountByTime: nil,
			})
		}
	}

	return ret, nil
}
