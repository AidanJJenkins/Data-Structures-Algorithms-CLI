package dijkstras

import "testing"

func TestDijktraList(t *testing.T) {

	list := make([][]Edge, 7)
	for i := range list {
		list[i] = make([]Edge, 0)
	}

	list[0] = append(list[0], Edge{to: 1, weight: 3})
	list[0] = append(list[0], Edge{to: 2, weight: 1})

	list[1] = append(list[1], Edge{to: 0, weight: 3})
	list[1] = append(list[1], Edge{to: 2, weight: 4})
	list[1] = append(list[1], Edge{to: 4, weight: 1})

	list[2] = append(list[2], Edge{to: 1, weight: 4})
	list[2] = append(list[2], Edge{to: 3, weight: 7})
	list[2] = append(list[2], Edge{to: 0, weight: 1})

	list[3] = append(list[3], Edge{to: 2, weight: 7})
	list[3] = append(list[3], Edge{to: 4, weight: 5})
	list[3] = append(list[3], Edge{to: 6, weight: 1})

	list[4] = append(list[4], Edge{to: 1, weight: 1})
	list[4] = append(list[4], Edge{to: 3, weight: 5})
	list[4] = append(list[4], Edge{to: 5, weight: 2})

	list[5] = append(list[5], Edge{to: 6, weight: 1})
	list[5] = append(list[5], Edge{to: 4, weight: 2})
	list[5] = append(list[5], Edge{to: 2, weight: 18})

	list[6] = append(list[6], Edge{to: 3, weight: 1})
	list[6] = append(list[6], Edge{to: 5, weight: 1})

	result := DijkstraList(0, 6, list)
	expected := []int{0, 1, 4, 5, 6}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("got %v, want %v", result[i], expected[i])
		}
	}
}
