package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"
	"unicode"
)

type Point struct {
	x, y int
}

type Node struct {
	location Point
	value    int
	through  *Node
}

type Edge struct {
	node   *Node
	weight int
}

// --- Graph ---

type Graph struct {
	Nodes []*Node
	Edges map[Point][]*Edge
	mutex sync.RWMutex
}

func newGraph() *Graph {
	return &Graph{Edges: make(map[Point][]*Edge)}
}

func (g *Graph) GetNode(x, y int) (node *Node) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	for _, n := range g.Nodes {
		if reflect.DeepEqual(n.location, Point{x, y}) {
			node = n
		}
	}
	return
}

func (g *Graph) AddNode(n *Node) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(n1, n2 *Node, weight int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Edges[n1.location] = append(g.Edges[n1.location], &Edge{n2, weight})
	// g.Edges[n2.location] = append(g.Edges[n2.location], &Edge{n1, weight})
}

// --- HEAP ---

type Heap struct {
	elements []*Node
	mutex    sync.RWMutex
}

func (h *Heap) Size() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return len(h.elements)
}

// push an element to the heap, re-arrange the heap
func (h *Heap) Push(element *Node) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.elements = append(h.elements, element)
	i := len(h.elements) - 1
	for ; h.elements[i].value < h.elements[parent(i)].value; i = parent(i) {
		h.swap(i, parent(i))
	}
}

// pop the top of the heap, which is the min value
func (h *Heap) Pop() (i *Node) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	i = h.elements[0]
	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]
	h.rearrange(0)
	return
}

// rearrange the heap
func (h *Heap) rearrange(i int) {
	smallest := i
	left, right, size := leftChild(i), rightChild(i), len(h.elements)
	if left < size && h.elements[left].value < h.elements[smallest].value {
		smallest = left
	}
	if right < size && h.elements[right].value < h.elements[smallest].value {
		smallest = right
	}
	if smallest != i {
		h.swap(i, smallest)
		h.rearrange(smallest)
	}
}

func (h *Heap) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

//  --- Utils ---

func getSpecialCoordinate(data []string, c rune) Point {
	for i, row := range data {
		for j, v := range row {
			if v == c {
				return Point{i, j}
			}
		}
	}
	return Point{}
}

func intAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func replaceSpecialChars(a byte) byte {
	switch a {
	case 'S':
		a = 'a'
	case 'E':
		a = 'z'
	}
	return a
}

func calcHeightDiff(a, b byte) int {
	// replace special characters
	a, b = replaceSpecialChars(a), replaceSpecialChars(b)
	return intAbs(int(a) - int(b))
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

//  --- Main ---

func parseFile(fp string) []string {
	file, err := os.Open(fp)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	data := strings.TrimRightFunc(string(bytes), unicode.IsSpace)
	return strings.Split(data, "\n")
}

func buildGraph(data []string) *Graph {
	defer timeTrack(time.Now(), "Building graph")

	graph := newGraph()
	height := len(data)
	width := len(data[0])

	for i, row := range data {
		for j := range row {
			n := &Node{Point{i, j}, math.MaxInt, nil}
			graph.AddNode(n)
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			n := graph.GetNode(i, j)
			h := data[i][j]

			if i > 0 {
				graph.AddEdge(n, graph.GetNode(i-1, j), calcHeightDiff(data[i-1][j], h))
			}
			// bottom
			if i < height-1 {
				graph.AddEdge(n, graph.GetNode(i+1, j), calcHeightDiff(data[i+1][j], h))
			}
			// left
			if j > 0 {
				graph.AddEdge(n, graph.GetNode(i, j-1), calcHeightDiff(data[i][j-1], h))
			}
			// right
			if j < width-1 {
				graph.AddEdge(n, graph.GetNode(i, j+1), calcHeightDiff(data[i][j+1], h))
			}
		}
	}

	return graph
}

func dijkstra(graph *Graph, start Point, end Point) int {
	defer timeTrack(time.Now(), "Dijkstra's algorithm")

	visited := make(map[Point]bool)
	heap := &Heap{}

	startNode := graph.GetNode(start.x, start.y)
	startNode.value = 0
	heap.Push(startNode)

	for heap.Size() > 0 {
		current := heap.Pop()
		visited[current.location] = true
		edges := graph.Edges[current.location]
		for _, edge := range edges {
			if visited[edge.node.location] || edge.weight > 1 {
				continue
			}
			heap.Push(edge.node)
			if current.value+edge.weight < edge.node.value {
				edge.node.value = current.value + 1
				edge.node.through = current
				if reflect.DeepEqual(edge.node.location, end) {
					return edge.node.value
				}
			}
		}
	}
	return 0
}

func main() {
	input_fp := "day_12/test_input.txt"
	fileData := parseFile(input_fp)

	start, end := getSpecialCoordinate(fileData, 'S'), getSpecialCoordinate(fileData, 'E')
	fmt.Printf("Start: %v, End: %v\n", start, end)
	fmt.Println("Building graph...")
	graph := buildGraph(fileData)

	fmt.Println("Starting Dijkstra's algorithm...")
	shortestDistance := dijkstra(graph, start, end)
	fmt.Printf("\nShortest distance from %v to %v: %d steps\n", start, end, shortestDistance)

	// draw taken path
	var shortestPath []Point
	for n := graph.GetNode(end.x, end.y); n.through != nil; n = n.through {
		shortestPath = append(shortestPath, n.location)
	}
	fmt.Println()
	for i, row := range fileData {
		for j, cell := range row {
			p := Point{i, j}
			for _, v := range shortestPath {
				if reflect.DeepEqual(p, v) {
					cell = '.'
				}
			}
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}
