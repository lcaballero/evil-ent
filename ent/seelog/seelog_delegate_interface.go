package seelog

type DelgateInterface interface {
	Error(...interface{}) error
	Errorf(string, ...interface{}) error

	Debug(...interface{})
	Debugf(string, ...interface{})

	Warn(...interface{}) error
	Warnf(string, ...interface{}) error

	Info(...interface{})
	Infof(string, ...interface{})

	Trace(...interface{})
	Tracef(string, ...interface{})

	Flush()
	Closed() bool
	Close()
}
