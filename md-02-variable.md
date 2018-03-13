```go
package main

import "fmt"

func main() {
    // Long Declaration
    var x string = "Hello, World"
    fmt.Println(x)

    var y string
    y = "Hello, World"
    fmt.Println(y)

    // Short Declaration
    // Type Inference
    z := "Hello, World"
    fmt.Println(z)
    fmt.Printf("Type: %T\n", z)
}
```


JAVA
```java
String greeting = "Hello";
```
Go
```go
var greeting String = "Hello"
```
```
Declares a variable named
greeting is a String with a value of "Hello"
```
  
  
Scope
```go
package main

import "fmt"

func main() {
    var x string = "Hello, World"
    fmt.Println(x)
}
```
  
another way
```go
package main

import "fmt"

var x string = "Hello, World"

func main() {
    fmt.Println(x)
}
```

```go
package main

import "fmt"

var x string = "Hello, World"

func main() {
    fmt.Println(x)
}

func f() {
    fmt.Println(x)
}
```

```go
package main

import "fmt"

func main() {
    var x string = "Hello, World"
    fmt.Println(x)
}

func f() {
    fmt.Println(x)
}
```
What's wrong?  

Constants
```go
package main

func main() {
    const x string = "Hello, World"
    x = "Other string"
}
```
What's happened?  


Define Multiple Variables
```go
var (
    a = 5
    b = 10
    c = 15
)
```

swapping variable
```go
foo, bar := 1, 2
foo, bar = bar, foo
```