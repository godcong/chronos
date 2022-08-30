package chronos

// solar ...
type solar struct {
}

func (s *solar) Year() int {
	//TODO implement me
	panic("implement me")
}

func (s *solar) Month() int {
	//TODO implement me
	panic("implement me")
}

func (s *solar) Day() int {
	//TODO implement me
	panic("implement me")
}

func (s *solar) Hour() int {
	//TODO implement me
	panic("implement me")
}

var _ Solar = &solar{}
