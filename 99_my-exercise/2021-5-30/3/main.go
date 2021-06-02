type server struct{
	Idx int
	Weight int
	CompleteTime int
}

type baseQueue []*server

func (b baseQueue) Len()int  {
	return Len(baseQueue)
}