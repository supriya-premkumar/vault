package audit

import (
	"encoding/json"
	"fmt"
	"io"
)

// JSONFormatWriter is an AuditFormatWriter implementation that structures data into
// a JSON format.
type JSONFormatWriter struct {
	Prefix string
}

func (f *JSONFormatWriter) WriteRequest(w io.Writer, req *AuditRequestEntry) error {
	if req == nil {
		return fmt.Errorf("request entry was nil, cannot encode")
	}

	if len(f.Prefix) > 0 {
		_, err := w.Write([]byte(f.Prefix))
		if err != nil {
			return err
		}
	}

	enc := json.NewEncoder(w)
	return enc.Encode(req)
}

func (f *JSONFormatWriter) WriteResponse(w io.Writer, resp *AuditResponseEntry) error {
	if resp == nil {
		return fmt.Errorf("response entry was nil, cannot encode")
	}

	if len(f.Prefix) > 0 {
		_, err := w.Write([]byte(f.Prefix))
		if err != nil {
			return err
		}
	}

	enc := json.NewEncoder(w)
	return enc.Encode(resp)
}
