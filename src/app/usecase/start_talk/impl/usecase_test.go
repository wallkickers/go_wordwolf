package impl

import (
	"testing"

	"github.com/go-server-dev/src/app/domain"
	"github.com/go-server-dev/src/app/usecase/join"
	"github.com/go-server-dev/src/app/usecase/join/mocks"
	"github.com/stretchr/testify/assert"
)

// 正常系
func TestExecute_withInput_success(t *testing.T) {
	// ダミーデータ
	dummyInput := join.Input{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	dummyMember := domain.NewMember(dummyInput.MemberID, "testMemberName_Taro")
	dummyGameMaster := domain.NewGameMaster(dummyInput.GroupRoomID, domain.GroupRoomType(dummyInput.GroupRoomType))
	dummyGameMaster.SetMember(dummyInput.MemberID)
	dummyOutput := join.Output{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		MemberName:    dummyMember.Name(),
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	// 参照用リポジトリMock
	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	readOnlyRepositoryMock.On("FindMemberByID", dummyInput.MemberID).Return(dummyMember, nil)
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", dummyInput.GroupRoomID).Return(dummyGameMaster, nil)
	// 更新用リポジトリMock
	gameMasterRepositoryMock := new(mocks.GameMasterRepository)
	gameMasterRepositoryMock.On("Save", dummyGameMaster).Return(nil)
	// プレゼンターMock
	presenterMock := new(mocks.Presenter)
	presenterMock.On("Execute", dummyOutput)
	joinUseCase := UseCaseImpl{
		gameMasterRepository: gameMasterRepositoryMock,
		readOnlyRepository:   readOnlyRepositoryMock,
		presenter:            presenterMock,
	}

	// 実行
	expect := dummyOutput
	actual := joinUseCase.Excute(dummyInput)

	// 検証
	readOnlyRepositoryMock.AssertExpectations(t)
	gameMasterRepositoryMock.AssertExpectations(t)
	presenterMock.AssertExpectations(t)
	assert.Equal(t, expect, actual, "they should be equal output")
}
