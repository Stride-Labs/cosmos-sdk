package log

type NoOp struct{}

func NewNopLogger() Logger {
	return &NoOp{}
}

func (l NoOp) Debug(msg string, keyvals ...interface{}) {}
func (l NoOp) Info(msg string, keyvals ...interface{})  {}
func (l NoOp) Error(msg string, keyvals ...interface{}) {}
