package seelog

import "fmt"

type MockLogger struct {
	Writes    map[string][]string
	IsClosed  bool
	IsFlushed bool
}

func NewMockLogger() *MockLogger {
	m := &MockLogger{
		Writes:   make(map[string][]string),
		IsClosed: false,
	}
	return m
}

func (m *MockLogger) Add(key, val string) {
	entries, ok := m.Writes[key]
	if ok {
		m.Writes[key] = append(entries, val)
	} else {
		m.Writes[key] = []string{val}
	}
}

func (m *MockLogger) Error(v ...interface{}) error {
	m.Add("Error", fmt.Sprintln(v...))
	return nil
}

func (m *MockLogger) Errorf(format string, v ...interface{}) error {
	m.Add("Errorf", fmt.Sprintf(format, v...))
	return nil
}

func (m *MockLogger) Debug(v ...interface{}) {
	m.Add("Debug", fmt.Sprintln(v...))
}

func (m *MockLogger) Debugf(format string, v ...interface{}) {
	m.Add("Debugf", fmt.Sprintf(format, v...))
}

func (m *MockLogger) Warn(v ...interface{}) error {
	m.Add("Warn", fmt.Sprintln(v...))
	return nil
}
func (m *MockLogger) Warnf(format string, v ...interface{}) error {
	m.Add("Warnf", fmt.Sprintf(format, v...))
	return nil
}

func (m *MockLogger) Info(v ...interface{}) {
	m.Add("Info", fmt.Sprintln(v...))

}
func (m *MockLogger) Infof(format string, v ...interface{}) {
	m.Add("Infof", fmt.Sprintf(format, v...))
}

func (m *MockLogger) Trace(v ...interface{}) {
	m.Add("Trace", fmt.Sprintln(v...))
}

func (m *MockLogger) Tracef(format string, v ...interface{}) {
	m.Add("Tracef", fmt.Sprintf(format, v...))
}

func (m *MockLogger) Fatal(v ...interface{}) {
	m.Add("Fatal", fmt.Sprintln(v...))
}

func (m *MockLogger) Fatalf(format string, v ...interface{}) {
	m.Add("Fatalf", fmt.Sprintf(format, v...))
}

func (m *MockLogger) Flush() {
	m.IsFlushed = true
}

func (m *MockLogger) Closed() bool {
	return m.IsClosed
}

func (m *MockLogger) Close() {
	m.IsClosed = true
}
