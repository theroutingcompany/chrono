package chrono

import "time"

func (t *Time) Add(p1 time.Duration) time.Time {
	return t.T().Add(p1)
}
func (t *Time) AddDate(p1 int, p2 int, p3 int) time.Time {
	return t.T().AddDate(p1, p2, p3)
}
func (t *Time) After(p1 time.Time) bool {
	return t.T().After(p1)
}
func (t *Time) AppendFormat(p1 []uint8, p2 string) []uint8 {
	return t.T().AppendFormat(p1, p2)
}
func (t *Time) Before(p1 time.Time) bool {
	return t.T().Before(p1)
}
func (t *Time) Clock() (int, int, int) {
	return t.T().Clock()
}
func (t *Time) Date() (int, time.Month, int) {
	return t.T().Date()
}
func (t *Time) Day() int {
	return t.T().Day()
}
func (t *Time) Equal(p1 time.Time) bool {
	return t.T().Equal(p1)
}
func (t *Time) Format(p1 string) string {
	return t.T().Format(p1)
}
func (t *Time) GobEncode() ([]uint8, error) {
	return t.T().GobEncode()
}
func (t *Time) Hour() int {
	return t.T().Hour()
}
func (t *Time) ISOWeek() (int, int) {
	return t.T().ISOWeek()
}
func (t *Time) In(p1 *time.Location) time.Time {
	return t.T().In(p1)
}
func (t *Time) IsZero() bool {
	return t.T().IsZero()
}
func (t *Time) Local() time.Time {
	return t.T().Local()
}
func (t *Time) Location() *time.Location {
	return t.T().Location()
}
func (t *Time) MarshalBinary() ([]uint8, error) {
	return t.T().MarshalBinary()
}
func (t *Time) MarshalText() ([]uint8, error) {
	return t.T().MarshalText()
}
func (t *Time) Minute() int {
	return t.T().Minute()
}
func (t *Time) Month() time.Month {
	return t.T().Month()
}
func (t *Time) Nanosecond() int {
	return t.T().Nanosecond()
}
func (t *Time) Round(p1 time.Duration) time.Time {
	return t.T().Round(p1)
}
func (t *Time) Second() int {
	return t.T().Second()
}
func (t *Time) Sub(p1 time.Time) time.Duration {
	return t.T().Sub(p1)
}
func (t *Time) Truncate(p1 time.Duration) time.Time {
	return t.T().Truncate(p1)
}
func (t *Time) UTC() time.Time {
	return t.T().UTC()
}
func (t *Time) Unix() int64 {
	return t.T().Unix()
}
func (t *Time) UnixNano() int64 {
	return t.T().UnixNano()
}
func (t *Time) Weekday() time.Weekday {
	return t.T().Weekday()
}
func (t *Time) Year() int {
	return t.T().Year()
}
func (t *Time) YearDay() int {
	return t.T().YearDay()
}
func (t *Time) Zone() (string, int) {
	return t.T().Zone()
}
