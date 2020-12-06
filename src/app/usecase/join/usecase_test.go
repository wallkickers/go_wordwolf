package join

import (
	"github.com/go-server-dev/src/app/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type readOnlyRepositoryMock struct {
	mock.Mock
}

func (m *readOnlyRepositoryMock) FindMemberByID(memberID string) (*domain.Member, error) {
	args := m.Called(memberID)
	return args.Get(0).(*domain.Member), args.Error(1)
}

func (m *readOnlyRepositoryMock) FindGameMasterByGroupID(groupID string) (*domain.GameMaster, error) {
	args := m.Called(groupID)
	return args.Get(0).(*domain.GameMaster), args.Error(1)
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

func (m *presenterMock) Execute(output Output) {
	m.Called(output)
}

// 正常系
func TestExecute_withInput_success(t *testing.T) {
	// ダミーデータ
	dummyInput := Input{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "testGroup",
	}
	dummyMember := domain.NewMember(dummyInput.MemberID, "testMemberName_Taro")
	dummyGameMaster := domain.NewGameMaster(dummyInput.GroupRoomID, domain.GroupRoomType(dummyInput.GroupRoomType))
	dummyGameMaster.SetMember(dummyInput.MemberID)
	dummyOutput := Output{
		ReplyToken:    "testReplyToken",
		MemberID:      "testMemberID12345",
		MemberName:    dummyMember.Name(),
		GroupRoomID:   "testGroup12345",
		GroupRoomType: "testGroup",
	}
	// 参照用リポジトリMock
	readOnlyRepositoryMock := new(readOnlyRepositoryMock)
	readOnlyRepositoryMock.On("FindMemberByID", dummyInput.MemberID).Return(dummyMember, nil)
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", dummyInput.GroupRoomID).Return(dummyGameMaster, nil)
	// 更新用リポジトリMock
	gameMasterRepositoryMock := new(gameMasterRepositoryMock)
	gameMasterRepositoryMock.On("Save", dummyGameMaster).Return(nil)
	// プレゼンターMock
	presenterMock := new(presenterMock)
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
