package cyb

import "encoding/json"

func parseJson(data []byte, v any) (err error) {
	defer func() {
		if err == nil {
			err = validate(v)
		}
	}()

	return json.Unmarshal(data, v)
}

func toJson(v any) (b []byte, err error) {
	if err = validate(v); err != nil {
		return
	}

	return json.Marshal(v)
}
