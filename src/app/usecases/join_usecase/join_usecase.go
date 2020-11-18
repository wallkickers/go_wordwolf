package join_usecase

type JoinUseCase interface {
	// ゲームに参加する
	JoinGame(JoinInputDto) JoinOutputDto
}
