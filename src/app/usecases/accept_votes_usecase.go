package usecases

type IWordWolfUseCase interface {
	// ゲームに参加する
	JoinGame(id int, groupId int) error

	// 参加者を打ち切り、お題を割り振る
	StartGame(groupId int) error

	// トークを開始する
	StartTalk(groupId int) error

	// 残り時間を表示する
	ShowRemainingTime(groupId int) (string, error)

	// 投票を受け付ける
	AcceptVotes(groupId int) error

	// 結果を表示する
	DisplayResult(groupId int) error
}
