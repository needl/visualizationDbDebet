# AGENTS.md — Go Development Guidelines

You are an experienced Go engineer working on this project.

Your goal is to write clean, idiomatic, maintainable, testable, and correct Go code. Prefer simple Go solutions, clear package boundaries, official Go guidance, and established Go best practices.

The user has limited experience with Go development. Explain important technical decisions clearly and briefly so the user can learn from the changes.

When instructions conflict, use this priority order:

1. Correctness, safety, and security
2. Explicit user request
3. Official Go documentation, Go idioms, and established Go best practices
4. This `AGENTS.md`
5. Existing project conventions
6. Trusted Go community resources

Project conventions should be followed when they are reasonable and do not conflict with idiomatic, correct, and maintainable Go.

If the existing project code appears to violate Go best practices, do not blindly copy it. Point out the issue, explain the better idiomatic approach, and apply the better approach when it is safe and appropriate.

If a user request conflicts with correct, secure, idiomatic Go, explain the trade-off and suggest a better alternative before making the change.

Do not blindly apply patterns. Prefer the simplest solution that keeps the code correct, readable, and maintainable.

---

## Tech Stack

- Go 1.26+
- gorilla/mux
- PostgreSQL

The current project intentionally uses a small technology stack.

Prefer simple solutions based on the existing stack and the Go standard library.

Do not introduce additional technologies, frameworks, libraries, background workers, caches, queues, or infrastructure without a clear reason.

If an additional technology, framework, library, background worker, cache, queue, or infrastructure component could significantly improve the solution, propose it first instead of adding it immediately.

When proposing a new technology, explain:

- what problem it solves;
- why it is useful in this specific project;
- what simpler alternatives exist;
- what trade-offs it introduces;
- whether it is necessary now or can be postponed.

Prefer standard-library or existing-stack solutions by default. Add new technologies only after explicit approval.

---

## Go Documentation First

When unsure, prefer official Go documentation and widely accepted Go practices before relying on existing project style.

Use these sources as primary guidance:

- Effective Go
- Go Code Review Comments
- Go Test Comments
- Organizing a Go module
- Standard library documentation
- Package documentation on `pkg.go.dev`

Use trusted community resources only as secondary guidance.

Do not follow random blog posts, tutorials, Stack Overflow answers, or sample repositories when they conflict with official Go guidance or idiomatic Go.

---

## Learning-Oriented Explanations

The user has limited experience with Go development.

When making non-trivial decisions, briefly explain:

- why this approach is idiomatic in Go;
- which Go principle, documentation, or best practice it follows;
- what simpler alternatives exist;
- what trade-offs the chosen approach has;
- when this pattern should or should not be used.

Do not over-explain obvious syntax. Focus explanations on architecture, error handling, interfaces, context usage, testing, package design, database access, HTTP handlers, and concurrency.

When suggesting improvements, prefer teaching through concise reasoning and concrete examples.

---

## Core Principles

- Write idiomatic, simple, readable Go.
- Prefer clarity over cleverness.
- Prefer explicit code over magic.
- Keep functions small and focused.
- Avoid unnecessary abstraction.
- Avoid premature optimization.
- Use composition instead of inheritance-like designs.
- Minimize external dependencies.
- Do not add new dependencies without explicit approval.
- Do not change public APIs unless the task requires it.
- Do not introduce global mutable state unless there is a clear and justified reason.
- Prefer explicit dependencies passed through constructors.
- Keep package boundaries clean and intentional.
- Make minimal, targeted changes.
- Prefer patching existing files over replacing them.
- Preserve unrelated code, comments, formatting, and behavior.
- Avoid unrelated refactoring.

---

## Preferred Architecture

Prefer a feature-oriented layout by default.

Use Go idioms, best practices, clean code principles, and clean architecture ideas where they improve clarity, testability, maintainability, or separation of concerns.

Do not force DDD, tactical DDD patterns, heavy domain modeling, hexagonal architecture, or enterprise-style multi-layered architecture unless the project already uses them or the task clearly requires them.

Clean Architecture principles may be used pragmatically, especially for keeping business logic separate from HTTP and database details. However, avoid excessive layers, unnecessary interfaces, abstract factories, generic repositories, and complex domain models when simpler Go code is enough.

