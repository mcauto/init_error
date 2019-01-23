package B

type InstanceB struct {
	Data string
}

func NewInstanceB() *InstanceB{
  return &InstanceB{Data:"hello world!"}
}

