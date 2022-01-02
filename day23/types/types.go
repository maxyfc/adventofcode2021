package types

import (
	"fmt"
	"strings"
)

type World struct {
	hallways [numHallways]*Pod
	rooms    [numRooms]*Room
	roomSize int
}

func (w *World) IsSolved() bool {
	for _, r := range w.rooms {
		if !r.IsSolved() {
			return false
		}
	}
	return true
}

func (w *World) NextMoves() []*Move {
	var moves []*Move

	// Move from room to hallway
	for ri, room := range w.rooms {
		pod, rpos := room.NextPod()
		if pod == nil {
			continue
		}

		for _, hallways := range [][]int{roomMoveLeft[ri], roomMoveRight[ri]} {
			for _, hpos := range hallways {
				if w.hallways[hpos] != nil {
					break
				}

				cost := (roomHallwaySteps[ri][hpos] + rpos) * podStepCosts[pod.t]
				moves = append(moves, &Move{
					t:          MoveTypeRoomToHallway,
					roomIndex:  ri,
					roomPos:    rpos,
					hallwayPos: hpos,
					cost:       cost,
				})
			}
		}
	}

	// Move from hallway to room
	for hpos, pod := range w.hallways {
		if pod == nil {
			continue
		}

		ri := podRoom[pod.t]
		if !w.allEmpty(hpos, ri) {
			continue
		}

		room := w.rooms[ri]
		if !room.AllowEntry() {
			continue
		}

		rpos := room.NextFreePos()
		cost := (roomHallwaySteps[ri][hpos] + rpos) * podStepCosts[pod.t]
		moves = append(moves, &Move{
			t:          MoveTypeHallwayToRoom,
			roomIndex:  ri,
			roomPos:    rpos,
			hallwayPos: hpos,
			cost:       cost,
		})
	}

	return moves
}

func (w *World) allEmpty(hallwayPos, room int) bool {
	for _, hpos := range hallwayRoomPassThrough[hallwayPos][room] {
		if w.hallways[hpos] != nil {
			return false
		}
	}
	return true
}

func (w *World) Apply(m *Move) {
	room := w.rooms[m.roomIndex]
	if m.t == MoveTypeRoomToHallway {
		pod := room.RemovePod()
		w.hallways[m.hallwayPos] = pod
	} else {
		pod := w.hallways[m.hallwayPos]
		w.hallways[m.hallwayPos] = nil
		room.AddPod(pod)
	}
}

func (w *World) Copy() *World {
	copy := &World{
		roomSize: w.roomSize,
	}
	for i, p := range w.hallways {
		if p != nil {
			copy.hallways[i] = p.Copy()
		}
	}
	for ri, r := range w.rooms {
		copy.rooms[ri] = r.Copy()
	}
	return copy
}

func (w *World) CacheKey() string {
	var s strings.Builder
	s.WriteString("h:")
	for _, p := range w.hallways {
		if p == nil {
			s.WriteByte('-')
		} else {
			s.WriteString(p.t.String())
		}
	}
	for ri, r := range w.rooms {
		s.WriteString(fmt.Sprintf(",r%d:", ri))
		for _, p := range r.pos {
			if p == nil {
				s.WriteByte('-')
			} else {
				s.WriteString(p.t.String())
			}
		}
	}
	return s.String()
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

	return s.String()
}

type Room struct {
	pos []*Pod
	t   PodType
}

func (r *Room) AllowEntry() bool {
	for _, p := range r.pos {
		if p != nil && p.t != r.t {
			return false
		}
	}
	return r.pos[0] == nil
}

func (r *Room) RemovePod() (p *Pod) {
	for i := 0; i < len(r.pos); i++ {
		if r.pos[i] == nil {
			continue
		}
		p = r.pos[i]
		r.pos[i] = nil
		break
	}
	return
}

func (r *Room) AddPod(p *Pod) int {
	if p == nil {
		panic("Cannot add nil")
	}

	for i := len(r.pos) - 1; i >= 0; i-- {
		if r.pos[i] != nil {
			continue
		}
		r.pos[i] = p
		return i
	}
	return -1
}

