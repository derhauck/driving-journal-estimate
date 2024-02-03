package logger

import (
	"os"
	"testing"
)

func TestParseLevel(t *testing.T) {
	type args struct {
		level string
	}
	tests := []struct {
		name    string
		args    args
		want    Level
		wantErr bool
	}{
		{name: "test debug",
			args: args{
				level: levels[DEBUG],
			},
			want:    DEBUG,
			wantErr: false,
		},
		{
			name: "test info",
			args: args{
				level: levels[INFO],
			},
			want:    INFO,
			wantErr: false,
		},
		{
			name: "test warn",
			args: args{
				level: levels[WARNING],
			},
			want:    WARNING,
			wantErr: false,
		},
		{
			name: "test error",
			args: args{
				level: levels[ERROR],
			},
			want:    ERROR,
			wantErr: false,
		},
		{
			name: "test invalid level",
			args: args{
				level: levels[9999999],
			},
			want:    DEFAULT,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLevel(tt.args.level)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseLevel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_bootstrapLogging(t *testing.T) {
	type fields struct {
		level         Level
		defaultWriter *os.File
	}
	type args struct {
		level  Level
		v      any
		result bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test debug", fields: fields{
				level:         DEFAULT,
				defaultWriter: os.Stdout,
			},
			args: args{
				level:  DEBUG,
				v:      "test debug",
				result: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &logger{
				level:         tt.fields.level,
				defaultWriter: tt.fields.defaultWriter,
			}
			if result := l.bootstrapLogging(tt.args.level, tt.args.v); result != tt.args.result {
				t.Errorf("bootstrapLogging() = %v, want %v, MSG LEVEL:%s, LOGGER LEVL:%s", result, tt.args.result, tt.args.level, tt.fields.level)
			}

		})
	}
}
