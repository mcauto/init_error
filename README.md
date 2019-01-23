# init_error
golang init function error 

```go
package A
import (
	"fmt"
	"init_error/B"
)

var instance *B.InstanceB

func init(){
    instance := B.NewInstanceB() // package level scope의 instance를 사용하는 것이 아닌 init()의 local scope로 변수가 새로 생성됨
    instance = B.NewInstanceB() // 따라서 :=가 아닌 =를 이용하여 package level 변수를 사용하도록 하면 된다.
    fmt.Println(instance)
}

func NewFormat() string{
    data := instance.Data
    return fmt.Sprintf("%s",data)
}
```
