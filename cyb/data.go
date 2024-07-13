package cyb

type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Content any    `json:"content"`
}

type ChannelData struct {
	UUID    string `json:"uuid"`
	Method  string `json:"method"`
	Channel string `json:"channel"`
}
