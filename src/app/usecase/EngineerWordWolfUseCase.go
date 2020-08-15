type EngineerWordWolfUseCase struct {
	memberAction *MemberAction
	gameMasterAction *GameMasterAction
}

func NewEngineerWordWolfUseCase(MemberAction MemberAction,GameMasterAction GameMasterAction) *EngineerWordWolfUseCase{
	return $EngineerWordWolfUseCase {
		MemberAction: MemberAction,
		GameMasterAction: GameMasterAction
	}
}

// ゲームに参加する
func (u *EngineerWordWolfUseCase) joinGame(id int, groupId int) error {
	result, err := memberAction.JoinGame(id, groupId)
	if err != nil {
		return err
	}
	return nil
}

// 参加者を打ち切り、お題を割り振る
func (u *EngineerWordWolfUseCase) startGame(groupId int) error {
	result, err := gameMasterAction.AssignTheme(groupId)
	if err != nil {
		return err
	}
	return nil
}

// トークを開始する
func (u *EngineerWordWolfUseCase) startTalk(groupId int) error {
	result, err := gameMasterAction.StartTalk(groupId)
	if err != nil {
		return err
	}
	result, err := gameMasterAction.StartToMesureTime(groupId)
	if err != nil {
		return err
	}
	result, err := gameMasterAction.GetLimitTime(groupId)
	if err != nil {
		return err
	}
	return nil
}

// 残り時間を表示する
func (u *EngineerWordWolfUseCase) showRemainingTime(groupId int) (string,error) {
	result, err := gameMasterAction.GetLimitTime(groupId)
	if err != nil {
		return (nil,err)
	}
	return (result,nil)
}

// 投票を受け付ける
func (u *EngineerWordWolfUseCase) acceptVotes(groupId int) error {
	result, err := memberAction.Vote(id, groupId, id)
	if err != nil {
		return err
	}
	return nil
}

// 結果を表示する
func (u *EngineerWordWolfUseCase) displayResult(groupId int) error {
	result, err := gameMasterAction.GetResult(groupId)
	if err != nil {
		return err
	}
 	return nil
}
