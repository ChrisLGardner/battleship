package main

import (
	"testing"

	"github.com/chrislgardner/battleship/package/battleship"
)

func Test_playerFire(t *testing.T) {
	type args struct {
		shot string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "valid shot",
			args:    args{"1,1"},
			want:    "hot",
			wantErr: false,
		},
		{
			name:    "too many elements",
			args:    args{"1,1,1"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "too few elements",
			args:    args{"11"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "invalid x",
			args:    args{"X,1"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "invalid y",
			args:    args{"1,Y"},
			want:    "",
			wantErr: true,
		},
	}
	battleship.NewGame(8, 8)
	battleship.Ship1 = battleship.Position{X: 2, Y: 2}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := playerFire(tt.args.shot)
			if (err != nil) != tt.wantErr {
				t.Errorf("playerFire() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("playerFire() = %v, want %v", got, tt.want)
			}
		})
	}
}
