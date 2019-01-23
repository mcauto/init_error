package A
import (
	"fmt"
	"init_error/B"
)

var instance *B.InstanceB

func init(){
    var err error
    instance, err = B.NewInstanceB()
    fmt.Println(instance)
}

func NewFormat() string{
    data := instance.Data
    return fmt.Sprintf("%s",data)
}
