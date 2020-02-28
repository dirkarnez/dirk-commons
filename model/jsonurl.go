package model

import "net/url"

// JSONURL for parsing only
type JSONURL struct {
	*url.URL
}

// UnmarshalJSON Unmarshal JSON
func (j *JSONURL) UnmarshalJSON(b []byte) error {
	u, err := url.Parse(string(b[1 : len(b)-1]))
	if err != nil {
		return err
	}
	j.URL = u
	return nil
}
