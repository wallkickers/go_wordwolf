package usecases

type IJoinUseCase interface {
	// ゲームに参加する
	JoinGame(id int, groupId int) error
}