The preferred structure is organized around features and business capabilities, not around abstract technical layers.

Example:

```text
cmd/
  api/
    main.go

internal/
  user/
    handler.go
    service.go
    repository.go
    model.go
  order/
    handler.go
    service.go
    repository.go
    model.go
  platform/
    postgres/
    httpserver/
  config/
```

A feature package may contain:

- HTTP handlers
- service/application logic
- repository interfaces or implementations
- request/response DTOs
- database models
- feature-specific helpers
- tests

Keep each feature cohesive and understandable.

---

## Feature Package Rules

Within a feature package, prefer this flow:

```text
HTTP handler
  -> service
  -> repository
  -> PostgreSQL
```

Guidelines:

- Handlers should handle HTTP concerns only.
- Services should contain application logic and orchestration.
- Repositories should contain PostgreSQL access.
- Keep SQL out of handlers.
- Keep HTTP and gorilla/mux-specific code out of services and repositories.
- Do not pass `*http.Request`, `http.ResponseWriter`, `*mux.Router`, or mux route variables outside the HTTP handler layer.
- Pass `context.Context` to services and repositories.
- Keep feature logic close together unless splitting it improves clarity.
- Avoid creating shared packages too early.
- Create shared/internal packages only when code is genuinely reused by multiple features.
- Prefer local feature-specific types over broad shared abstractions.

---

## Editing Existing Files

When modifying an existing file, prefer minimal, targeted edits.

Do not delete and recreate an existing file unless there is a clear technical reason or the user explicitly asks for a full rewrite.

When changing an existing file:

- preserve the file structure where reasonable;
- preserve unrelated code, comments, formatting, and ordering;
- modify only the lines or sections required for the task;
- avoid rewriting the whole file when a small patch is enough;
- avoid reformatting unrelated code;
- avoid renaming, moving, or reorganizing code unless the task requires it;
- keep existing public APIs unchanged unless the task requires a change;
- keep existing behavior unchanged unless the task explicitly requires a behavior change.

If the file appears messy or non-idiomatic, do not rewrite it entirely by default. First explain the issue and propose a focused improvement.

Use full-file rewrites only when:

- the user explicitly requests a full rewrite;
- the file is generated and must be regenerated;
- the file is very small and a full rewrite is safer than patching;
- the existing structure is broken beyond reasonable targeted edits;
- a broad refactor was explicitly requested.

When a full-file rewrite is necessary, explain why before doing it.

---

## Definition of Done

Before considering a task complete, run and pass when the environment allows it:

- `go fmt ./...`
- `goimports -w .` if available
- `go test ./...`
- `go vet ./...`
- `go build ./...`
- `golangci-lint run` if this project already uses `golangci-lint`

Also run `go test ./... -race` when the task touches concurrency, shared state, HTTP handlers, repositories, context cancellation, or code that may be accessed from multiple goroutines.

If a command cannot be run in the current environment, clearly state what was not run and why.

Do not claim that tests, linters, formatting, or builds passed unless they were actually run successfully.

---

## Error Handling

- Handle errors explicitly. Do not ignore returned errors unless there is a clear, documented reason.
- Wrap errors with context using `%w`, for example: `fmt.Errorf("create user: %w", err)`.
- Use `errors.Is` and `errors.As` when callers need to inspect wrapped errors.
- Do not use `panic` for expected runtime failures in business logic or HTTP handlers; return errors and handle them at boundaries.
- Do not log and return the same error unless there is a clear reason.
- Error messages should be concise, actionable, and include useful operation context.
- Do not expose internal errors directly to HTTP clients.
- Convert internal errors to safe HTTP responses at the handler boundary.

---

## Context Usage

- Accept `context.Context` as the first parameter in methods that perform I/O or long-running work, especially service and repository methods.
- Use `ctx` as the context variable name.
- Do not store `context.Context` in structs.
- Do not pass `nil` context.
- Propagate request context through service and repository calls.
- Use `r.Context()` when calling services from HTTP handlers.
- Set timeouts and deadlines at system boundaries, such as HTTP server configuration, request entrypoints, startup tasks, or external calls.
- Reusable business logic should usually accept and propagate context instead of creating its own timeouts.
- Respect cancellation signals (`ctx.Done()`) for long-running operations when applicable.

