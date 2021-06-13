package internal

var _ RecoverService = new(RecoverServiceImpl)

type RecoverServiceImpl struct {
	err interface{}
}

func (s *RecoverServiceImpl) Panic(err interface{}) {
	if s == nil {
		return
	}
	s.err = err
}

func (s *RecoverServiceImpl) Recover() interface{} {
	if s == nil {
		return nil
	}
	return s.err
}
