package impl

import (
	"errors"
	"testing"
	"time"

	"github.com/go-server-dev/src/app/domain"
	"github.com/go-server-dev/src/app/usecase/start_talk"
	"github.com/go-server-dev/src/app/usecase/start_talk/mocks"
)

// 正常系
func TestExecute_withInput_success(t *testing.T) {
	dummyInput := start_talk.Input{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	dummyGameMaster := domain.NewGameMaster(dummyInput.GroupRoomID, domain.GroupRoomType(dummyInput.GroupRoomType))
	dummyOutput := start_talk.Output{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
		TalkTimeMin:   dummyGameMaster.TalkTimeMin(), //トーク時間が返却
		Err:           nil,                           //エラーがnil
	}

	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", dummyInput.GroupRoomID).Return(dummyGameMaster, nil)
	presenterMock := new(mocks.Presenter)
	presenterMock.On("Execute", dummyOutput)
	joinUseCase := NewUseCaseImpl(readOnlyRepositoryMock, presenterMock)

	// 実行
	joinUseCase.Excute(dummyInput)

	// 検証
	readOnlyRepositoryMock.AssertExpectations(t)
	presenterMock.AssertExpectations(t)
}

// 異常系
// ゲームマスター取得失敗
func TestExecute_getGM_fail(t *testing.T) {
	dummyInput := start_talk.Input{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	err := errors.New("get GM Failure")
	dummyOutput := start_talk.Output{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
		Err:           err, //エラー出力
	}

	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", dummyInput.GroupRoomID).Return(nil, err)
	presenterMock := new(mocks.Presenter)
	presenterMock.On("Execute", dummyOutput)
	joinUseCase := NewUseCaseImpl(readOnlyRepositoryMock, presenterMock)

	// 実行
	joinUseCase.Excute(dummyInput)

	// 検証
	readOnlyRepositoryMock.AssertExpectations(t)
	presenterMock.AssertExpectations(t)
}

// 異常系
// 既にトーク中
func TestExecute_alreadyStart_fail(t *testing.T) {
	dummyInput := start_talk.Input{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	dummyGameMaster := domain.NewGameMaster(dummyInput.GroupRoomID, domain.GroupRoomType(dummyInput.GroupRoomType))
	dummyGameMaster.StartTalk()
	err := errors.New("is already Started")
	dummyOutput := start_talk.Output{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
		Err:           err,
	}

	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", dummyInput.GroupRoomID).Return(dummyGameMaster, nil)
	presenterMock := new(mocks.Presenter)
	presenterMock.On("Execute", dummyOutput)
	joinUseCase := NewUseCaseImpl(readOnlyRepositoryMock, presenterMock)

	// 実行
	joinUseCase.Excute(dummyInput)

	// 検証
	readOnlyRepositoryMock.AssertExpectations(t)
	presenterMock.AssertExpectations(t)
}

// 正常系
// 終了タイマー関数の確認
func TestSetFinishTimer(t *testing.T) {
	dummyInput := start_talk.Input{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	dummyOutput := start_talk.FinishTalkOutput{
		ReplyToken:    "testReplyToken",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}

	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	presenterMock := new(mocks.Presenter)
	presenterMock.On("FinishTalk", dummyOutput)
	joinUseCase := NewUseCaseImpl(readOnlyRepositoryMock, presenterMock)

	// 実行
	setFinishTimer(joinUseCase, dummyInput, time.Second*1)

	// 検証
	presenterMock.AssertExpectations(t)
}
