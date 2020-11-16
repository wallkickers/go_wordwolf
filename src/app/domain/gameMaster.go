package domain

import (
	"time"
)

type GameMaster struct {
	groupId         int
	accountList     []int
	themeManagement map[int]string
	voteManagement  map[int]int
	startTime       time.Time
	talkTime        time.Time
}

func NewGameMaster(groupId int) *GameMaster {
	return &GameMaster{
		groupId: groupId,
	}
}

func (g *GameMaster) GetGroupId() int {
	return g.groupId
}

func (g *GameMaster) SetTalkTime(time time.Time) {
	g.talkTime = time
}

// インターフェイス
type GameMasterAction interface {
	AssignTheme(groupId int) error
	StartTalk(groupId int) error
	StartToMesureTime(groupId int) error
	GetLimitTime(groupId int) (time.Time, error)
	GetResult(groupId int) error
	AddMember(id, groupId int) error
	//ManageVote(id, id int) error
}
