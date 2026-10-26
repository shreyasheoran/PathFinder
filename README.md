# PathFinder

Interactive grid pathfinding visualizer: a React frontend and a Go backend that compare two ways of finding a path on a 20×20 grid with obstacles.

- **Backtracking DFS** with branch-and-bound pruning and Manhattan-distance move ordering. It explores every viable route, so it returns a shortest path, but its cost grows quickly as the grid opens up.
- **Dijkstra**: returns a shortest path without exploring every route, so it stays fast on open grids.

The app shows how long each algorithm takes, which makes the difference visible.

## Features

- Click to set start and end points, then toggle obstacles
- Run DFS or Dijkstra and see the path highlighted
- Time taken shown for each run
- Reset the grid

## Tech

React (Create React App), Axios, Go (net/http), CSS

## Run locally

Backend (port 8080):

```bash
cd go-backend
go run main.go
```

Frontend (port 3000):

```bash
cd grid-pathfinder
npm install
npm start
```

Open http://localhost:3000

## API

| Endpoint | Algorithm |
|---|---|
| `/find-path-dfs` | Backtracking DFS |
| `/find-path-dijkstra` | Dijkstra |
