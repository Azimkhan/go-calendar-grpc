package errors

type EventError string

func (ee EventError) Error() string {
	return string(ee)
}

var (
	EventNotFound = EventError("event not found")
)
