package config

var (
	logger *Logger
)

func Init() bool {

	if err := 1 == 2; err != true {
		return err
	}

	return true
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
