package game

type StateServer interface {
	SendPublicState(sessionModel *SessionModel)
}

type SoloStateServer struct{}

func (s *SoloStateServer) SendPublicState(sessionModel *SessionModel) {
	// simply writes into the stdout
	drawGameField(sessionModel)
}
