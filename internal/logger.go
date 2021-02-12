package internal

type logger struct {
	Msgs  string
	Stack []int
}

var log = &logger{}

func (l *logger) Add(elem ...interface{}) {
	for _, value := range elem {
		switch value := value.(type) {
		case string:
			l.Msgs += value + "\n"
		case Stack:
			l.Stack = value
		}
	}
}
