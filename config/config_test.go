package config

import (
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	os.Setenv("SERVER_ADDR", "asdasdasdadasdasd")
	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		{
			name: "ttt",
			args: args{opts: []Option{WithGRPCServer("qweqweqw", "qweqwe")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
