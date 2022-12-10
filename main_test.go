package windirection

import (
	"testing"
)

func TestDegreesToCardinal(t *testing.T) {
	type args struct {
		degrees float32
		points  int8
		lang    string
	}
	tests := []struct {
		name    string
		args    args
		want    *Cardinal
		wantErr bool
	}{
		{
			name: "Verify with invalid low value",
			args: args{
				degrees: -0.1,
				points:  16,
				lang:    "es",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Verify with invalid high value",
			args: args{
				degrees: 360.01,
				points:  16,
				lang:    "es",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid high points number",
			args: args{
				degrees: 0,
				points:  17,
				lang:    "es",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid low points number",
			args: args{
				degrees: 0,
				points:  0,
				lang:    "es",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid points number valids are only 4, 8, 16",
			args: args{
				degrees: 0,
				points:  12,
				lang:    "es",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Valid value with north 16 ES",
			args: args{
				degrees: 360,
				points:  16,
				lang:    "es",
			},
			want:    &Cardinal{"N", "Norte"},
			wantErr: false,
		},
		{
			name: "Valid value with north 8 ES",
			args: args{
				degrees: 350,
				points:  8,
				lang:    "es",
			},
			want:    &Cardinal{"N", "Norte"},
			wantErr: false,
		},
		{
			name: "Valid value with north 4 ES",
			args: args{
				degrees: 10,
				points:  4,
				lang:    "es",
			},
			want:    &Cardinal{"N", "Norte"},
			wantErr: false,
		},
		{
			name: "Valid translation with north 4 EN",
			args: args{
				degrees: 10,
				points:  4,
				lang:    "english",
			},
			want:    &Cardinal{"N", "North"},
			wantErr: false,
		},
		{
			name: "Valid value with south 16 ES",
			args: args{
				degrees: 185,
				points:  16,
				lang:    "es",
			},
			want:    &Cardinal{"S", "Sur"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DegreesToCardinal(tt.args.degrees, tt.args.points, tt.args.lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("DegreesToCardinal() error = %v, wantErr %v", err, tt.wantErr)
				if tt.want != nil {
					t.Errorf("DegreesToCardinal() cardinal must be nil")
				}
				return
			}
			// Compare values returned by cardinal
			if err == nil && (got.Short != tt.want.Short || got.Long != tt.want.Long) {
				t.Errorf("DegreesToCardinal() = %v, want %v", got, tt.want)
			}
		})
	}
}
