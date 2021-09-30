package schedules

import (
	"context"
	"database/sql"
	"time"
)

type Domain struct {
	ID          uint
	DoctorId    uint
	Room        string
	Message     string
	TanggalJaga time.Time
	JamAwal     string
	JamAkhir    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}

// type Domain struct {
// 	ID          uint
// 	TanggalJaga time.Time
// 	JamAwal     string
// 	JamAkhir    string
// 	CreatedAt   time.Time
// 	UpdatedAt   time.Time
// 	DeletedAt   gorm.DeletedAt
// }

type Usecase interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	Remove(ctx context.Context, domain Domain) (Domain, error)
	Modificate(ctx context.Context, domain Domain) (Domain, error)
	Show(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) (Domain, error)
	// InsertDoctor(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	Remove(ctx context.Context, domain Domain) (Domain, error)
	Modificate(ctx context.Context, domain Domain) (Domain, error)
	Show(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context, domain Domain) (Domain, error)
	// InsertDoctor(ctx context.Context, domain Domain) (Domain, error)
}
