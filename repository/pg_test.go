package repository

import (
	"context"
	"testing"

	"github.com/felipeweb/clean-arch/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestPG_Save(t *testing.T) {
	type args struct {
		port *entity.Port
	}
	tests := []struct {
		name    string
		dsn     string
		args    args
		wantErr bool
		rows    int
	}{
		{
			name: "test save port on postgres",
			dsn:  "postgresql://postgres:postgres@localhost:5432/testdb?sslmode=disable",
			args: args{
				port: &entity.Port{
					Key: "AEAJM",
					PortInfo: entity.PortInfo{
						Name:        "Ajman",
						City:        "Ajman",
						Country:     "United Arab Emirates",
						Alias:       []string{},
						Regions:     []string{},
						Coordinates: []float64{55.5136433, 25.4052165},
						Province:    "Ajman",
						Timezone:    "Asia/Dubai",
						Unlocs:      []string{"AEAJM"},
						Code:        "52000",
					},
				},
			},
			wantErr: false,
			rows:    1,
		},
	}
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			db, err := gorm.Open(postgres.Open(tt.dsn), &gorm.Config{})
			if (err != nil) != tt.wantErr {
				t.Errorf("PG.Save() gorm db error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			p := NewPG(db)
			if err := p.Save(ctx, tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("PG.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			ports := []entity.Port{}
			db.Find(&ports)
			if len(ports) != tt.rows {
				t.Errorf("PG.Save() rows = %v, want %v", len(ports), tt.rows)
			}
		})
	}
}
