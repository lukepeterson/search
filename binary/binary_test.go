package binary

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		data   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty input",
			args: args{data: []int{}, target: 1},
			want: -1,
		},
		{
			name: "single element, found",
			args: args{data: []int{4}, target: 4},
			want: 0,
		},
		{
			name: "single element, not found",
			args: args{data: []int{4}, target: 1},
			want: -1,
		},
		{
			name: "multiple elements, found",
			args: args{data: []int{5, 6, 7, 8}, target: 7},
			want: 2,
		},
		{
			name: "multiple elements, not found",
			args: args{data: []int{5, 6, 7, 8}, target: 10},
			want: -1,
		},
		{
			name: "negative elements, found",
			args: args{data: []int{-2, -3, -4, -5}, target: -3},
			want: 1,
		},
		{
			name: "negative elements, not found",
			args: args{data: []int{-2, -3, -4, -5}, target: -9},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.data, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