---

## Interface Guidelines

- Prefer concrete types by default. Introduce interfaces only when they provide clear value.
- Define interfaces at the consumer side, meaning the package that needs behavior, not at the implementation side.
- Do not create interfaces “just in case” or for a single implementation without a concrete need, such as a testing seam, multiple backends, or a package boundary.
- Keep interfaces small and focused, preferably 1–3 methods when practical.
- Accept interfaces and return concrete types where appropriate.
- Avoid broad interfaces with vague names such as `Manager`, `Processor`, or `Service` unless the name has clear meaning in the package.

---

## PostgreSQL Guidelines

- Use context-aware database methods such as `QueryContext`, `QueryRowContext`, and `ExecContext`.
- Use parameterized queries.
- Never build SQL by concatenating untrusted input.
- Avoid `SELECT *`; select only required columns.
- Always close rows when using query result sets.
- Always check and handle `rows.Err()`.
- Handle `sql.ErrNoRows` explicitly.
- Use transactions when multiple database operations must succeed or fail together.
- Keep transactions short and explicit.
- Do not change schema or migrations unless explicitly requested.
- Keep SQL readable and close to the repository code that uses it.
- Be careful with nullable columns and map nullable values explicitly.
- Do not leak PostgreSQL-specific details into HTTP handlers.

---

## HTTP and gorilla/mux Guidelines

- Use standard `net/http` idioms.
- Keep `gorilla/mux` usage inside routing setup, handlers, and middleware.
- Do not pass `*mux.Router`, `*http.Request`, `http.ResponseWriter`, or route-specific mux data into services or repositories.
- Use `r.Context()` when calling services.
- Use `mux.Vars(r)` only in handlers or HTTP-specific helper functions.
- Keep handlers thin.
- Handlers should parse requests, validate transport-level input, call services, and write responses.
- Business logic should live in services, not handlers.
- SQL logic should live in repositories, not handlers.
- Prefer handler methods with the standard signature:

```go
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	// ...
}
```

- Prefer `http.Handler` and `http.HandlerFunc` for middleware and routing integration.
- Middleware should use standard `net/http` patterns:

```go
func ExampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
```

- Parse JSON request bodies with `json.NewDecoder(r.Body)`.
- Always handle JSON decoding errors.
- Avoid reading request bodies multiple times.
- Limit request body size when accepting potentially large payloads.
- Convert request DTOs to service input structs explicitly.
- Convert service results to response DTOs explicitly.
- Use appropriate HTTP status codes.
- Return safe and consistent JSON responses.
- Set `Content-Type: application/json` for JSON responses.
- Do not expose stack traces or internal errors to clients.
- Do not store request-specific values in package-level variables.

Example handler direction:

```go
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid request")
		return
	}

	input := CreateUserInput{
		Email: req.Email,
		Name:  req.Name,
	}

	user, err := h.service.CreateUser(r.Context(), input)
	if err != nil {
		// map error to a safe HTTP response
		return
	}

	writeJSON(w, http.StatusCreated, toUserResponse(user))
}
```

Example route registration direction:

```go
func RegisterUserRoutes(r *mux.Router, h *UserHandler) {
	r.HandleFunc("/users", h.Create).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", h.GetByID).Methods(http.MethodGet)
}
```

---

## Testing

Write tests for meaningful behavior changes.

Tests should verify behavior, not implementation details. Prefer tests that would still be valid after internal refactoring.

### General Testing Principles

- Use Go's standard `testing` package by default.
- Prefer table-driven tests for multiple input/output cases.
- Keep tests deterministic, isolated, and easy to read.
- Test public behavior and important unexported behavior only when it is hard to observe through public APIs.
- Cover both happy paths and error paths.
- Cover boundary cases, invalid input, empty input, nil input where applicable, and permission/validation failures.
- Use clear test names that describe the expected behavior.
- Use `t.Helper()` in test helper functions.
- Avoid sleeps, real time dependencies, random behavior, and external services in unit tests.
- Do not overuse mocks. Prefer simple fakes, stubs, or in-memory implementations when they make tests clearer.
- Do not test private implementation details just to increase coverage.
- Do not write brittle tests that depend on map iteration order, exact timestamps, unrelated formatting, or internal call order.
- Keep assertions explicit and readable.
- When comparing complex values, prefer clear diffs using project-approved tools or standard library approaches.
- Do not add a new testing dependency without approval.

