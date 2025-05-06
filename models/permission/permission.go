package permission

type Permission int

const (
	Admin Permission = iota
	Receptionist
	Doctor
	Patient
	Recovery
)

func (d Permission) String() string {
	return [...]string{"admin", "receptionist", "doctor", "patient", "recovery"}[d]
}
