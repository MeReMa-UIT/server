package permission

type Permission int

const (
	Admin Permission = iota
	Receptionist
	Doctor
	Patient
)

func (d Permission) String() string {
	return [...]string{"admin", "receptionist", "doctor", "patient"}[d]
}
