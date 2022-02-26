package battleship

import (
	"testing"
)

func TestPlayerMove(t *testing.T) {
	type args struct {
		move  Position
		ship1 Position
		ship2 Position
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Ship 1 is hit",
			args: args{
				move:  Position{1, 1},
				ship1: Position{1, 1},
				ship2: Position{5, 5},
			},
			want: "hit",
		},
		{
			name: "Ship 2 is hit",
			args: args{
				move:  Position{1, 1},
				ship1: Position{5, 5},
				ship2: Position{1, 1},
			},
			want: "hit",
		},
		{
			name: "Ship 1 is closer and hot",
			args: args{
				move:  Position{1, 1},
				ship1: Position{1, 2},
				ship2: Position{5, 5},
			},
			want: "hot",
		},
		{
			name: "Ship 1 is closer and warm",
			args: args{
				move:  Position{1, 1},
				ship1: Position{1, 4},
				ship2: Position{5, 5},
			},
			want: "warm",
		},
		{
			name: "Ship 1 is closer and cold",
			args: args{
				move:  Position{1, 1},
				ship1: Position{1, 6},
				ship2: Position{5, 5},
			},
			want: "cold",
		},
		{
			name: "Ship 2 is closer and hot",
			args: args{
				move:  Position{1, 1},
				ship1: Position{5, 5},
				ship2: Position{1, 2},
			},
			want: "hot",
		},
		{
			name: "Ship 2 is closer and warm",
			args: args{
				move:  Position{1, 1},
				ship1: Position{5, 5},
				ship2: Position{1, 4},
			},
			want: "warm",
		},
		{
			name: "Ship 2 is closer and cold",
			args: args{
				move:  Position{1, 1},
				ship1: Position{5, 5},
				ship2: Position{1, 6},
			},
			want: "cold",
		},
		{
			name: "Ship 1 is closer and hot, ship 2 is already hit",
			args: args{
				move:  Position{1, 1},
				ship1: Position{1, 2},
				ship2: Position{99, 99},
			},
			want: "hot",
		},
		{
			name: "Player move x is out of bounds",
			args: args{
				move:  Position{9, 1},
				ship1: Position{1, 2},
				ship2: Position{99, 99},
			},
			want: "out",
		},
		{
			name: "Player move y is out of bounds",
			args: args{
				move:  Position{1, 9},
				ship1: Position{1, 2},
				ship2: Position{99, 99},
			},
			want: "out",
		},
	}

	NewGame(8, 8)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Ship1 = tt.args.ship1
			Ship2 = tt.args.ship2
			if got := PlayerMove(tt.args.move); got != tt.want {
				t.Errorf("PlayerMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
