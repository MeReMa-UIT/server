package permission

type Permission int

const (
	Admin Permission = iota
	Receptionist
	Doctor
	Patient
	Recovery
	PatientRegistration
	StaffRegistration
)

func (d Permission) String() string {
	return [...]string{
		"admin",
		"receptionist",
		"doctor",
		"patient",
		"recovery",
		"patient_registration",
		"staff_registration"}[d]
}
