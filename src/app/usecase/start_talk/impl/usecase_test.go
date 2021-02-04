package impl

import (
	"errors"
	"testing"
	"time"

	"github.com/go-server-dev/src/app/domain"
	"github.com/go-server-dev/src/app/usecase/start_talk"
	"github.com/go-server-dev/src/app/usecase/start_talk/mocks"
)

/*
 * [正常系]
 */
func TestExecute_withInput_success(t *testing.T) {
	// テスト対象のインスタンス作成
	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	gameMasterRepositoryMock := new(mocks.GameMasterRepository)
	presenterMock := new(mocks.Presenter)
	joinUseCase := NewUseCaseImpl(gameMasterRepositoryMock, readOnlyRepositoryMock, presenterMock)

	dummyInput := start_talk.Input{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	dummyGameMaster := domain.NewGameMaster(dummyInput.GroupRoomID, domain.GroupRoomType(dummyInput.GroupRoomType))
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", dummyInput.GroupRoomID).Return(dummyGameMaster, nil)
	gameMasterRepositoryMock.On("Save", dummyGameMaster).Return(nil)
	dummyOutput := start_talk.Output{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
		TalkTimeMin:   dummyGameMaster.TalkTimeMin(), //トーク時間が返却
		Err:           nil,                           //エラーがnil
	}
	presenterMock.On("Execute", dummyOutput)

	// 実行
	joinUseCase.Excute(dummyInput)

	// Mockの入力値を検証
	readOnlyRepositoryMock.AssertExpectations(t)
	gameMasterRepositoryMock.AssertExpectations(t)
	presenterMock.AssertExpectations(t)
}

/*
 * [異常系]
 * システムエラー：ゲームマスター取得失敗
 */
func TestExecute_getGM_fail(t *testing.T) {
	// テスト対象のインスタンス作成
	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	gameMasterRepositoryMock := new(mocks.GameMasterRepository)
	presenterMock := new(mocks.Presenter)
	joinUseCase := NewUseCaseImpl(gameMasterRepositoryMock, readOnlyRepositoryMock, presenterMock)

	dummyInput := start_talk.Input{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	err := errors.New("get GM Failure")
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", dummyInput.GroupRoomID).Return(nil, err)
	dummyOutput := start_talk.Output{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
		Err:           err, //エラー出力
	}
	presenterMock.On("Execute", dummyOutput)

	// 実行
	joinUseCase.Excute(dummyInput)

	// Mockの入力値を検証
	readOnlyRepositoryMock.AssertExpectations(t)
	presenterMock.AssertExpectations(t)
}

/*
 * [異常系]
 * システムエラー：ゲームマスター保存失敗
 */
func TestExecute_saveGM_fail(t *testing.T) {
	// テスト対象のインスタンス作成
	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	gameMasterRepositoryMock := new(mocks.GameMasterRepository)
	presenterMock := new(mocks.Presenter)
	joinUseCase := NewUseCaseImpl(gameMasterRepositoryMock, readOnlyRepositoryMock, presenterMock)

	dummyInput := start_talk.Input{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	err := errors.New("save GM Failure")
	dummyGameMaster := domain.NewGameMaster(dummyInput.GroupRoomID, domain.GroupRoomType(dummyInput.GroupRoomType))
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", dummyInput.GroupRoomID).Return(dummyGameMaster, nil)
	gameMasterRepositoryMock.On("Save", dummyGameMaster).Return(err)
	dummyOutput := start_talk.Output{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
		Err:           err, //エラー出力
	}
	presenterMock.On("Execute", dummyOutput)

	// 実行
	joinUseCase.Excute(dummyInput)

	// Mockの入力値を検証
	readOnlyRepositoryMock.AssertExpectations(t)
	presenterMock.AssertExpectations(t)
}

/*
 * [異常系]
 * 業務エラー：既にトーク中
 */
func TestExecute_alreadyStart_fail(t *testing.T) {
	// テスト対象のインスタンス作成
	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	gameMasterRepositoryMock := new(mocks.GameMasterRepository)
	presenterMock := new(mocks.Presenter)
	joinUseCase := NewUseCaseImpl(gameMasterRepositoryMock, readOnlyRepositoryMock, presenterMock)

	dummyInput := start_talk.Input{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	dummyGameMaster := domain.NewGameMaster(dummyInput.GroupRoomID, domain.GroupRoomType(dummyInput.GroupRoomType))
	dummyGameMaster.StartTalk()
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", dummyInput.GroupRoomID).Return(dummyGameMaster, nil)
	err := errors.New("is already Started")
	dummyOutput := start_talk.Output{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
		Err:           err,
	}
	presenterMock.On("Execute", dummyOutput)

	// 実行
	joinUseCase.Excute(dummyInput)

	// Mockの入力値を検証
	readOnlyRepositoryMock.AssertExpectations(t)
	presenterMock.AssertExpectations(t)
}

/*
 * [正常系]
 * 終了通知タイマーのテスト
 */
func TestSetFinishTimer(t *testing.T) {
	// テスト対象のインスタンス作成
	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	gameMasterRepositoryMock := new(mocks.GameMasterRepository)
	presenterMock := new(mocks.Presenter)
	joinUseCase := NewUseCaseImpl(gameMasterRepositoryMock, readOnlyRepositoryMock, presenterMock)

	dummyInput := start_talk.Input{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	dummyGameMaster := domain.NewGameMaster(dummyInput.GroupRoomID, domain.GroupRoomType(dummyInput.GroupRoomType))
	dummyGameMaster.StartTalk() //終了時間通知前はトーク中
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", dummyInput.GroupRoomID).Return(dummyGameMaster, nil)
	dummyGameMasterAfter := domain.NewGameMaster(dummyInput.GroupRoomID, domain.GroupRoomType(dummyInput.GroupRoomType))
	dummyGameMasterAfter.EndTalk() //終了時間通知後はトーク終了
	gameMasterRepositoryMock.On("Save", dummyGameMasterAfter).Return(nil)
	dummyOutput := start_talk.FinishTalkOutput{
		ReplyToken:    "testReplyToken",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
		Err:           nil,
	}
	presenterMock.On("FinishTalk", dummyOutput)

	// 実行
	setFinishTimer(joinUseCase, dummyInput, time.Microsecond*1)

	// Mockの入力値を検証
	readOnlyRepositoryMock.AssertExpectations(t)
	gameMasterRepositoryMock.AssertExpectations(t)
	presenterMock.AssertExpectations(t)
}

/*
 * [異常系]
 * 終了通知タイマーのテスト
 * ゲームマスター取得失敗
 */
func TestSetFinishTimer_getGM_failure(t *testing.T) {
	// テスト対象のインスタンス作成
	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	gameMasterRepositoryMock := new(mocks.GameMasterRepository)
	presenterMock := new(mocks.Presenter)
	joinUseCase := NewUseCaseImpl(gameMasterRepositoryMock, readOnlyRepositoryMock, presenterMock)

	dummyInput := start_talk.Input{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	err := errors.New("get GM Failure")
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", dummyInput.GroupRoomID).Return(nil, err)
	dummyOutput := start_talk.FinishTalkOutput{
		ReplyToken:    "testReplyToken",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
		Err:           err,
	}
	presenterMock.On("FinishTalk", dummyOutput)

	// 実行
	setFinishTimer(joinUseCase, dummyInput, time.Microsecond*1)

	// Mockの入力値を検証
	readOnlyRepositoryMock.AssertExpectations(t)
	presenterMock.AssertExpectations(t)
}

/*
 * [異常系]
 * 終了通知タイマーのテスト
 * ゲームマスター保存失敗
 */
func TestSetFinishTimer_saveGM_failure(t *testing.T) {
	// テスト対象のインスタンス作成
	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	gameMasterRepositoryMock := new(mocks.GameMasterRepository)
	presenterMock := new(mocks.Presenter)
	joinUseCase := NewUseCaseImpl(gameMasterRepositoryMock, readOnlyRepositoryMock, presenterMock)

	dummyInput := start_talk.Input{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	dummyGameMaster := domain.NewGameMaster(dummyInput.GroupRoomID, domain.GroupRoomType(dummyInput.GroupRoomType))
	dummyGameMaster.StartTalk()
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", dummyInput.GroupRoomID).Return(dummyGameMaster, nil)
	err := errors.New("save GM Failure")
	gameMasterRepositoryMock.On("Save", dummyGameMaster).Return(err)
	dummyOutput := start_talk.FinishTalkOutput{
		ReplyToken:    "testReplyToken",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
		Err:           err,
	}
	presenterMock.On("FinishTalk", dummyOutput)

	// 実行
	setFinishTimer(joinUseCase, dummyInput, time.Microsecond*1)

	// Mockの入力値を検証
	readOnlyRepositoryMock.AssertExpectations(t)
	presenterMock.AssertExpectations(t)
}
