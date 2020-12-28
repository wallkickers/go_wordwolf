package impl

import (
	"testing"

	"github.com/go-server-dev/src/app/domain"
	acceptVotes "github.com/go-server-dev/src/app/usecase/accept_votes"
	"github.com/go-server-dev/src/app/usecase/accept_votes/mocks"
	"github.com/stretchr/testify/assert"
)

// 正常系
func TestExecute_accept_votes_withInput_success(t *testing.T) {
	// ダミーデータ
	dummyInput := acceptVotes.Input{
		ReplyToken:    "testReplyToken",
		FromMemberID:  "testMemberIDFrom",
		ToMemberID:    "testMemberIDTo",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	dummyGameMaster := domain.NewGameMaster(dummyInput.GroupRoomID, domain.GroupRoomType(dummyInput.GroupRoomType))
	dummyGameMaster.SetVoteManagement(dummyInput.FromMemberID, dummyInput.ToMemberID)
	dummyOutput := acceptVotes.Output{
		ReplyToken:    "testReplyToken",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	// 参照用リポジトリMock
	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", dummyInput.GroupRoomID).Return(dummyGameMaster, nil)
	// 更新用リポジトリMock
	gameMasterRepositoryMock := new(mocks.GameMasterRepository)
	gameMasterRepositoryMock.On("Save", dummyGameMaster).Return(nil)
	// プレゼンターMock
	presenterMock := new(mocks.Presenter)
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
