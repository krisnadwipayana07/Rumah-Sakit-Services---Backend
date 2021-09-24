package doctors

type DoctorLogin struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type DoctorUpdate struct {
	Id            int    `validate:"required"`
	Email         string `validate:"required"`
	Password      string `validate:"required"`
	Name          string `validate:"required"`
	Address       string `validate:"required"`
	Nip           string `validate:"required"`
	DoctorJob     string `validate:"required"`
	Token         string `validate:"required"`
	Description   string `validate:"required"`
	ContactPerson string `validate:"required"`
}