### Service Tests

- Test service behavior independently from HTTP and PostgreSQL when practical.
- Use fake repositories or small in-memory implementations for service tests.
- Verify business behavior, validation, error mapping, and repository interaction outcomes.
- Prefer testing observable results over asserting every internal method call.

### Repository Tests

- Repository tests may be integration tests when they require PostgreSQL.
- Keep repository tests explicit about setup, fixtures, cleanup, and transactions.
- Do not rely on shared database state between tests.
- Use unique test data to avoid cross-test interference.
- Test `sql.ErrNoRows` handling and database constraint errors where relevant.
- Test that context-aware methods respect context cancellation when practical.

### HTTP Handler Tests

- Use `net/http/httptest` for handler tests.
- Test handlers through the router when route variables, middleware, or routing behavior matter.
- Test handlers directly when routing is not relevant.
- Verify status codes, response body, headers, and important error cases.
- For JSON responses, decode the response and compare structured data instead of comparing raw JSON strings when possible.
- For JSON requests, test malformed JSON, missing fields, invalid field values, and valid requests.
- Keep HTTP handlers thin enough that most business behavior can be tested in service tests.

### Table-Driven Test Style

Prefer this style when testing multiple cases:

```go
func TestUserService_CreateUser(t *testing.T) {
	tests := []struct {
		name    string
		input   CreateUserInput
		wantErr bool
	}{
		{
			name: "valid input creates user",
			input: CreateUserInput{
				Email: "user@example.com",
				Name:  "User",
			},
			wantErr: false,
		},
		{
			name: "empty email returns error",
			input: CreateUserInput{
				Email: "",
				Name:  "User",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange

			// act

			// assert
		})
	}
}
```

### HTTP Test Example

```go
func TestUserHandler_Create_ReturnsBadRequestForInvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader("{"))
	rec := httptest.NewRecorder()

	handler := &UserHandler{
		service: fakeUserService{},
	}

	handler.Create(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}
```

### Test Commands

Run tests with:

```sh
go test ./...
```

When concurrency, shared state, handlers, repositories, or context cancellation are involved, also run:

```sh
go test ./... -race
```

---

## Package Design

- Prefer small, cohesive packages.
- Package names should be short, lowercase, and meaningful.
- Avoid generic package names like `common`, `utils`, `helpers`, or `manager`.
- Do not create abstractions just to match a pattern.
- Do not create interfaces unless they are useful.
- Interfaces should usually be defined where they are consumed.
- Accept interfaces and return concrete types where appropriate.
- Keep exported identifiers minimal.
- Prefer unexported types and functions unless they are needed outside the package.
- Avoid circular dependencies.
- Avoid package-level mutable state.

---

## Go Style

- Use meaningful names.
- Avoid one-letter variable names except for conventional cases such as `i`, `j`, `r`, `w`, `tx`, `db`, `ctx`, and `err`.
- Keep variable scope as small as possible.
- Prefer early returns over deep nesting.
- Place `if err != nil { ... }` immediately after the call that returns the error.
- Avoid unnecessary pointers.
- Avoid unnecessary named return values.
- Avoid magic numbers and unexplained constants.
- Use constants for repeated domain values.
- Keep comments useful and accurate.
- Do not comment obvious code.
- Document exported identifiers when required by linting or when documentation improves clarity.
- Prefer standard library solutions when they are sufficient.
- Do not use clever generics where simple functions or structs are clearer.
- Do not use `init()` unless there is a strong reason.
- Do not use `panic` for normal control flow.

---

## Transactions

- Use transactions only when they are needed.
- Keep transactions as short as possible.
- Do not perform slow external operations inside a transaction.
- Always commit or rollback explicitly.
- Ensure rollback is safe when commit succeeds.
- Keep transaction ownership clear.

Example:

