package convert

import (
	"testing"
)

func BenchmarkStrToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrToInt[int]("2147483647")
	}
}

func BenchmarkStrToInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrToInt[int64]("9223372036854775807")
	}
}

func BenchmarkStrToUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrToInt[uint]("4294967295")
	}
}

func BenchmarkStrToUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrToInt[uint64]("18446744073709551615")
	}
}

func BenchmarkStrToUintptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrToInt[uintptr]("18446744073709551615")
	}
}

func TestStrToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "pos",
			args: args{"2147483647"},
			want: 2147483647,
		},

		{
			name: "neg",
			args: args{"-2147483648"},
			want: -2147483648,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToInt[int](tt.args.s); got != tt.want {
				t.Errorf("StrToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToInt64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "pos",
			args: args{"9223372036854775807"},
			want: 9223372036854775807,
		},

		{
			name: "neg",
			args: args{"-9223372036854775808"},
			want: -9223372036854775808,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToInt[int64](tt.args.s); got != tt.want {
				t.Errorf("StrToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastStrToUint(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "pos",
			args: args{"4294967295"},
			want: 4294967295,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToInt[uint](tt.args.s); got != tt.want {
				t.Errorf("StrToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastStrToUint64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "pos",
			args: args{"18446744073709551615"},
			want: 18446744073709551615,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToInt[uint64](tt.args.s); got != tt.want {
				t.Errorf("StrToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastStrToUintptr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want uintptr
	}{
		{
			name: "pos",
			args: args{"18446744073709551615"},
			want: 18446744073709551615,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToInt[uintptr](tt.args.s); got != tt.want {
				t.Errorf("StrToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
