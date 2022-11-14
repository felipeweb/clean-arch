package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/felipeweb/clean-arch/entity"
	"github.com/felipeweb/clean-arch/repository"
)

func TestPortService_Save(t *testing.T) {
	type args struct {
		key  string
		info *entity.PortInfo
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.Port
		wantErr bool
	}{
		{
			name: "test save port",
			args: args{
				info: &entity.PortInfo{
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
				key: "AEAJM",
			},
			want: &entity.Port{
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PortService{
				repo: repository.NewInMemory(),
			}
			got, err := s.Save(context.Background(), tt.args.key, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("PortService.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PortService.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}
