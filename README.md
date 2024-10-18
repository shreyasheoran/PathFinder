Pathfinder App

A 20x20 Grid Pathfinding Application Using DFS and Dijkstra Algorithms
Welcome to Pathfinder, a web-based application that allows users to select start and end points on a 20x20 grid and calculate the shortest path between them using either the Depth-First Search (DFS) or Dijkstra algorithms. Users can also add obstacles to the grid to see how the algorithms navigate around them.

Features:
Interactive Grid Layout: A 20x20 clickable grid where users can select the start point, end point, and toggle obstacles.
DFS and Dijkstra Algorithms: Choose between two pathfinding algorithms, Depth-First Search (DFS) and Dijkstra, to calculate the shortest path.
Path Highlighting: The path found between the start and end points is visually highlighted.
Processing Time Display: The time taken for each algorithm to calculate the path is displayed after completion.
Responsive UI: The app adjusts nicely to different screen sizes, providing a professional and modern look.

Getting Started
Prerequisites
Before you begin, ensure you have the following installed:

Node.js (v14 or higher)
npm (comes with Node.js)
Go (if you want to run the backend locally)
Installation
Clone the repository to your local machine:

bash
git clone https://github.com/shreyasheoran/pathfinder-app.git
cd pathfinder-app


Install dependencies for the React frontend:

bash
npm install
Start the React frontend:

bash
npm start


Set up the Go backend:

The Go backend is responsible for calculating the path using DFS or Dijkstra algorithms. You can run it locally.

Run the backend locally:

go run main.go


Access the application:

Open your browser and go to http://localhost:3000 to view the app.
Usage
Select start and end points:

Click any cell on the grid to set the start point, and click another cell to set the end point.
Add obstacles:

After setting the start and end points, click other grid cells to toggle obstacles. These cells will be marked as obstacles for the algorithms to avoid.
Choose an algorithm:

Click Find Path Using DFS to use Depth-First Search or Find Path Using Dijkstra to calculate the shortest path using Dijkstra's algorithm.
See the result:

The path between the start and end points will be highlighted, and the time taken to calculate the path will be displayed below the buttons.
Reset:

Click the Reset Grid button to clear the grid and start over.


Technologies Used
React.js for the frontend
Axios for making HTTP requests to the backend
Go (Golang) for backend algorithms (DFS and Dijkstra)
CSS for styling and responsive layout
