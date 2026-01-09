package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

type OutputMode struct {
	JSON  bool
	Plain bool
}

func printJSON(w io.Writer, data any) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}

func printJSONBytes(w io.Writer, data []byte) error {
	if len(data) == 0 {
		return nil
	}
	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("invalid JSON response: %w", err)
	}
	return printJSON(w, v)
}

func printTable(w io.Writer, headers []string, rows [][]string, plain bool) {
	tw := tabwriter.NewWriter(w, 0, 4, 2, ' ', 0)
	if !plain && len(headers) > 0 {
		fmt.Fprintln(tw, joinRow(headers))
	}
	for _, row := range rows {
		fmt.Fprintln(tw, joinRow(row))
	}
	_ = tw.Flush()
}

func joinRow(cols []string) string {
	return strings.Join(cols, "\t")
}
