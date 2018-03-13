#### For
```
for i < 10 {}
```
```
for i:= 0; i < 10; i++ {}
```
```
for { 
    .
    break
}
```
```
for key, value := range kvs {}
``` 

#### If/Else
```
if i % 2 == 0 {
    // even
} else {
    // odd
}
```
```
i := f()
if i > 0 {
    do something
}
```

```
if i := f(); i > 0 {
    do something
}
```
```
if i > 0 {
    ...
} else if i < 0 {
    ...
} else {
    ...
}
```


#### Switch
```
i := num % 2
switch i {
    case 0: fmt.Println("even")
    default: fmt.Println("odd")
}
```
```
switch {
    case i % 2 == 0:  
        fmt.Println("even")
    default:
        fmt.Println("odd")
}
```
```
switch i := num % 2 {
    case 0: fmt.Println("even")
    default: fmt.Println("odd")
}
```
