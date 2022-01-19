package wxmp

import (
    "testing"
    "time"
)

func TestActivityExtraService_CreateActivityExtraInfo(t *testing.T) {
    tim := time.Date(2021, 2, 11, 10, 0, 0, 0, time.Local)
    for i := -5; i <= 5; i++ {
        t.Log(periodicFuncs[5](tim, i))
    }
}