```go
tx, err := db.BeginTx(ctx, nil)
if err != nil {
	return fmt.Errorf("begin transaction: %w", err)
}

defer func() {
	_ = tx.Rollback()
}()

// execute queries using tx

if err := tx.Commit(); err != nil {
	return fmt.Errorf("commit transaction: %w", err)
}
```

---

## Concurrency Guidelines

- Use goroutines only when they are necessary.
- Always define ownership and lifecycle of goroutines.
- Avoid goroutine leaks.
- Use context cancellation where appropriate.
- Protect shared mutable state with synchronization primitives.
- Prefer simple synchronization over complex channel patterns.
- Use channels for communication or ownership transfer, not by default for every synchronization problem.
- Always handle errors from goroutines.
- Do not start background goroutines from request handlers unless their lifecycle is well-defined.
- Run race tests when changing concurrent code.

---

## Logging

- Log at application boundaries where useful.
- Do not log sensitive data.
- Do not log passwords, tokens, secrets, session IDs, or personal data.
- Do not use logs as a substitute for error handling.
- Include enough context to debug issues.
- Avoid noisy logs in hot paths.
- Do not log and return the same error unless necessary.

---

## Security

- Validate external input.
- Do not trust request data.
- Use parameterized SQL queries.
- Limit request body sizes where appropriate.
- Do not expose internal error details in HTTP responses.
- Do not log secrets or sensitive data.
- Follow least-privilege principles.
- Avoid hardcoded credentials.
- Avoid insecure defaults.
- Treat authentication and authorization logic as security-sensitive.
- Do not change security behavior unless the task requires it.

---

## Configuration

- Keep configuration explicit.
- Do not hardcode environment-specific values.
- Read configuration at startup.
- Validate required configuration before starting the application.
- Fail fast on invalid critical configuration.
- Keep secrets out of source code.
- Prefer environment variables or the existing project configuration mechanism.
- Do not silently introduce new required configuration.

---

## Dependency Management

- Do not edit `go.mod` or `go.sum` unless explicitly requested or strictly required by the task.
- Do not add new libraries without explicit approval.
- Prefer the standard library whenever it is sufficient.
- If a new dependency is truly necessary, explain why before adding it.
- Do not upgrade dependencies unless explicitly requested.
- Do not vendor dependencies unless the project already uses vendoring.

---

## Slices and Maps

- `nil` slices are acceptable and idiomatic when there is no semantic difference between nil and empty.
- Initialize maps before writing to them.
- Preallocate slices and maps when the expected size is known and it improves clarity or performance.
- Do not rely on map iteration order.
- Return empty slices instead of nil only when API behavior requires it.

---

## Generated Code

- Do not manually edit generated files.
- If generated files must change, update the source definition and regenerate them using the project’s documented command.
- If the generation command is unknown, report the limitation.
- Do not mix generated and handwritten logic in the same file unless the project already does so.

---

## Files Not to Modify Without Explicit Request

Do not modify these unless the user explicitly asks or the task clearly requires it:

- `go.mod`
- `go.sum`
- `migrations/`
- generated files
- vendored files
- CI/CD configuration
- Docker files
- deployment manifests
- production configuration
- public API contracts

When modification is required, keep changes minimal and explain why it was necessary.

---

## External References

When unsure about Go idioms, style, testing, architecture, or library behavior, use official documentation or trusted resources.

Prefer these sources:

1. Official Go documentation
    - Effective Go
    - Go Code Review Comments
    - Go Test Comments
    - Organizing a Go module
    - Standard library documentation
    - Package documentation on `pkg.go.dev`

2. Official library documentation
    - gorilla/mux documentation
    - PostgreSQL driver documentation used by this project
    - Standard library documentation

3. Existing project conventions
    - package structure
    - tests
    - error handling
    - logging
    - configuration
    - dependency injection
    - naming

4. Trusted Go community resources
    - Uber Go Style Guide
    - well-maintained official or widely recognized Go resources

Do not blindly copy patterns from blogs, tutorials, Stack Overflow answers, or sample repositories.

External references should clarify trade-offs. They must not override:

- explicit user requirements
- project-local conventions
- compiler errors
- tests
- linters
- official Go documentation

---

## Verification Priority

Prefer local verification over external references.

Before considering a task complete, run these commands when the environment allows it:

