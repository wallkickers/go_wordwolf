package domain

import (
	"testing"
	"time"
)

//残り時間表示（終了時間-現在時間）

//トークスタートと同時に終了タイマーを作成、一定時間後に告知表示
//トークスタートと同時に終了タイマー・残り時間タイマーセット・一定時間後に告知表示
//トーク時間追加
//	1.ステート状態確認
//		トーク前の場合→「トークが始まっていません」メッセージ
//		トーク後の場合→「時間切れ」メッセージ
//		トーク中の場合→　終了時間を延長＋タイマーをリセット・再度タイマーをセット
func TestSingleton(t *testing.T) {
	t1, t2 := GetTimerManeger(), GetTimerManeger()
	t1.talkTimeMin = time.Microsecond * 2
	t2.talkTimeMin = time.Microsecond * 3
	println("t1", t1)
	println("t2", t2)
	if !(t1 == t2) {
		t.Fail()
	}
	//後処理
	shared = nil
}

// 非同期でアクセス時のシングルトンのテスト
func TestThreadSeafSingleton(t *testing.T) {
	ch := make(chan interface{})
	go func() {
		t1 := GetTimerManeger()
		ch <- t1
	}()
	go func() {
		t2 := GetTimerManeger()
		ch <- t2
	}()
	t1, t2 := <-ch, <-ch
	println("t1", t1)
	println("t2", t2)
	if !(t1 == t2) {
		t.Fail()
	}
	//後処理
	shared = nil
}

func TestSetTimer(t *testing.T) {
	//前処理
	g := GetTimerManeger()
	print(g)
	//now := time.Now() //現在時刻
	//実行
	//g.SetTimer(now)
	// //検証
	// if !(g.endTime == now.Add(g.talkTimeMin)) {
	// 	t.Fail()
	// }
}
