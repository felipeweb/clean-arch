package entity

import (
	"os"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPort_UnmarshalJSON(t *testing.T) {

	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		info    PortInfo
		Key     string
	}{
		{
			name: "test unmarshal port",
			args: args{
				filepath: "../testdata/port.json",
			},
			wantErr: false,
			info: PortInfo{
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
			Key: "AEAJM",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Port{}
			b, err := os.ReadFile(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Port.UnmarshalJSON() read error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := p.UnmarshalJSON(b); (err != nil) != tt.wantErr {
				t.Errorf("Port.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(p.PortInfo, tt.info) {
				t.Errorf("Port.UnmarshalJSON() Info = %v, want %v", p.PortInfo, tt.info)
			}
			if p.Key != tt.Key {
				t.Errorf("Port.UnmarshalJSON() Key = %v, want %v", p.Key, tt.Key)
			}
		})
	}
}

func TestPort_MarshalJSON(t *testing.T) {
	type fields struct {
		Key      string
		PortInfo PortInfo
	}
	tests := []struct {
		name     string
		fields   fields
		filepath string
		wantErr  bool
	}{
		{
			name: "test marshal port",
			fields: fields{
				Key: "AEAJM",
				PortInfo: PortInfo{
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
			filepath: "../testdata/port.json",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Port{
				Key:      tt.fields.Key,
				PortInfo: tt.fields.PortInfo,
			}
			got, err := p.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Port.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			b, err := os.ReadFile(tt.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Port.UnmarshalJSON() read error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gs := Port{}
			bs := Port{}
			if err := gs.UnmarshalJSON(got); err != nil {
				t.Errorf("Port.UnmarshalJSON() error = %v", err)
				return
			}
			if err := bs.UnmarshalJSON(b); err != nil {
				t.Errorf("Port.UnmarshalJSON() error = %v", err)
				return
			}
			diff := cmp.Diff(gs, bs)
			if diff != "" {
				t.Errorf("Port.MarshalJSON() = %v", diff)
			}
		})
	}
}
