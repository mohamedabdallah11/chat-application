
# Chat Application in Go

This project is a simple **chat application** implemented in **Go**. It allows multiple clients to connect to a server, send and receive messages in real-time. The server broadcasts messages to all connected clients. The project demonstrates the use of **RPC (Remote Procedure Call)** in Go for communication between the client and server.

---

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [Directory Structure](#directory-structure)
- [Code Implementation](#code-implementation)
- [Contributing](#contributing)

---

## Features

- **Client-Server Communication**: The server handles multiple clients using **RPC** and broadcasts messages to all connected clients.
- **Dynamic Port Allocation**: The client dynamically assigns a port for communication.
- **Real-time Messaging**: Messages sent by one client are broadcasted to all other clients in real-time.
- **Message Reception**: Clients can receive messages from other clients.

---

## Prerequisites

To run this project, you'll need to have the following installed:

- **Go (Golang)**: Version 1.16 or higher
  - Install Go: [Download Go](https://golang.org/dl/)
- **Git**: For cloning the repository
  - Install Git: [Download Git](https://git-scm.com/downloads)

---
## Installation and Running the Application

1. **Clone the repository**:

   Open your terminal and run the following command to clone the repository:

   ```bash
   git clone https://github.com/mohamedabdallah11/chat-application.git
   cd chat-application
Install dependencies (if any) using Go modules:

Run the following command to install the necessary dependencies:

bash
Copy code
go mod tidy
Run the Server:

In the terminal, navigate to the server folder and run the following command to start the server:

bash
Copy code
cd server
go run server.go
The server will start and listen on port 42586 by default.

Run the Client:

Open a new terminal window (keeping the server running in the first terminal), navigate to the client folder, and run the following command to start the client:

bash
Copy code
cd client
go run client.go
The client will prompt you for your name. After you enter your name, it will connect to the server and dynamically allocate a port for communication. You can now send and receive messages.
