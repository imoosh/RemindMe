package wxmp

import (
    "errors"
    "fmt"
    "strconv"
    "strings"
    "time"
)

type activityCursor struct {
    ts    int64
    id    uint
    subId uint
}

func parseCursor(cursor string) (*activityCursor, error) {
    if cursor == "" {
        return &activityCursor{ts: time.Now().Unix(), id: 0, subId: 0}, nil
    }

    var ss = strings.Split(cursor, ".")
    if len(ss) != 3 {
        return nil, errors.New("非法的分页游标")
    }
    var (
        ts, _    = strconv.ParseInt(ss[0], 10, 64)
        id, _    = strconv.ParseInt(ss[1], 10, 64)
        subId, _ = strconv.ParseInt(ss[2], 10, 64)
    )
    return &activityCursor{ts: ts, id: uint(id), subId: uint(subId)}, nil
}

func newCursor(ts int64, id, subId uint) *activityCursor {
    return &activityCursor{ts: ts, id: id, subId: subId}
}

func (c *activityCursor) String() string {
    return fmt.Sprintf("%v.%v.%v", c.ts, c.id, c.subId)
}
