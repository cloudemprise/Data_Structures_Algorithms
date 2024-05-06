package main

import (
	"reflect"
	"testing"
)

func TestReadBinaryWatch(t *testing.T) {

	checkResult := func(t *testing.T, got, want map[string]bool) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("turnedOn=1", func(t *testing.T) {
		turnedOn := 1

		want := map[string]bool{
			"0:01": true,
			"0:02": true,
			"0:04": true,
			"0:08": true,
			"0:16": true,
			"0:32": true,
			"1:00": true,
			"2:00": true,
			"4:00": true,
			"8:00": true,
		}

		got := make(map[string]bool, len(want))
		for _, str := range readBinaryWatch(turnedOn) {
			got[str] = true
		}

		checkResult(t, got, want)
	})

	t.Run("turnedOn=9", func(t *testing.T) {
		turnedOn := 9

		want := map[string]bool{}

		got := make(map[string]bool, len(want))
		for _, str := range readBinaryWatch(turnedOn) {
			got[str] = true
		}

		checkResult(t, got, want)
	})

	t.Run("turnedOn=2", func(t *testing.T) {
		turnedOn := 2

		want := map[string]bool{
			"0:03":  true,
			"0:05":  true,
			"0:06":  true,
			"0:09":  true,
			"0:10":  true,
			"0:12":  true,
			"0:17":  true,
			"0:18":  true,
			"0:20":  true,
			"0:24":  true,
			"0:33":  true,
			"0:34":  true,
			"0:36":  true,
			"0:40":  true,
			"0:48":  true,
			"1:01":  true,
			"1:02":  true,
			"1:04":  true,
			"1:08":  true,
			"1:16":  true,
			"1:32":  true,
			"2:01":  true,
			"2:02":  true,
			"2:04":  true,
			"2:08":  true,
			"2:16":  true,
			"2:32":  true,
			"3:00":  true,
			"4:01":  true,
			"4:02":  true,
			"4:04":  true,
			"4:08":  true,
			"4:16":  true,
			"4:32":  true,
			"5:00":  true,
			"6:00":  true,
			"8:01":  true,
			"8:02":  true,
			"8:04":  true,
			"8:08":  true,
			"8:16":  true,
			"8:32":  true,
			"9:00":  true,
			"10:00": true,
		}

		got := make(map[string]bool, len(want))
		for _, str := range readBinaryWatch(turnedOn) {
			got[str] = true
		}

		checkResult(t, got, want)
	})

}
