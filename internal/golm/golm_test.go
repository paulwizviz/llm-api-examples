package golm

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type bytesReadCloser struct {
	*bytes.Reader
}

func (b bytesReadCloser) Close() error {
	return nil // No-op close
}

func TestParseStream(t *testing.T) {
	testcases := []struct {
		input string
		want  int
	}{
		{
			input: `data: {"id":"cmpl-woqeinll1jb0wdppr152d","object":"text_completion","created":1741514769,"model":"gemma-2-2b-it","choices":[{"index":0,"text":"<","logprobs":null,"finish_reason":null}]}
data: {"id":"cmpl-woqeinll1jb0wdppr152d","object":"text_completion","created":1741514769,"model":"gemma-2-2b-it","choices":[{"index":0,"text":"ul","logprobs":null,"finish_reason":null}]}
data: {"id":"cmpl-eegcovqioujw48xwjo3me","object":"text_completion","created":1741473481,"model":"gemma-2-2b-it","choices":[{"index":0,"text":"ul","logprobs":null,"finish_reason":null}]}
data: {"id":"cmpl-eegcovqioujw48xwjo3me","object":"text_completion","created":1741473481,"model":"gemma-2-2b-it","choices":[{"index":0,"text":">","logprobs":null,"finish_reason":null}]}
data: {"id":"cmpl-eegcovqioujw48xwjo3me","object":"text_completion","created":1741473481,"model":"gemma-2-2b-it","choices":[{"index":0,"text":"\n","logprobs":null,"finish_reason":null}]}
`,
			want: 5,
		},
		{
			input: `data: {"id":"cmpl-eegcovqioujw48xwjo3me","object":"text_completion","created":1741473481,"model":"gemma-2-2b-it","choices":[{"index":0,"text":"\n\n","logprobs":null,"finish_reason":null}]}
data: {"id":"cmpl-eegcovqioujw48xwjo3me","object":"text_completion","created":1741473481,"model":"gemma-2-2b-it","choices":[{"index":0,"text":"<","logprobs":null,"finish_reason":null}]}
data: {"id":"cmpl-eegcovqioujw48xwjo3me","object":"text_completion","created":1741473481,"model":"gemma-2-2b-it","choices":[{"index":0,"text":"ul","logprobs":null,"finish_reason":null}]}
data: {"id":"cmpl-eegcovqioujw48xwjo3me","object":"text_completion","created":1741473481,"model":"gemma-2-2b-it","choices":[{"index":0,"text":">","logprobs":null,"finish_reason":null}]}
`,
			want: 4,
		},
	}

	for i, tc := range testcases {
		bReader := bytes.NewReader([]byte(tc.input))
		bR := bytesReadCloser{bReader}
		got := parseGenerateStream(bR)
		assert.Equal(t, tc.want, len(got), fmt.Sprintf("Case: %d Want: %v Got: %v", i, tc.want, len(got)))

	}
}
