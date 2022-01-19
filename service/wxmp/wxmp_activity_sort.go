package wxmp

import (
    "RemindMe/model/wxmp"
    "sort"
)

// Bucket 定义一个通用的结构体
type Bucket struct {
    Slice []wxmp.Activity             //承载以任意结构体为元素构成的Slice
    By    func(a, b interface{}) bool //排序规则函数,当需要对新的结构体slice进行排序时，只需定义这个函数即可
}

func (this Bucket) Len() int { return len(this.Slice) }

func (this Bucket) Swap(i, j int) { this.Slice[i], this.Slice[j] = this.Slice[j], this.Slice[i] }

func (this Bucket) Less(i, j int) bool { return this.By(this.Slice[i], this.Slice[j]) }

func sortActivitiesByTime(list []wxmp.Activity) {
    f := func(a, b interface{}) bool {
        t1, t2 := a.(wxmp.Activity).Time.Time, b.(wxmp.Activity).Time.Time
        return t1.Before(t2) || (t1.Equal(t2) && a.(wxmp.Activity).ID < b.(wxmp.Activity).ID)
    }
    results := Bucket{By: f, Slice: list}
    sort.Sort(results)
}