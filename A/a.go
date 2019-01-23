package A
import (
	"fmt"
	"init_error/B"
)

var instance *B.InstanceB

func init(){
    instance = B.NewInstanceB() // reference count의 변화가 없을 것
    fmt.Println(instance)
}

func NewFormat() string{
    data := instance.Data
    return fmt.Sprintf("%s",data)
}
