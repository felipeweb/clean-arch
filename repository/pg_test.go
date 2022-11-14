package repository

import (
	"context"
	"testing"

	"github.com/felipeweb/clean-arch/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"zombiezen.com/go/postgrestest"
)

func TestPG_Save(t *testing.T) {

	type args struct {
		port *entity.Port
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		rows    int
	}{
		{
			name: "test save port on postgres",
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
	srv, err := postgrestest.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(srv.Cleanup)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			godb, err := srv.NewDatabase(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("PG.Save() godb error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			db, err := gorm.Open(postgres.New(postgres.Config{
				Conn: godb,
			}), &gorm.Config{})
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
