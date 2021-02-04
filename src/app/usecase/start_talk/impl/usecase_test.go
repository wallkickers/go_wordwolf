package impl

import (
	"errors"
	"testing"
	"time"

	"github.com/go-server-dev/src/app/domain"
	"github.com/go-server-dev/src/app/usecase/start_talk"
	"github.com/go-server-dev/src/app/usecase/start_talk/mocks"
	"github.com/stretchr/testify/assert"
)

// 正常系
func TestExecute_withInput_success(t *testing.T) {
	// ダミーデータ
	dummyInput := start_talk.Input{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
	}
	dummyOutput := start_talk.Output{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "group",
		Err:           nil,
	}
	dummyGameMaster := domain.NewGameMaster(dummyInput.GroupRoomID, domain.GroupRoomType(dummyInput.GroupRoomType))
	dummyOutput.TalkTimeMin = dummyGameMaster.TalkTimeMin()

	// 参照用リポジトリMock
	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", dummyInput.GroupRoomID).Return(dummyGameMaster, nil)

	// プレゼンターMock
	presenterMock := new(mocks.Presenter)
	presenterMock.On("Execute", dummyOutput).Return(nil)
	joinUseCase := UseCaseImpl{
		readOnlyRepository: readOnlyRepositoryMock,
		presenter:          presenterMock,
	}

	// 実行
	expect := dummyOutput
	actual := joinUseCase.Excute(dummyInput)

	// 検証
	readOnlyRepositoryMock.AssertExpectations(t)
	presenterMock.AssertExpectations(t)
	assert.Equal(t, expect, actual, "they should be equal output")
}

// 異常系
// ゲームマスター取得失敗
func TestExecute_getGM_fail(t *testing.T) {
	// ダミーデータ
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
		Err:           err,
	}

	// 参照用リポジトリMock
	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", dummyInput.GroupRoomID).Return(nil, err)

	// プレゼンターMock
	presenterMock := new(mocks.Presenter)
	presenterMock.On("Execute", dummyOutput).Return(nil)
	joinUseCase := UseCaseImpl{
		readOnlyRepository: readOnlyRepositoryMock,
		presenter:          presenterMock,
	}

	// 実行
	expect := dummyOutput
	actual := joinUseCase.Excute(dummyInput)

	// 検証
	readOnlyRepositoryMock.AssertExpectations(t)
	presenterMock.AssertExpectations(t)
	assert.Equal(t, expect, actual, "they should be equal output")
}

// 異常系
// 既にトーク中
func TestExecute_alreadyStart_fail(t *testing.T) {
	// ダミーデータ
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

	// 参照用リポジトリMock
	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", dummyInput.GroupRoomID).Return(dummyGameMaster, nil)

	// プレゼンターMock
	presenterMock := new(mocks.Presenter)
	presenterMock.On("Execute", dummyOutput).Return(nil)
	joinUseCase := UseCaseImpl{
		readOnlyRepository: readOnlyRepositoryMock,
		presenter:          presenterMock,
	}

	// 実行
	expect := dummyOutput
	actual := joinUseCase.Excute(dummyInput)

	// 検証
	readOnlyRepositoryMock.AssertExpectations(t)
	presenterMock.AssertExpectations(t)
	assert.Equal(t, expect, actual, "they should be equal output")
}

// 正常系
// 終了タイマー関数の確認
func TestSetFinishTimer(t *testing.T) {
	// ダミーデータ
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

	// 参照用リポジトリMock
	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)

	// プレゼンターMock
	presenterMock := new(mocks.Presenter)
	presenterMock.On("FinishTalk", dummyOutput).Return(nil)
	joinUseCase := UseCaseImpl{
		readOnlyRepository: readOnlyRepositoryMock,
		presenter:          presenterMock,
	}

	// 実行
	setFinishTimer(&joinUseCase, dummyInput, time.Second*1)

	// 検証
	presenterMock.AssertExpectations(t)
}
