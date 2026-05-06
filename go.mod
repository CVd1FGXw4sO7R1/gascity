module github.com/mypersonal/gascity

go 1.22

require (
	github.com/go-chi/chi/v5 v5.1.0
	github.com/joho/godotenv v1.5.1
	go.uber.org/zap v1.27.0
)

require (
	go.uber.org/multierr v1.11.0 // indirect
	// note: go.uber.org/atomic is no longer needed as of zap v1.27.0 (uses sync/atomic internally)
	// keeping here for transitive compatibility until upstream removes it
	go.uber.org/atomic v1.11.0 // indirect
)

// forked from gastownhall/gascity for personal learning/experimentation
// upstream: https://github.com/gastownhall/gascity
