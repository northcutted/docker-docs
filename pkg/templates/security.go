package templates

import (
	"bytes"
	"context"
	"fmt"
	"text/template"
	"time"
)

// SecurityConfig defines security limits for template execution.
type SecurityConfig struct {
	// MaxExecutionTime is the maximum duration a template can execute.
	MaxExecutionTime time.Duration
	// MaxOutputSize is the maximum number of bytes the template output can produce.
	MaxOutputSize int64
}

// DefaultSecurityConfig returns the default security configuration.
func DefaultSecurityConfig() SecurityConfig {
	return SecurityConfig{
		MaxExecutionTime: 30 * time.Second,
		MaxOutputSize:    10 * 1024 * 1024, // 10MB
	}
}

// ExecuteWithLimits runs a template with timeout and output size limits.
func ExecuteWithLimits(tmpl *template.Template, data interface{}, sec SecurityConfig) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), sec.MaxExecutionTime)
	defer cancel()

	type result struct {
		output string
		err    error
	}

	resultChan := make(chan result, 1)

	go func() {
		buf := &limitedWriter{limit: sec.MaxOutputSize}
		err := tmpl.Execute(buf, data)
		resultChan <- result{output: buf.String(), err: err}
	}()

	select {
	case res := <-resultChan:
		if res.err != nil {
			return "", fmt.Errorf("template execution failed: %w", res.err)
		}
		return res.output, nil
	case <-ctx.Done():
		return "", fmt.Errorf("template execution timeout exceeded (%s)", sec.MaxExecutionTime)
	}
}

// limitedWriter wraps a bytes.Buffer with a size limit.
type limitedWriter struct {
	buf   bytes.Buffer
	limit int64
}

func (w *limitedWriter) Write(p []byte) (n int, err error) {
	if int64(w.buf.Len()+len(p)) > w.limit {
		return 0, fmt.Errorf("template output size exceeded limit (%d bytes)", w.limit)
	}
	return w.buf.Write(p)
}

func (w *limitedWriter) String() string {
	return w.buf.String()
}
