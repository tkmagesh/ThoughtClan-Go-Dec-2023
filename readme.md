# Go Foundation

## Magesh Kuppan

## Schedule
- Commence      : 9:00 AM
- Tea Break     : 10:30 AM (20 mins)
- Lunch Break   : 12:30 PM (1 hr)
- Tea Break     : 3:00 PM (20 mins)
- Wind up       : 5:00 PM

## Methodology
- No powerpoint
- 100% Code driven

## Software Requirements
- Go Tools (https://go.dev/dl)
- Visual Studio Code (https://code.visualstudio.com)
- Go Extension for VSCode
- Docker Desktop

## Repository
- https://github.com/tkmagesh/thoughtclan-go-dec-2023

## Why Go?
- Simplicity
    - ONLY 25 keywords
    - No access modifiers (public/private/protected)
    - No classes (only structs)
    - No inheritance (only composition)
    - No function overloading
    - No implicit type conversion
    - No exceptions (only errors)
    - No "try catch finally" constructs
    - No pointer arithmatic
- Performance
    - Equivalent to C++
    - Close to the hardware
        - Compiled to machine code
        - Platform specific builds are needed
        - No JIT compilation
    - Support for cross compilation
- Managed Concurrency
    - Concurrency
        - Ability to have more than one execution path
    - Built in Scheduler
        - N : M scheduler where N = # of goroutines & M = # of OS threads
        - N >= M
    - Concurrent operations are represented as Goroutines (cheap, ONLY 4KB)
    - Concurrency features are built in the language
        - go keyword, channel data type, channel operator ( <- ), range, select-case constructs
    - APIs support
        - "sync" package
        - "sync/atomic" package

## Go Programs
### To compile & execute
> go run [filename.go]
### To compile
- > go build [filename.go]
- > go build -o [binary_name] [filename.go]
### To list the environment variables
- > go env
- > go env [var_1] [var_2]...
### To update the environment variables
- > go env -w [var_1 = val_1] [var_2 = val_2]...
### Environment variables for cross compilation
- GOARCH
- GOOS
### To get the supported platforms for cross compilation
> go tool dist list
### To cross compile
> GOOS=[target_os] GOARCH=[target_arch] go build [filename.go]

## Data Types
- bool
- string
- integer types
    - int8
    - int16
    - int32
    - int64
    - int
- unsigned integer types
    - uint8
    - uint16
    - uint32
    - uint64
    - uint
- float types
    - float32
    - float64
- complex types
    - complex64 (real[float32] + imaginary[float32])
    - complex128 (real[float64] + imaginary[float64])
- type aliases
    - byte (alias for uint8)
    - rune (alias for int32 / unicode code point)
### Variable Declaration
- Using "var" keyword
- Using ":="
- Function scoped variables
    - Can use ":="
    - Cannot have unused variables
- Package scoped variables
    - Cannot use ":="
    - Can have unused variables
### Programming Constructs
## if else

## switch case

## for

## func
