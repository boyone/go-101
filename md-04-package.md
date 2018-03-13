### enough with main package :)  
#### It has so many responsibility.  
#### We have to move behaviours to package that it's belong.  

1. Type go env and looking for GOPATH  
```
>go env
```

2. Create new folder call vendingMachine  

3. Create new file call engine.go  

4. open file exercies/vending-machine.go  

5. move every single func and struct to engine.go and the first line add:   
```
package vendingMachine
```

vending-machine.go should have only func main   
```
pacage main

import "fmt"

func main() {
    ...
}
```

### It's time to import vending-machine.  

6. import vendingMachine to vending-machine.go
```
import (
    "fmt"
    "./vendingMachine"
)
```
then add vendingMachine in front of NewVendingMachine()  
```
vm := vendingMachine.NewVendingMachine(coins, items)
```

### It's work but not the right way.

7. move vendingMachine folder to GOPATH/src/   
```
GOPATH
    |- src
        |- vendingMachine
```

8. chang import in vending-machine.go  
from  
```
import (
    "fmt"
    "./vendingMachine"
)

to  

```
import (
    "fmt"
    "vendingMachine"
)
```
```

### share code with your team  

9. move vendingMachine folder to github.com/  
```
GOPATH
    |- src
        |- github.com
            |- your git username
                |- vendingMachine
```

example  
```
GOPATH
    |- src
        |- github.com
            |- boyone
                |- vendingMachine
```

10. change import in vending-machine.go  
```
import (
    "fmt"
    "github.com/boyone/vendingMachine"
)
```

11. change directory to $GOPATH/src/github.com/boyone/vendingMachine  
12. type git init  
13. new repository call vendingMachine  
14. add remote  
15. push source code to github  
16. remove vendingMachine folder
17. go get vendingMachine  
```
>go get github.com/<username>/vendingMachine
```