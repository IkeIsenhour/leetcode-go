This is a way to use a slice as a stack in Go, without actually have a stack implementation.

1. Initialize a slice

```go
stack := []string{}
```

2. Add to the stack

```go
stack = append(stack, "INSERT STRING")
```

3. Pop from the stack. This code removes the last index from the slice.

```go
if len(stack) > 0 {
    stack = stack[:len(stack) - 1]
}
```

Full code block below:

```go
// Initialize stack
stack := []string{}

// Add to stack
stack = append(stack, "INSERT STRING")

// Pop from stack
if len(stack) > 0 {
    stack = stack[:len(stack) - 1]
}
```
