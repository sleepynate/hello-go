package main

import "testing"

func TestJustIan(t *testing.T) {
	got := hasIan("ian")
	if !got {
		t.Errorf("ian should be true")
	}
}

func Test_hasIan(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasIan(tt.args.input); got != tt.want {
				t.Errorf("hasIan() = %v, want %v", got, tt.want)
			}
		})
	}
}