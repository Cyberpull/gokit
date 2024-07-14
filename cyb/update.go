package cyb

type Update struct {
	Data
	ChannelData
}

func (x Update) name() string {
	return "update"
}

func (x Update) prefix() string {
	return "UPDATE::"
}
