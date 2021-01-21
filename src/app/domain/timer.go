package domain

import (
	"sync"
	"time"
)

var shared *TimerManager
var once sync.Once

// TimerManager 時間管理マネージャ
type TimerManager struct {
	timer       *time.Timer
	endTime     time.Time
	talkTimeMin time.Duration
}

// GetTimerManeger 時間管理インスタンス取得
func GetTimerManeger() *TimerManager {
	once.Do(func() { //スレッドセーフ
		shared = &TimerManager{
			talkTimeMin: time.Second * 3, //3分間
		}
	})
	return shared
}

func (t *TimerManager) SetTimer(now time.Time) {
	t.timer = time.NewTimer(t.talkTimeMin)
	t.endTime = now.Add(t.talkTimeMin)
}
