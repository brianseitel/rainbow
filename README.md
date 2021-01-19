# rainbow

## Description

A simple library to match color names and hex codes.

### Example: Find a Color by Hex

```go
c, _ := colorful.Hex("#ff0808")
name, distance := color.Match(c)

// returns "candy apple red" 98.05212330136119
```

### Example: Find Hex by Color

```go
hex, err := color.Find("pink")
if err != nil {
    // handle
}

// hex is "#ffc0cb"

## Install

`go get -u github.com/brianseitel/rainbow`


