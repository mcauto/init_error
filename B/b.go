package B

type InstanceB struct {
	Data string
}

func NewInstanceB() (*InstanceB, error){
  return &InstanceB{Data:"hello world!"}, nil
}

