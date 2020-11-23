package domain

import (
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	println("[test start]")
	m.Run()
	println("[test finish]")
}

// 現在時刻 < 終了時刻のとき、正の値を返す
func TestRemainingTime_withNowBeforeEndTime_positive(t *testing.T) {
	// 事前準備
	g := NewGameMaster("test123")
	endTime := time.Date(2020, 6, 1, 17, 00, 13, 0, time.Local)
	nowTime := time.Date(2020, 6, 1, 16, 00, 13, 0, time.Local)
	// 条件_現在時刻 < 終了時刻
	if !nowTime.Before(endTime) {
		t.Error("条件エラー")
	}
	g.endTime = endTime
	// 実行
	result := g.RemainingTime(nowTime)
	expect, _ := time.ParseDuration("0ms")
	// 検証_正の値
	if result < expect {
		t.Fail()
	}
	// 後処理
	t.Log("\n期待値： > ", expect, "\n実際値： ", result)
	t.Log("TestRemainingTime")
}

// 現在時刻 > 終了時刻のとき、負の値を返す
func TestRemainingTime_withNowBeforeEndTime_negative(t *testing.T) {
	// 事前準備
	g := NewGameMaster("test123")
	endTime := time.Date(2020, 6, 1, 17, 00, 13, 0, time.Local)
	nowTime := time.Date(2020, 6, 1, 18, 30, 13, 0, time.Local)
	// 条件_現在時刻 < 終了時刻
	if !nowTime.After(endTime) {
		t.Error("条件エラー")
	}
	g.endTime = endTime
	// 実行
	result := g.RemainingTime(nowTime)
	expect, _ := time.ParseDuration("0ms")
	// 検証_負の値
	if result > expect {
		t.Fail()
	}
	// 後処理
	t.Log("\n期待値： > ", expect, "\n実際値： ", result)
	t.Log("TestRemainingTime")
}
