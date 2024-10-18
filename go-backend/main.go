package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"sort"

	"github.com/rs/cors"
)

// Coordinates represents a single tile in the grid.
type Coordinates struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

// PathRequest defines the structure for the request body.
type PathRequest struct {
	Start     Coordinates   `json:"start"`
	End       Coordinates   `json:"end"`
	Obstacles []Coordinates `json:"obstacles"`
}

// PathResponse defines the structure for the response body.
type PathResponse struct {
	Path []Coordinates `json:"path"`
}

// Directions (up, down, left, right).
var directions = []Coordinates{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

// Helper function to check if a coordinate is an obstacle.
func isObstacle(coord Coordinates, obstacles []Coordinates) bool {
	for _, obstacle := range obstacles {
		if coord.Row == obstacle.Row && coord.Col == obstacle.Col {
			return true
		}
	}
	return false
}

// Manhattan distance heuristic to prioritize neighbors
func manhattanDistance(a, b Coordinates) int {
	return abs(a.Row-b.Row) + abs(a.Col-b.Col)
}

// Optimized DFS function with early exit, path pruning, and heuristic
func optimizedDFS(start, end Coordinates, obstacles []Coordinates, gridSize int) []Coordinates {
	visited := make([][]bool, gridSize)
	for i := range visited {
		visited[i] = make([]bool, gridSize)
	}

	var bestPath []Coordinates
	var currentPath []Coordinates
	bestLength := math.MaxInt64 // Keep track of the best path length found

	var dfsRecursive func(current Coordinates, pathLength int)
	dfsRecursive = func(current Coordinates, pathLength int) {
		log.Printf("Visiting node: %+v with path length: %d", current, pathLength)

		// Base case: reached the end point
		if current.Row == end.Row && current.Col == end.Col {
			log.Printf("Reached end: %+v", current)
			currentPath = append(currentPath, current)
			if pathLength < bestLength {
				// Update best path if current path is shorter
				bestPath = make([]Coordinates, len(currentPath))
				copy(bestPath, currentPath)
				bestLength = pathLength
				log.Printf("New best path: %+v with length: %d", bestPath, bestLength)
			}
			currentPath = currentPath[:len(currentPath)-1] // Backtrack
			return
		}

		// Prune paths longer than the best path found so far
		if pathLength >= bestLength {
			log.Printf("Pruning path at node: %+v with length: %d", current, pathLength)
			return
		}

		// Mark the current cell as visited
		visited[current.Row][current.Col] = true
		currentPath = append(currentPath, current)

		// Sort directions by Manhattan distance to prioritize better directions
		directionsWithPriority := directions
		sort.Slice(directionsWithPriority, func(i, j int) bool {
			nextA := Coordinates{Row: current.Row + directions[i].Row, Col: current.Col + directions[i].Col}
			nextB := Coordinates{Row: current.Row + directions[j].Row, Col: current.Col + directions[j].Col}
			return manhattanDistance(nextA, end) < manhattanDistance(nextB, end)
		})

		// Explore neighbors
		for _, direction := range directionsWithPriority {
			neighbor := Coordinates{
				Row: current.Row + direction.Row,
				Col: current.Col + direction.Col,
			}

			// Ensure neighbor is within bounds, not visited, and not an obstacle
			if neighbor.Row >= 0 && neighbor.Row < gridSize &&
				neighbor.Col >= 0 && neighbor.Col < gridSize &&
				!visited[neighbor.Row][neighbor.Col] && !isObstacle(neighbor, obstacles) {

				log.Printf("Exploring neighbor: %+v", neighbor)
				dfsRecursive(neighbor, pathLength+1)
			}
		}

		// Unmark the current cell and backtrack
		visited[current.Row][current.Col] = false
		currentPath = currentPath[:len(currentPath)-1]
	}

	// Start the DFS
	log.Printf("Starting DFS from: %+v to %+v", start, end)
	dfsRecursive(start, 0)

	return bestPath
}

// Dijkstra's algorithm for shortest path considering obstacles.
func dijkstra(start, end Coordinates, obstacles []Coordinates) []Coordinates {
	gridSize := 20
	distances := make([][]float64, gridSize)
	for i := range distances {
		distances[i] = make([]float64, gridSize)
		for j := range distances[i] {
			distances[i][j] = math.Inf(1)
		}
	}
	distances[start.Row][start.Col] = 0

	parent := make(map[Coordinates]*Coordinates)
	queue := []Coordinates{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Row == end.Row && current.Col == end.Col {
			break
		}

		for _, direction := range directions {
			neighbor := Coordinates{
				Row: current.Row + direction.Row,
				Col: current.Col + direction.Col,
			}

			if neighbor.Row >= 0 && neighbor.Row < gridSize &&
				neighbor.Col >= 0 && neighbor.Col < gridSize &&
				!isObstacle(neighbor, obstacles) {
				newDist := distances[current.Row][current.Col] + 1

				if newDist < distances[neighbor.Row][neighbor.Col] {
					distances[neighbor.Row][neighbor.Col] = newDist
					parent[neighbor] = &current
					queue = append(queue, neighbor)
				}
			}
		}
	}

	var path []Coordinates
	for current := end; current != start; current = *parent[current] {
		path = append([]Coordinates{current}, path...)
	}
	path = append([]Coordinates{start}, path...)

	return path
}

// Handler for DFS pathfinding.
func findPathDFS(w http.ResponseWriter, r *http.Request) {
	var request PathRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Received DFS pathfinding request: %+v", request)

	// Use optimized DFS to find the path.
	path := optimizedDFS(request.Start, request.End, request.Obstacles, 20)
	response := PathResponse{Path: path}

	log.Printf("DFS path response: %+v", response)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handler for Dijkstra's pathfinding.
func findPathDijkstra(w http.ResponseWriter, r *http.Request) {
	var request PathRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Received Dijkstra pathfinding request: %+v", request)

	// Use Dijkstra to find the path.
	path := dijkstra(request.Start, request.End, request.Obstacles)
	response := PathResponse{Path: path}

	log.Printf("Dijkstra path response: %+v", response)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	c := cors.Default()

	mux := http.NewServeMux()
	mux.HandleFunc("/find-path-dfs", findPathDFS)         // Endpoint for DFS
	mux.HandleFunc("/find-path-dijkstra", findPathDijkstra) // Endpoint for Dijkstra

	handler := c.Handler(mux)

	log.Println("Server is starting on port 8080...")

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// Utility function for absolute values
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
