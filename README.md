# Simple Unique Epoch

"**S**imple **U**nique **E**poch" (SUE) is a simple id generator based on UNIX epoch time.

## Install

  go get github.com/buriwoy/sue

## Usage

```go
microId := sue.New()             // a unique id based on microseconds
fmt.Println(microId)             // sample output: "3L4CTbdm"

microId := sue.New()             // a unique id based on nanoseconds
fmt.Println(nanoId)              // sample output: "rrnlFXBWL"
```
