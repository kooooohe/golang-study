package main

/**
type error interface{
	Error() string
}

**/
type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}
func RaiseError() error {
	return &MyError{Message: "えらーがおきた", ErrCode: 1234}
}

func main() {

}

func tmp() (err error) {
	return
}
