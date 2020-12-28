package impl

import (
	"github.com/go-server-dev/src/app/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type readOnlyRepositoryMock struct {
	mock.Mock
}

func (m *readOnlyRepositoryMock) FindGameMasterByGroupID(groupID string) (*domain.GameMaster, error) {
	args := m.Called(groupID)
	return args.(*domain.GameMaster), args.Error(1)
}

type gameMasterRepositoryMock struct {
	mock.Mock
}

func (m *gameMasterRepositoryMock) Save(gameMaster *domain.GameMaster) error {
	args := m.Called(gameMaster)
	return args.Error(0)
}

type presenterMock struct {
	mock.Mock
}

func (m *presenterMock) Execute(output Output) error {
	ret := m.Called(output)

	var r0 error
	if rf, ok := ret.Get(0).(func(Output) error); ok {
		r0 = rf(output)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// 正常系
func TestExecute_accept_votes_withInput_success(t *testing.T) {
	// ダミーデータ
	dummyInput := Input{
		ReplyToken:    "testReplyToken",
		FromMemberID:  "testMemberIDFrom",
		ToMemberID:    "testMemberIDTo",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	dummyGameMaster := domain.NewGameMaster(dummyInput.GroupRoomID, domain.GroupRoomType(dummyInput.GroupRoomType))
	dummyGameMaster.SetVoteManagement(dummyInput.FromMemberID, dummyInput.ToMemberID)
	dummyOutput := Output{
		ReplyToken:    "testReplyToken",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	// 参照用リポジトリMock
	readOnlyRepositoryMock := new(readOnlyRepositoryMock)
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", dummyInput.GroupRoomID).Return(dummyGameMaster, nil)
	// 更新用リポジトリMock
	gameMasterRepositoryMock := new(gameMasterRepositoryMock)
	gameMasterRepositoryMock.On("Save", dummyGameMaster).Return(nil)
	// プレゼンターMock
	presenterMock := new(presenterMock)
	presenterMock.On("Execute", dummyOutput)
	// mockの注入
	acceptVotesUseCase := UseCaseImpl{
		gameMasterRepository: gameMasterRepositoryMock,
		readOnlyRepository:   readOnlyRepositoryMock,
		presenter:            presenterMock,
	}

	// 実行
	expect := dummyOutput
	actual := acceptVotesUseCase.Excute(dummyInput)

	// 検証
	readOnlyRepositoryMock.AssertExpectations(t)
	gameMasterRepositoryMock.AssertExpectations(t)
	presenterMock.AssertExpectations(t)
	assert.Equal(t, expect, actual, "they should be equal output")
}
