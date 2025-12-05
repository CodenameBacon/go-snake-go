package game

type StateServer interface {
	SendPublicState(sessionModel *SessionModel)
}

type SoloStateServer struct {
	StateChan chan *SessionModel
}

func (s *SoloStateServer) SendPublicState(sessionModel *SessionModel) {
	s.StateChan <- sessionModel
}
