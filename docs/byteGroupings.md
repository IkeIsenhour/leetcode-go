In Go, Bytes are stores as their ASCII equivalent value. This allows us to iterate over characters as if they were numbers, because under the hood
Go is incrementing the ASCII value.

Loop over every lowercase letter

```go
for char := 'a'; char <= 'z'; char++ {
    fmt.Printf("%c ", char)
}
```

Loop over every uppercase letter

```go
for char := 'A'; char <= 'Z'; char++ {
    fmt.Printf("%c ", char)
}
```

Loop over every digit

```go
for num := 0; num <= 9; num++ {
    fmt.Printf("%v ", num)
}
```

Using this logic, we can loop over a string and remove all non-alphanumeric characters

```go
func strip(s string) string {
    var result strings.Builder

    for i := 0; i < len(s); i++ {
        curr := s[i]
        if ('a' <= curr && curr <= 'z') || ('A' <= curr && curr <= 'Z') || ('0' <= curr && curr <= '9') {
            result.WriteByte(curr)
        }
    }

    return result.String()
}
```
