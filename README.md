# Personal Finance Tracker API

## Description
A simple modular monolith backend API built with go, with proper authentication and authorization, to help the user keep track of all income and expenses, and reports describing the state of the user's finances, monthly or all time.

## Purpose
This project was built with the goal of exploring API creation in go, by building a well-designed modular monolithic architecture, with a clear and well-organized project structure. This project also includes a terminal interface that displays the interaction of a client application communicating with the API, with terminal commands.

## Current Limitations
Tests are not currently included. Testing was intentionally left as a follow-up task due to the project's original time constraints. Instead, this project currently focuses on presenting a clear architecture and structure of a modular monolith API.

## Installation & How to Run
1. Clone the repo:
```
git clone 
```

2. Run the server:
First, move to the location of `main.go`:
```
cd api/cmd/api
```

Now, run:
```
go run .
```

3. Run with docker:
Build the image:
```
docker build -t docker-pet .
```

Run:
```
docker run -p 8000:8000 docker-pet
```

## Usage
1. Using the TUI
To run the TUI, in a new terminal, go to the directory:
```
cd tui
```

Now, run:
```
go run main.go
```

## Future Work
- Testing: The next step is to add tests around domain modules and API boundaries.
- Spending alerts: Notify the user when total spending reaches a defined threshold.
- Cost of life comparison: Compare the monthly spending to the average cost of living, based on the address of the user.
- Better TUI interface: Use packages like BubbleTea to improve the interface of the client application.
- AI summaries: Use AI to produce better reports, and give improved spending recommendations.

## License
MIT License
