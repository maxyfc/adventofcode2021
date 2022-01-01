package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRoomAllowEntry(t *testing.T) {
	podA := &Pod{t: PodTypeA}
	podB := &Pod{t: PodTypeB}

	tests := []struct {
		r        *Room
		expected bool
	}{
		{
			&Room{pos: []*Pod{nil, nil, nil}, t: PodTypeA},
			true,
		},
		{
			&Room{pos: []*Pod{nil, nil, podA}, t: PodTypeA},
			true,
		},
		{
			&Room{pos: []*Pod{podA, podA, podA}, t: PodTypeA},
			false,
		},
		{
			&Room{pos: []*Pod{podA, podB, podA}, t: PodTypeA},
			false,
		},
		{
			&Room{pos: []*Pod{podB, podA, podA}, t: PodTypeA},
			false,
		},
		{
			&Room{pos: []*Pod{podA, podA, podB}, t: PodTypeA},
			false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test%03d", i), func(t *testing.T) {
			output := test.r.AllowEntry()
			if test.expected != output {
				t.Errorf("Expected: %v Got: %v", test.expected, output)
			}
		})
	}
}

func TestRoomRemovePod(t *testing.T) {
	podA := &Pod{t: PodTypeA}
	podB := &Pod{t: PodTypeB}
	podC := &Pod{t: PodTypeC}

	tests := []struct {
		r        *Room
		expected []*Pod
	}{
		{
			&Room{pos: []*Pod{nil, nil, nil}, t: PodTypeA},
			[]*Pod{},
		},
		{
			&Room{pos: []*Pod{nil, nil, podA}, t: PodTypeA},
			[]*Pod{podA},
		},
		{
			&Room{pos: []*Pod{nil, podB, podA}, t: PodTypeA},
			[]*Pod{podB, podA},
		},
		{
			&Room{pos: []*Pod{podC, podB, podA}, t: PodTypeA},
			[]*Pod{podC, podB, podA},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test%03d", i), func(t *testing.T) {
			var pods []*Pod
			for {
				pod := test.r.RemovePod()
				if pod == nil {
					break
				}
				pods = append(pods, pod)
			}

			if len(test.expected) != len(pods) {
				t.Fatalf("Expected length: %d Got: %d", len(test.expected), len(pods))
			}

			for i := 0; i < len(test.expected); i++ {
				if !reflect.DeepEqual(test.expected[i], pods[i]) {
					t.Errorf("i: %d Expected: %v Got: %v", i, test.expected[i], pods[i])
				}
			}
		})
	}
}

func TestRoomAddPod(t *testing.T) {
	podA := &Pod{t: PodTypeA}
	podB := &Pod{t: PodTypeB}
	podC := &Pod{t: PodTypeC}
	podD := &Pod{t: PodTypeD}

	tests := []struct {
		r        *Room
		p        *Pod
		expected int
	}{
		{
			&Room{pos: []*Pod{nil, nil, nil}, t: PodTypeA},
			podA,
			2,
		},
		{
			&Room{pos: []*Pod{nil, nil, podA}, t: PodTypeA},
			podB,
			1,
		},
		{
			&Room{pos: []*Pod{nil, podB, podA}, t: PodTypeA},
			podC,
			0,
		},
		{
			&Room{pos: []*Pod{podC, podA, podB}, t: PodTypeA},
			podD,
			-1,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test%03d", i), func(t *testing.T) {
			output := test.r.AddPod(test.p)
			if test.expected != output {
				t.Fatalf("Expected: %d Got: %d", test.expected, output)
			}

			if output >= 0 {
				pod := test.r.RemovePod()
				if !reflect.DeepEqual(test.p, pod) {
					t.Errorf("Expected remove: %v Got: %v", test.p, pod)
				}
			}
		})
	}
}

func TestRoomNextPod(t *testing.T) {
	podA := &Pod{t: PodTypeA}
	podB := &Pod{t: PodTypeB}

	tests := []struct {
		r        *Room
		expected *Pod
		pos      int
	}{
		{
			&Room{pos: []*Pod{nil, nil, nil}, t: PodTypeA},
			nil, -1,
		},
		{
			&Room{pos: []*Pod{nil, nil, podA}, t: PodTypeA},
			nil, -1,
		},
		{
			&Room{pos: []*Pod{podA, podA, podA}, t: PodTypeA},
			nil, -1,
		},
		{
			&Room{pos: []*Pod{nil, nil, podB}, t: PodTypeA},
			podB, 2,
		},
		{
			&Room{pos: []*Pod{nil, podB, podA}, t: PodTypeA},
			podB, 1,
		},
		{
			&Room{pos: []*Pod{nil, podA, podB}, t: PodTypeA},
			podA, 1,
		},
		{
			&Room{pos: []*Pod{podA, podB, podB}, t: PodTypeA},
			podA, 0,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test%03d", i), func(t *testing.T) {
			pod, pos := test.r.NextPod()
			if test.expected != pod {
				t.Errorf("Expected pod: %v Got: %v", test.expected, pod)
			}
			if test.pos != pos {
				t.Errorf("Expected positon: %d Got: %d", test.pos, pos)
			}
		})
	}
}

func TestRoomNextFreePos(t *testing.T) {
	podA := &Pod{t: PodTypeA}
	podB := &Pod{t: PodTypeB}
	podC := &Pod{t: PodTypeC}

	tests := []struct {
		r        *Room
		expected int
	}{
		{
			&Room{pos: []*Pod{nil, nil, nil}, t: PodTypeA},
			2,
		},
		{
			&Room{pos: []*Pod{nil, nil, podA}, t: PodTypeA},
			1,
		},
		{
			&Room{pos: []*Pod{nil, podB, podA}, t: PodTypeA},
			0,
		},
		{
			&Room{pos: []*Pod{podC, podA, podB}, t: PodTypeA},
			-1,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test%03d", i), func(t *testing.T) {
			output := test.r.NextFreePos()
			if test.expected != output {
				t.Fatalf("Expected: %d Got: %d", test.expected, output)
			}
		})
	}
}