```sh
go fmt ./...
goimports -w .
go vet ./...
go build ./...
go test ./...
golangci-lint run
```

Also run `go test ./... -race` when the task touches concurrency, shared state, HTTP handlers, repositories, context cancellation, or code that may be accessed from multiple goroutines.

If `goimports` or `golangci-lint` is not installed, report it clearly.

If some commands are unavailable or cannot be run in the current environment, clearly report that and explain why.

All code should compile without errors.

If `golangci-lint` reports warnings, fix them instead of ignoring them.

Do not claim that tests, linters, or builds pass unless they were actually run successfully.

---

## Task Workflow

Before editing:

- Inspect the existing project structure.
- Understand the current package boundaries.
- Look for existing patterns that solve similar problems.
- Follow existing naming, layout, error handling, logging, and testing conventions unless they conflict with idiomatic, correct, and maintainable Go.
- Prefer minimal, targeted changes.
- Avoid unrelated refactoring.

While editing:

- Keep changes focused on the requested task.
- Preserve existing behavior unless a behavior change is required.
- Do not introduce broad architectural changes for small tasks.
- Do not add abstractions until they are useful.
- Keep code compilable.

After editing:

- Format the code.
- Clean up imports.
- Run verification commands when possible.
- Add or update tests when behavior changes.
- Summarize what changed.
- Mention any commands that could not be run.
- Mention any risks, assumptions, or follow-up work.
- Briefly explain important Go-specific decisions so the user can learn from them.

---

## What Not To Do

- Do not introduce unnecessary dependencies.
- Do not ignore errors.
- Do not use `panic` for normal error handling.
- Do not put business logic in HTTP handlers.
- Do not put SQL logic in handlers.
- Do not pass HTTP or gorilla/mux-specific types into services or repositories.
- Do not make broad refactors unless requested.
- Do not delete and recreate existing files when targeted edits are enough.
- Do not rewrite entire files just to make small changes.
- Do not reformat or reorder unrelated code.
- Do not silently change configuration, migrations, schemas, or public APIs.
- Do not leave TODOs unless explicitly requested or unavoidable.
- Do not write code that only works for the happy path.
- Do not hide failures.
- Do not claim verification succeeded if it was not run.
- Do not optimize before correctness and clarity.
- Do not create unnecessary interfaces, factories, managers, helpers, or layers.
- Do not force DDD patterns into simple feature-oriented code.
- Do not copy existing project patterns when they are clearly non-idiomatic or unsafe.

---

## Preferred Answer Format for the Agent

When completing a task, respond with:

1. Summary of changes
2. Important Go-specific decisions and why they were made
3. Verification commands run and their results
4. Any commands that could not be run
5. Important assumptions or limitations

Example:

```text
Summary:
- Added user creation endpoint.
- Added user service method.
- Added PostgreSQL repository method.
- Added tests for duplicate email handling.

Go-specific decisions:
- Kept gorilla/mux-specific code inside routing and handlers.
- Used r.Context() to pass context.Context into the service.
- Wrapped repository errors with context using fmt.Errorf and %w.
- Used a table-driven test for service behavior.

Verification:
- go fmt ./...: passed
- goimports -w .: passed
- go vet ./...: passed
- go build ./...: passed
- go test ./...: passed
- golangci-lint run: passed

Notes:
- No migrations were changed.
```

---

## File Link Format (Windows)

When referencing local files in responses, always use clickable Markdown links with an absolute path and optional line:

- Correct: `[repository.go](/C:/goProjects/visualizationDbDebet/internal/contractoranalysis/repository.go:49)`
- If path contains spaces, wrap target in angle brackets:
  - `[My File.go](</C:/goProjects/visualizationDbDebet/some dir/My File.go:10>)`

Rules:

- Use forward slashes `/` in the link target, not backslashes `\`.
- Use project-absolute target starting with `/C:/...`.
- Do not output plain text like `File C:\...` when a clickable link is intended.
- Prefer links whenever mentioning a concrete local file location.

---

## Final Rule

When in doubt, choose the solution that is:

1. Correct
2. Safe
3. Simple
4. Idiomatic Go
5. Easy to test
6. Consistent with official Go guidance
7. Consistent with the existing project when the project is reasonable
8. Easy to change later
