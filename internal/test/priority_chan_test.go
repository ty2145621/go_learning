package test

import (
	"testing"
)

func initTest() {
	Init()
}

func initTest(pos int) {
	initTest()

}

func Test_popHeartBeat(t *testing.T) {
	tests := []struct {
		name          string
		wantHeartBeat int
	}{
		{name: "1", wantHeartBeat: 1},
		{name: "2", wantHeartBeat: 2},
		{name: "3", wantHeartBeat: 3},
		{name: "4", wantHeartBeat: 4},
		{name: "5", wantHeartBeat: 5},
		{name: "6", wantHeartBeat: 6},
		{name: "7", wantHeartBeat: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHeartBeat := popHeartBeat(); gotHeartBeat != tt.wantHeartBeat {
				t.Errorf("popHeartBeat() = %v, want %v", gotHeartBeat, tt.wantHeartBeat)
			}
		})
	}
}

func Test_popHeartBeatV2(t *testing.T) {
	tests := []struct {
		name          string
		wantHeartBeat int
	}{
		{name: "1", wantHeartBeat: 1},
		{name: "2", wantHeartBeat: 2},
		{name: "3", wantHeartBeat: 3},
		{name: "4", wantHeartBeat: 4},
		{name: "5", wantHeartBeat: 5},
		{name: "6", wantHeartBeat: 6},
		{name: "7", wantHeartBeat: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHeartBeat := popHeartBeatV2(); gotHeartBeat != tt.wantHeartBeat {
				t.Errorf("popHeartBeatV2() = %v, want %v", gotHeartBeat, tt.wantHeartBeat)
			}
		})
	}
}
