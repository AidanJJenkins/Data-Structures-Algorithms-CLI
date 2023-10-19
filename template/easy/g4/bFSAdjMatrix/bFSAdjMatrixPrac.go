package bfsadjmatrix

type WeightedAdjacencyMatrix [][]int

func BFS(graph WeightedAdjacencyMatrix, source int, needle int) []int {

}

func rev(out []int) {
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
}
