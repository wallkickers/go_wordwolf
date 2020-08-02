package domain

type GameMember struct {
	groupId int
	accountList []int
	themeManagement map[int]string
	voteManagement map[int]int
	startTime time
	talkTime　time
}

func NewGameMember(groupId int) *GameMember {
	return &GameMember{
		groupId: groupId
	}
}

func (g *GameMember) GetGroupId() int {
	return g.groupId
}

func (g *GameMember) SetTalkTime(time) {
	g.talktime = time
}

// インターフェイス
type GameMemberAction interface {
	AssignTheme(id int) error
	StartToMesureTime() error
	GetLimitTime() error
	GetResult() error
	AddMember(id) error
	ManageVote(id, id) error
}
