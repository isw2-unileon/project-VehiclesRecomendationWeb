# Design Decisions

## Go as the backend language

Go was mandatory for this project. 

## Hexagonal Architecture

We chose to structure the backend around Hexagonal Architecture because we wanted the business logic to be completely independent of the infrastructure. In practice this means the core services do not import any database driver, HTTP framework, or AI SDK. They only depend on Go interfaces defined in the ports layer. This decision paid off during testing: we could write unit tests for every service and handler without spinning up a real database or making real API calls, just by swapping the real implementations for the mocks in `test/mocks`.

## Vue 3 via CDN for the frontend

The Go server already serves static files from the `public` directory, so the simplest approach was to put the entire frontend in a single `index.html` file loaded via CDN. Using Vue 3 this way gives us reactive data binding, computed properties, and a clean component-style structure without requiring any build tooling — no npm, no Vite, no Node.js pipeline. The whole project remains a single deployable unit: one Go binary that serves both the API and the UI. TailwindCSS is also loaded via CDN for the same reason.

## PostgreSQL as the database

SQL was a hard requirement for this project. Among the available SQL options, PostgreSQL was the natural choice because it is well-supported on Render with a free managed tier, has a solid Go driver (`lib/pq`), and supports full-text search natively, which we used for the text search endpoint. It also provides ACID compliance, which matters for the user account and favorites system where consistency is important.

## Groq with Llama 3.1 8B Instant for AI recommendations

We chose Groq over other AI providers primarily because of its inference speed. Groq's LPU architecture returns responses in under two seconds on the free tier, which makes the recommendation feature feel responsive rather than like a slow background job. Llama 3.1 8B Instant produces high-quality natural language recommendations while being fast enough for real-time use.

The recommendation service does not just forward the user's prompt to the AI. It first queries the car database to fetch relevant vehicles and injects that data into the prompt. This way the AI can only recommend cars that actually exist in our catalogue, which avoids the hallucination problem that would occur if we simply asked the model to recommend cars from its training data.

The AI provider is also abstracted behind an interface (`ports/ai_provider.go`), so swapping it for a different model or provider in the future would only require writing a new adapter.

## JWT for authentication and Bcrypt for passwords

JWT was chosen for session management because it is stateless — the server does not need to store any session data. After a successful login the client receives a signed token that it includes in subsequent requests. This keeps the Go server simple and makes it straightforward to scale horizontally if needed.

For password storage, Bcrypt is the standard choice in the Go ecosystem. It automatically handles salting, which means two users with the same password will always produce different hashes in the database. The `Password` field in the `User` struct is tagged with `json:"-"` so the hash can never accidentally leak into an API response.

## Render for deployment

Render was chosen because it offers a free tier that supports Go web services and managed PostgreSQL, which was all we needed for this project. Auto-deploy on push to `main` integrates naturally with the Trunk Based Development workflow required by the project specification. The deployment configuration is stored as code in `render.yaml`, which means anyone who forks the repository can reproduce the deployment environment. All secrets (`I/E-DATABASE_URL`, `GROQ_API_KEY`, `JWT_SECRET`) are stored as environment variables in the Render dashboard and are never committed to the repository.