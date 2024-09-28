# Authentication System in Go

This repository contains a robust **Authentication System** built using the Go programming language. The system includes essential features such as **user registration**, **login**, **JWT-based authentication**, and **role-based access control (RBAC)**, making it suitable for secure API services.

## Key Features

- **User Registration**: Secure user sign-up with email verification.
- **Login**: Authenticate users with secure password hashing using bcrypt.
- **JWT Authentication**: Issue and validate JSON Web Tokens (JWT) to ensure secure API access.
- **Role-Based Access Control (RBAC)**: Assign and enforce roles and permissions for different users.
- **Session Management**: Token refresh and session expiration handling to keep authentication smooth and secure.
- **Password Recovery**: Secure password reset functionality with email verification.
- **Middleware for Protected Routes**: Implement middleware to restrict access to routes that require user authentication.

## Tech Stack

- **Go**: Efficient, scalable back-end programming language.
- **JWT**: Secure token-based authentication for session management.
- **bcrypt**: Password hashing algorithm to ensure security.
- **Gin**: For routing HTTP requests in the application.
- **MongoDB**: Flexible support for popular databases (configurable).
- **Docker**: Containerized setup for easy and consistent deployment.

## Installation

### Prerequisites

- Go 1.16 or higher
- Docker (for containerized setup)
- PostgreSQL, MySQL, or MongoDB

### Steps to Run Locally

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Manikandan-Parasuraman/authentication-system-go.git
