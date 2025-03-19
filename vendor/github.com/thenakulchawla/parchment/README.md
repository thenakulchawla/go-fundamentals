Parchment is a lightweight, structured logging library designed to persist log context throughout the lifecycle of a request or process. Inspired by the historical practice of recording events on parchment, Parch ensures that logs remain structured, contextual, and seamlessly available across function calls—without requiring developers to manually pass loggers around.

# Create a new logger using your context
```go
log := parchment.New(context.Background())
```

# Add values to logger
```go
ctx = AddToLogger(ctx, []LoggerField{
		{Key: "first_key", Value: "first_value"},
	})
```

# Get logger from context
```go
log := parchment.FromContext(ctx)
```


