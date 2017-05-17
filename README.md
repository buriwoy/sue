# Simple Unique Epoch

"**S**imple **U**nique **E**poch" (SUE) is a simple id generator based on UNIX epoch time.

## Install

```bash
  go get github.com/buriwoy/sue
```

## Usage

```go

// a unique id based on microseconds
microId := sue.New()

// sample output: "3L4CTbdm"
fmt.Println(microId)



// a unique id based on nanoseconds
microId := sue.New()

// sample output: "rrnlFXBWL"
fmt.Println(nanoId)

```
