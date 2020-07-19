package barion

import (
	"fmt"
	"time"

	"github.com/leekchan/timeutil"
)

type TimeSpan time.Duration

func (t TimeSpan) MarshalText() (text []byte, err error) {
	tt := time.Time{}
	tt = tt.Add(time.Duration(t))
	days := fmt.Sprint(int(time.Duration(t).Hours() / 24))
	fmt.Println(days)
	timestamp := tt.Format("15:04:05")
	return []byte(days + "." + timestamp), nil
}

func (t *TimeSpan) UnmarshalText(text []byte) error {
	var days time.Duration
	var hours time.Duration
	var minutes time.Duration
	var seconds time.Duration
	_, err := fmt.Sscanf(string(text), "%d.%d:%d:%d", &days, &hours, &minutes, &seconds)
	if err != nil {
		return err
	}
	td := timeutil.Timedelta{Days: days, Hours: hours, Minutes: minutes, Seconds: seconds}
	*t = TimeSpan(td.Duration())
	return nil
}