func (r *Room) NextPod() (*Pod, int) {
	var i int
	for i = len(r.pos) - 1; i >= 0; i-- {
		if r.pos[i] == nil || r.pos[i].t != r.t {
			break
		}
	}

	for j, p := range r.pos {
		if j > i {
			break
		}
		if p == nil {
			continue
		}
		return p, j
	}
	return nil, -1
}

func (r *Room) NextFreePos() int {
	for i := len(r.pos) - 1; i >= 0; i-- {
		if r.pos[i] != nil {
			continue
		}
		return i
	}
	return -1
}

func (r *Room) Copy() *Room {
	copy := &Room{
		t:   r.t,
		pos: make([]*Pod, len(r.pos)),
	}
	for i, p := range r.pos {
		if p != nil {
			copy.pos[i] = p.Copy()
		}
	}
	return copy
}

func (r *Room) IsSolved() bool {
	for _, p := range r.pos {
		if p == nil || p.t != r.t {
			return false
		}
	}
	return true
}

type Pod struct {
	t PodType
}

func (p *Pod) Copy() *Pod {
	return &Pod{t: p.t}
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

type Move struct {
	t          MoveType
	roomIndex  int
	roomPos    int
	hallwayPos int
	cost       int
}

func (m *Move) Cost() int {
	return m.cost
}

const (
	MoveTypeRoomToHallway MoveType = iota
	MoveTypeHallwayToRoom
)

type MoveType int

// Number of rooms in the world
const numRooms = 4

// Number of hallway positions in the world
const numHallways = 7

/*
               +-+-+-+-+-+-+-+-+-+-+-+
Hallway index: |0|1| |2| |3| |4| |5|6|
               +-+-+-+-+-+-+-+-+-+-+-+
                   | | | | | | | |
       Room index: |0| |1| |2| |3|
                   | | | | | | | |
                   +-+ +-+ +-+ +-+
*/

var roomMoveLeft map[int][]int = map[int][]int{
	0: {1, 0},
	1: {2, 1, 0},
	2: {3, 2, 1, 0},
	3: {4, 3, 2, 1, 0},
}

var roomMoveRight map[int][]int = map[int][]int{
	0: {2, 3, 4, 5, 6},
	1: {3, 4, 5, 6},
	2: {4, 5, 6},
	3: {5, 6},
}

var podStepCosts map[PodType]int = map[PodType]int{
	PodTypeA: 1,
	PodTypeB: 10,
	PodTypeC: 100,
	PodTypeD: 1000,
}

var podRoom map[PodType]int = map[PodType]int{
	PodTypeA: 0,
	PodTypeB: 1,
	PodTypeC: 2,
	PodTypeD: 3,
}

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

var hallwayRoomPassThrough map[int]map[int][]int = map[int]map[int][]int{
	0: {
		0: {1},
		1: {1, 2},
		2: {1, 2, 3},
		3: {1, 2, 3, 4},
	},
	1: {
		0: {},
		1: {2},
		2: {2, 3},
		3: {2, 3, 4},
	},
	2: {
		0: {},
		1: {},
		2: {3},
		3: {3, 4},
	},
	3: {
		0: {2},
		1: {},
		2: {},
		3: {4},
	},
	4: {
		0: {3, 2},
		1: {3},
		2: {},
		3: {},
	},
	5: {
		0: {4, 3, 2},
		1: {4, 3},
		2: {4},
		3: {},
	},
	6: {
		0: {5, 4, 3, 2},
		1: {5, 4, 3},
		2: {5, 4},
		3: {5},
	},
}

func NewWorld(roomSize int, podTypes ...PodType) *World {
	if roomSize < 1 {
		panic(fmt.Sprintf("Room size must be more than zero. Got: %d", roomSize))
	}
	if len(podTypes)/roomSize != numRooms {
		panic(fmt.Sprintf("There should be a multiple of 4 pods. Got: %d", len(podTypes)))
	}

	return &World{
		roomSize: roomSize,
		rooms:    createRooms(roomSize, createPods(podTypes)),
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
