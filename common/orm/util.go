package orm

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type LocalTime struct {
	time.Time
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	//格式化秒
	tm := t.Time.Format("2006-01-02 15:04:05")
	var zeroTime time.Time
	if tm == zeroTime.Format("2006-01-02 15:04:05") {
		tm = ""
	}
	return json.Marshal(tm)
}

func (t *LocalTime) UnmarshalJSON(s []byte) (err error) {
	if string(s) == "null" || string(s) == `""` {
		return nil
	}
	t.Time, err = time.ParseInLocation("\"2006-01-02 15:04:05\"", string(s), time.Local)
	return
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	//支持插入空时间
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
