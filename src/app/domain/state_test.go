package domain

import (
	"testing"
)

// NewGameMaster ゲームマスター初期化時にStateにSettingTimeをセットする
func TestNewGameMaster_setState(t *testing.T) {
	gm := NewGameMaster("testRoom", Group)
	result := gm.state
	expect := settingTime
	if result != expect {
		t.Fail()
	}
}

// IsTalkTIme トーク中の場合、Trueを返す
func TestIsTalkTime_returnTrue(t *testing.T) {
	gm := NewGameMaster("testRoom", Group)
	gm.StartTalk()
	result := gm.IsTalkTime()
	if result != true {
		t.Fail()
	}
}

// IsTalkTimeトーク中でない場合、Falseを返す
func TestIsTalkTime_returnFalse(t *testing.T) {
	gm := NewGameMaster("testRoom", Group)
	gm.EndTalk()
	result := gm.IsTalkTime()
	if result != false {
		t.Fail()
	}
}
