# Architecture

## Overview

The project is structured following Hexagonal Architecture, also known as Ports and Adapters. The main goal of this pattern is to isolate the core business logic from everything external: the database, the HTTP layer, and the AI provider. This means the domain and services have no direct dependency on PostgreSQL, Groq, or any HTTP framework. They only interact with interfaces defined in the ports layer.

This separation makes the codebase significantly easier to test, since every external dependency can be replaced by a mock without touching the business logic. It also makes it easy to swap out any infrastructure piece.

## How the layers are organized

The `internal/core` directory is the heart of the application. Inside it, the `domain` package defines the central business entities: `Car` and `User`. These structs represent the data the application reasons about and contain no framework-specific code. The `ports` package defines Go interfaces that describe what the application needs from the outside world — for example, what operations a car repository must support, or what an AI provider must be able to do. The `services` package contains the actual business logic: authentication, car searches, favorites management, and AI recommendation generation. Services depend only on the port interfaces, never on concrete implementations.

The `internal/adapters` directory contains all the concrete implementations of those interfaces. The `handlers` package translates HTTP requests into service calls and writes the responses back. The `repositories` package implements the database operations using PostgreSQL. The `groq` package implements the AI provider interface by calling the Groq API. The `simulator` package contains a background goroutine that periodically updates vehicle prices, simulating real-time market fluctuations.

The `public` directory holds the Vue 3 frontend, which the Go server serves as a static file. This keeps the whole project as a single deployable unit — one binary serves both the API and the UI.

The `test/mocks` directory contains mock implementations of every port interface. These are used in unit tests to simulate database responses and AI outputs without any real network or database connections, making the test suite fast and deterministic.

## How a request flows through the system

To make the architecture concrete, here is how a request for an AI recommendation travels through the application. The user fills in a text prompt in the browser and clicks the button. The Vue frontend sends a POST request to `/api/recommend` with a JSON body containing the preferences and a JWT token in the Authorization header.

The Go HTTP server receives the request and routes it to the `RecommendationHandler` in the adapters layer. The handler validates the JWT, extracts the preferences from the body, and calls the `RecommendationService` in the core layer. The service then calls the `CarRepository` interface to fetch a relevant set of vehicles from the database. It builds a prompt combining the user's natural language input with the real car data, and passes it to the `AIProvider` interface. The `GroqClient` adapter takes over, sends the prompt to the Groq API, and returns the recommendation text. The service returns this text to the handler, which writes the JSON response back to the browser.

The key detail is that the `RecommendationService` never knows whether it is talking to a real PostgreSQL database or a mock, and never knows whether the AI response comes from Groq or a test double. It only talks to interfaces.

## Concurrency

The price simulator is launched as a goroutine at startup with `go sim.Start()` in `main.go`. It runs in the background for the lifetime of the application, updating vehicle prices periodically without blocking the main HTTP server thread.