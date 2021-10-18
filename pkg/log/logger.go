package log

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"runtime"
)

//var Logger io.Writer
//func init()  {
//	var err error
//	Logger,err = os.Create("./app.log")
//	log.Println(err)
//}

type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "[DEBUG]"
	case LevelInfo:
		return "[INFO]"
	case LevelWarn:
		return "[WARN]"
	case LevelError:
		return "[ERROR]"

	case LevelPanic:
		return "[PANIC]"
	}
	return ""
}

type Logger struct {
	newLogger *log.Logger

	fields  Fields
	callers []string
}

func NewLogger(w io.Writer) *Logger {
	l := log.New(w, "", log.LstdFlags)
	return &Logger{newLogger: l}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}
	for k, v := range f {
		ll.fields[k] = v
	}
	return ll
}
func (l *Logger) WithWriter(w io.Writer) *Logger {
	ll := l.clone()
	ll.newLogger = log.New(w, "", log.LstdFlags)
	return ll
}

func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}

	return ll
}

func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 3
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf(` %s: %d %s `, frame.File, frame.Line, frame.Function)

		callers = append(callers, s)
		if !more {
			break
		}
	}
	ll := l.clone()
	ll.callers = callers
	return ll
}

func (l *Logger) WithTrace() *Logger {
	return l.WithFields(Fields{
		"trace_id": "X-Trace-ID",
		"span_id":  "X-Span-ID",
	})
}

func (l *Logger) JSONFormat() map[string]interface{} {
	data := make(Fields, len(l.fields)+4)

	if len(l.callers) > 0 {
		data["callers"] = l.callers
	}

	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}

	return data
}

func (l *Logger) Output(level Level) {
	body, _ := json.Marshal(l.JSONFormat())
	content := level.String() + string(body)

	switch level {
	case LevelDebug, LevelInfo, LevelWarn, LevelError:
		l.newLogger.Println(fmt.Sprint(content))
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

func (l *Logger) Debug(key string, v ...interface{}) {
	fields := make(Fields, 1)
	fields[key] = v
	l.WithFields(fields).Output(LevelDebug)

}

func (l *Logger) Debugf(format string, v ...interface{}) {
	fields := make(Fields, 1)
	fields[format] = fmt.Sprintf(format, v...)
	l.WithFields(fields).Output(LevelDebug)
}

//func (l *Logger) Info( prefix string,v ...interface{}) {
//	l.Output(LevelInfo, ,fmt.Sprint(v...))
//}
//
//func (l *Logger) Infof( format string, v ...interface{}) {
//	l.Output(LevelInfo, fmt.Sprintf(format, v...))
//}
//
//func (l *Logger) Warn( v ...interface{}) {
//	l.Output(LevelWarn, fmt.Sprint(v...))
//}
//
//func (l *Logger) Warnf( format string, v ...interface{}) {
//	l.Output(LevelWarn, fmt.Sprintf(format, v...))
//}
//
func (l *Logger) Error(key string, v ...interface{}) {
	fields := make(Fields, 1)
	fields[key] = v
	l.WithFields(fields).WithCallersFrames().Output(LevelError)

}

//
func (l *Logger) Errorf(format string, v ...interface{}) {
	fields := make(Fields, 1)
	fields[format] = fmt.Sprintf(format, v)

	l.WithFields(fields).WithCallersFrames().Output(LevelError)
}

//
//
//func (l *Logger) Panic( v ...interface{}) {
//	l.Output(LevelPanic, fmt.Sprint(v...))
//}
//
//func (l *Logger) Panicf( format string, v ...interface{}) {
//	l.Output(LevelPanic, fmt.Sprintf(format, v...))
//}
