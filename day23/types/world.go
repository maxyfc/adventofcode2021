package types

import (
	"fmt"
	"strings"
)

type World struct {
	pods     []*Pod
	hallways [numHallways]*Pod
	rooms    [numRooms]*Room
	roomSize int
}

func (w *World) EnergyUsed() int {
	e := 0
	for _, p := range w.pods {
		e += p.energyUsed
	}
	return e
}

func (w *World) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf(`
#############
#%s%s %s %s %s %s%s#
###%s#%s#%s#%s###`,
		w.hallways[0],
		w.hallways[1],
		w.hallways[2],
		w.hallways[3],
		w.hallways[4],
		w.hallways[5],
		w.hallways[6],
		w.rooms[0].pos[0],
		w.rooms[1].pos[0],
		w.rooms[2].pos[0],
		w.rooms[3].pos[0],
	))

	for p := 1; p < w.roomSize; p++ {
		s.WriteString(fmt.Sprintf("\n  #%s#%s#%s#%s#",
			w.rooms[0].pos[p],
			w.rooms[1].pos[p],
			w.rooms[2].pos[p],
			w.rooms[3].pos[p]))
	}

	s.WriteString("\n  #########\n")
	s.WriteString(fmt.Sprintf("Energy Used: %d\n", w.EnergyUsed()))

	return s.String()
}

type Room struct {
	pos []*Pod
	t   PodType
}

type Pod struct {
	t          PodType
	energyUsed int
}

func (p *Pod) String() string {
	if p == nil {
		return " "
	}
	return p.t.String()
}

const (
	PodTypeA PodType = iota
	PodTypeB
	PodTypeC
	PodTypeD
)

type PodType int

func (t PodType) String() string {
	switch t {
	case PodTypeA:
		return "A"
	case PodTypeB:
		return "B"
	case PodTypeC:
		return "C"
	case PodTypeD:
		return "D"
	default:
		panic(fmt.Sprintf("Invalid pod type: %d", t))
	}
}

// Number of rooms in the world
const numRooms = 4

// Number of hallway positions in the world
const numHallways = 7

var roomHallwaySteps map[int]map[int]int = map[int]map[int]int{
	0: { // Room 0
		0: 3, // Hallway 0 - 3 steps
		1: 2,
		2: 2,
		3: 4,
		4: 6,
		5: 8,
		6: 9,
	},
	1: { // Room 1
		0: 5,
		1: 4,
		2: 2,
		3: 2,
		4: 4,
		5: 6,
		6: 7,
	},
	2: { // Room 2
		0: 7,
		1: 6,
		2: 4,
		3: 2,
		4: 2,
		5: 4,
		6: 5,
	},
	3: { // Room 3
		0: 9,
		1: 8,
		2: 6,
		3: 4,
		4: 2,
		5: 2,
		6: 3,
	},
}

func NewWorld(roomSize int, podTypes ...PodType) *World {
	if roomSize < 1 {
		panic(fmt.Sprintf("Room size must be more than zero. Got: %d", roomSize))
	}
	if len(podTypes)/roomSize != numRooms {
		panic(fmt.Sprintf("There should be a multiple of 4 pods. Got: %d", len(podTypes)))
	}

	pods := createPods(podTypes)
	return &World{
		pods:     pods,
		roomSize: roomSize,
		rooms:    createRooms(roomSize, pods),
	}
}

func createPods(podTypes []PodType) []*Pod {
	var pods []*Pod
	for _, t := range podTypes {
		pods = append(pods, &Pod{t: t})
	}
	return pods
}

func createRooms(size int, pods []*Pod) (rooms [numRooms]*Room) {
	for i := 0; i < numRooms; i++ {
		rooms[i] = &Room{
			pos: make([]*Pod, 0, size),
			t:   PodType(i),
		}
		for j := i * size; j < (i+1)*size; j++ {
			rooms[i].pos = append(rooms[i].pos, pods[j])
		}
	}
	return
}
