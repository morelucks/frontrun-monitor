# Frontrun Monitor

A high-performance frontrunning detection and monitoring system with a Go backend and modern frontend.

## Project Structure

```
frontrun-monitor/
├── backend/          # Go backend service
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   └── ...
├── frontend/         # Frontend application
│   ├── src/
│   ├── public/
│   ├── package.json
│   └── ...
└── README.md
```

## Technology Stack

### Backend
- **Language**: Go (Golang)
- **Purpose**: Core monitoring engine and API service
- **Features**:
  - High-performance transaction analysis
  - Real-time frontrunning detection
  - RESTful API endpoints
  - WebSocket support for live updates

### Frontend
- **Modern UI framework** for visualization and monitoring dashboard
- **Real-time data updates** via WebSocket connections
- **Responsive design** for multiple devices

## Getting Started

### Prerequisites
- Go 1.19 or higher (for backend)
- Node.js 16+ (for frontend)
- Your preferred code editor

### Backend Setup
```bash
cd backend
go mod download
go run main.go
```

### Frontend Setup
```bash
cd frontend
npm install
npm start
```

## Features

- **Real-time Detection**: Monitor mempool for suspicious transaction patterns
- **High Performance**: Built with Go for optimal speed and efficiency
- **Live Dashboard**: Interactive frontend for viewing detected threats
- **WebSocket Integration**: Real-time updates between backend and frontend

## API Documentation

The Go backend provides a comprehensive REST API for monitoring and analysis. See the backend documentation for detailed endpoint specifications.

## Contributing

Contributions are welcome! Please feel free to submit pull requests.

## License

MIT License
