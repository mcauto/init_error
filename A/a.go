package A
import (
	"fmt"
	"init_error/B"
)

var instance *B.InstanceB

func init(){
    instance := B.NewInstanceB()
    fmt.Println(instance)
}

func NewFormat() string{
    data := instance.Data
    return fmt.Sprintf("%s",data)
}
