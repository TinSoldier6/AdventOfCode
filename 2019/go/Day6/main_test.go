package main

import (
	"io"
	"reflect"
	"testing"
)

// Test set 1
// COM)B
// B)C
// C)D
// D)E
// E)F
// B)G
// G)H
// D)I
// E)J
// J)K
// K)L
// jumps = 42

// Test set 2
// COM)B
// B)C
// C)D
// D)E
// E)F
// B)G
// G)H
// D)I
// E)J
// J)K
// K)L
// K)YOU
// I)SAN
// jumps = 4

var part1Example = []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}

func Test_parseOrbits(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{
			name: "Part 1 Example Data",
			args: args{input: part1Example},
			want: [][]string{{"COM", "B"}, {"B", "C"}, {"C", "D"}, {"D", "E"}, {"E", "F"}, {"B", "G"}, {"G", "H"}, {"D", "I"}, {"E", "J"}, {"J", "K"}, {"K", "L"}},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseOrbits(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseOrbits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newOrbitalSystem(t *testing.T) {
	type args struct {
		orbits [][]string
	}
	tests := []struct {
		name       string
		args       args
		wantSystem orbitalSystem
	}{
		// TODO: Add test cases.
		{
			name:       "Part 1 Example Data",
			args:       args{parseOrbits(part1Example)},
			wantSystem: orbitalSystem{"B": "COM", "C": "B", "D": "C", "E": "D", "F": "E", "G": "B", "H": "G", "I": "D", "J": "E", "K": "J", "L": "K", "COM": ""},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSystem := newOrbitalSystem(tt.args.orbits); !reflect.DeepEqual(gotSystem, tt.wantSystem) {
				t.Errorf("newOrbitalSystem() = %v, want %v", gotSystem, tt.wantSystem)
			}
		})
	}
}

func Test_orbitalSystem_getPath(t *testing.T) {
	type args struct {
		to, from string
	}
	tests := []struct {
		name     string
		o        orbitalSystem
		args     args
		wantPath []string
	}{
		// TODO: Add test cases.
		{
			name:     "Part 1 Example Data",
			o:        orbitalSystem{"B": "COM", "C": "B", "D": "C", "E": "D", "F": "E", "G": "B", "H": "G", "I": "D", "J": "E", "K": "J", "L": "K", "COM": ""},
			args:     args{to: "COM", from: "I"},
			wantPath: []string{"D", "C", "B", "COM"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPath := tt.o.getPath(tt.args.to, tt.args.from); !reflect.DeepEqual(gotPath, tt.wantPath) {
				t.Errorf("orbitalSystem.getPath() = %v, want %v", gotPath, tt.wantPath)
			}
		})
	}
}

// var part2Example = []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "k)YOU", "I)SAN"}

func Test_countAllJumps(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Part 1 Example Data",
			args: args{input: part1Example},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAllJumps(tt.args.input); got != tt.want {
				t.Errorf("countAllJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_readLines(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name      string
		args      args
		wantLines []string
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLines, err := readLines(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("readLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLines, tt.wantLines) {
				t.Errorf("readLines() = %v, want %v", gotLines, tt.wantLines)
			}
		})
	}
}

func Test_check(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check(tt.args.err)
		})
	}
}
