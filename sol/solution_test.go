package sol

import "testing"

func Test_smashStones(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[2, 10, 4, 2, 20, 3]",
			args: args{
				stones: []int{2, 10, 4, 2, 20, 3},
			},
			want: 1,
		},
		{
			name: " [100, 1, 2]",
			args: args{
				stones: []int{100, 1, 2},
			},
			want: 97,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smashStones(tt.args.stones); got != tt.want {
				t.Errorf("smashStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
