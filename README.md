
# 🌐 Patricia's Go HTTP Server

A simple, lightweight HTTP web server built with Go's standard library. This project demonstrates core Go concepts including HTTP routing, environment-based configuration, and graceful error handling.

---

## 📋 Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Server](#running-the-server)
- [Configuration](#configuration)
- [Project Structure](#project-structure)
- [How It Works](#how-it-works)
- [Future Improvements](#future-improvements)

---

## Overview

This is a beginner-friendly Go HTTP server that:

- Listens for incoming HTTP requests on a configurable port
- Responds with a greeting message on the root route `/`
- Supports environment-based port configuration for easy deployment to cloud platforms like Railway, Render, or Heroku
- Handles server errors gracefully using Go's `log` package

---

## Prerequisites

Before running this project, make sure you have the following installed:

- [Go](https://go.dev/dl/) (version 1.18 or higher recommended)

Verify your installation by running:

```bash
go version
```

You should see output like `go version go1.21.0 ...`

---

## Installation

1. Clone or download this repository to your local machine:

```bash
git clone (https://github.com/Patricia05-alt/golang_text-_example)
cd golang_text-_example
```

2. No additional dependencies are needed — this project uses only Go's standard library!

---

## Running the Server


Start the server directly using `go run` (default port `:8080`):

```powershell
cd 'c:\Users\USER\OneDrive\Desktop\Hello world\golang_text-_example'
go run .
```

Start the server on a different port by setting the `PORT` environment variable (examples use PowerShell):

```powershell
cd 'c:\Users\USER\OneDrive\Desktop\Hello world\golang_text-_example'
$env:PORT = '8081'   # will be normalized to ":8081"
go run .
```

**Run in background (PowerShell / Windows)**

Launch the server in a detached process and capture logs to files:

```powershell
cd 'c:\Users\USER\OneDrive\Desktop\Hello world\golang_text-_example'
Start-Process -FilePath cmd -ArgumentList '/c','set PORT=8081&&go run .' -WorkingDirectory (Get-Location) -NoNewWindow -RedirectStandardOutput "$PWD\server.out" -RedirectStandardError "$PWD\server.err"
```

Check `server.out` for the startup message and `server.err` for errors.

**Build (produce executable)**

```powershell
cd 'c:\Users\USER\OneDrive\Desktop\Hello world\golang_text-_example'
go build -o hello-server.exe
.
```

**Verify the server**

Use PowerShell to request the root path and see the response:

```powershell
Invoke-RestMethod -Uri http://localhost:8081/
# or use curl
curl http://localhost:8081/
```

Expected response:

```
Hello, World...this Is Patricia first go project! This is a Go HTTP Server.
```

**Stop the server**

Find and stop the server process (example shows killing by PID found via `netstat`):

```powershell
netstat -ano | findstr :8081
Stop-Process -Id <PID> -Force
```

Alternatively, if you started the server in a visible terminal, press `Ctrl+C` there.

## Configuration

- The program reads the `PORT` environment variable and will automatically add a leading `:` if you provide only the port number (for example `8081`).
- If port `8080` (or your chosen port) is already in use, either stop the process using it or pick another port.

This makes it compatible with most cloud hosting platforms that inject a `PORT` variable at runtime.

---

## Project Structure

```
golang_text-_example/
│
└── main.go       # Entry point — contains the handler and server setup
```

---

## How It Works

### 1. The Handler Function

```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World...this Is Patricia first go project!")
}
```

The `handler` function is called every time a request comes in. It receives two arguments:
- `w http.ResponseWriter` — used to write the response back to the client
- `r *http.Request` — contains details about the incoming request (method, headers, body, etc.)

### 2. Route Registration

```go
http.HandleFunc("/", handler)
```

This tells Go: *"When someone visits the root path `/`, call the `handler` function."*

### 3. Port Configuration

```go
portEnv := os.Getenv("PORT")
```

The server reads the `PORT` environment variable, making it cloud-deployment ready out of the box.

### 4. Starting the Server

```go
if err := http.ListenAndServe(port, nil); err != nil {
    log.Fatalf("Error starting server: %v", err)
}
```

`http.ListenAndServe` starts the server and blocks (runs forever) until it's stopped. If something goes wrong (e.g. the port is already in use), `log.Fatalf` prints the error and exits the program cleanly.

---

## Future Improvements

Some ideas for extending this project:

- Add more routes (e.g. `/about`, `/api/hello`)
- Return JSON responses for API-style endpoints
- Add middleware for logging incoming requests
- Serve static HTML files
- Add unit tests using Go's `testing` package
- Use a router library like [gorilla/mux](https://github.com/gorilla/mux) or [chi](https://github.com/go-chi/chi) for more advanced routing

---

## Author

**Patricia** — first Go project 🎉

---

## License

This project is open source and available under the [MIT License](LICENSE).
