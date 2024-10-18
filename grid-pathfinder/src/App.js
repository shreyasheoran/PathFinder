import React, { useState } from "react";
import axios from "axios";
import "./App.css";

const App = () => {
  // State for managing the grid (20x20 cells)
  const [grid, setGrid] = useState(
    Array(20)
      .fill(null)
      .map(() => Array(20).fill(false))
  );

  // State for managing the start, end points, and the calculated path
  const [start, setStart] = useState(null);
  const [end, setEnd] = useState(null);
  const [path, setPath] = useState([]);

  // State to track loading status, processing text, and time taken
  const [loading, setLoading] = useState(false);
  const [processingText, setProcessingText] = useState(""); // Text while processing
  const [timeTakenMessage, setTimeTakenMessage] = useState(""); // Message for time taken

  // Handles cell clicks to set start, end points, or toggle obstacles
  const handleCellClick = (row, col) => {
    if (start && start.row === row && start.col === col) return; // Ignore clicks on start
    if (end && end.row === row && end.col === col) return; // Ignore clicks on end

    // Set start or end, or toggle obstacles
    if (!start) {
      setStart({ row, col });
    } else if (!end) {
      setEnd({ row, col });
    } else {
      const newGrid = [...grid];
      newGrid[row][col] = !newGrid[row][col]; // Toggle obstacle
      setGrid(newGrid);
    }
  };

  // Find path using the selected algorithm (DFS or Dijkstra)
  const findPath = async (algorithm) => {
    // Ensure start and end points are selected
    if (!start || !end) {
      alert("Please select both start and end points.");
      return;
    }

    setLoading(true); // Disable buttons while processing
    setProcessingText(`Processing path using ${algorithm.toUpperCase()}...`);
    setTimeTakenMessage(""); // Clear previous time message when starting a new search

    try {
      const startTime = Date.now(); // Track start time

      // Send request to the backend API
      const response = await axios.post(
        `http://localhost:8080/find-path-${algorithm}`,
        {
          start,
          end,
          obstacles: grid
            .flat()
            .map((cell, index) =>
              cell
                ? {
                    row: Math.floor(index / 20),
                    col: index % 20,
                  }
                : null
            )
            .filter(Boolean), // Filter out null values
        }
      );

      // Update the path with backend response
      setPath(response.data.path);

      // Calculate and display time taken for pathfinding
      const endTime = Date.now();
      const timeTaken = ((endTime - startTime) / 1000).toFixed(2); // Time in seconds
      setProcessingText(""); // Clear processing message
      setTimeTakenMessage(
        `Path found using ${algorithm.toUpperCase()} in ${timeTaken} seconds.`
      );
    } catch (error) {
      console.error("Network error:", error);
      alert(
        "Failed to connect to the server. Please ensure the backend is running."
      );
    } finally {
      setLoading(false); // Re-enable buttons after processing
    }
  };

  // Reset the grid, start, end, and path
  const resetGrid = () => {
    setGrid(
      Array(20)
        .fill(null)
        .map(() => Array(20).fill(false))
    );
    setStart(null);
    setEnd(null);
    setPath([]);
    setProcessingText(""); // Clear processing text
    setTimeTakenMessage(""); // Clear time taken message
  };

  return (
    <div className="App">
      {/* Control panel for pathfinding and grid reset */}
      <div className="controls">
        <h1>Pathfinder</h1>
        <p>
          Select start, end points, and obstacles to calculate the shortest
          path!
        </p>

        {/* Buttons for DFS and Dijkstra algorithms */}
        <button onClick={() => findPath("dfs")} disabled={loading}>
          Find Path Using DFS
        </button>
        <button onClick={() => findPath("dijkstra")} disabled={loading}>
          Find Path Using Dijkstra
        </button>

        {/* Reset button */}
        <button onClick={resetGrid} disabled={loading}>
          Reset Grid
        </button>

        {/* Show processing text while loading */}
        {loading && <div className="processing">{processingText}</div>}

        {/* Show time taken message after path is found */}
        {timeTakenMessage && (
          <div className="processing">{timeTakenMessage}</div>
        )}
      </div>

      {/* Grid rendering */}
      <div className="grid-container">
        {grid.map((row, rowIndex) =>
          row.map((cell, colIndex) => {
            const isStart =
              start && start.row === rowIndex && start.col === colIndex;
            const isEnd = end && end.row === rowIndex && end.col === colIndex;
            const isPath = path.some(
              (p) => p.row === rowIndex && p.col === colIndex
            );
            const isObstacle = cell;

            return (
              <div
                key={`${rowIndex}-${colIndex}`}
                className={`grid-cell ${isStart ? "start" : ""} ${
                  isEnd ? "end" : ""
                } ${isPath ? "path-highlight" : ""} ${
                  isObstacle ? "obstacle" : ""
                }`}
                onClick={() => handleCellClick(rowIndex, colIndex)}
              >
                {/* Small text for start and end */}
                {isStart ? (
                  <small>Start</small>
                ) : isEnd ? (
                  <small>End</small>
                ) : (
                  ""
                )}
              </div>
            );
          })
        )}
      </div>
    </div>
  );
};

export default App;
