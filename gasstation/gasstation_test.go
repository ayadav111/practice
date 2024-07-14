package gasstation

import "testing"

func Test_canCompleteStatation(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "canCompleteStation",
			args: args{
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			want: 3,
		},
		{
			name: "cantCompleteStation",
			args: args{
				gas:  []int{2, 3, 4},
				cost: []int{3, 4, 3},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCompleteStatation(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("canCompleteStatation() = %v, want %v", got, tt.want)
			}
		})
	}
}
