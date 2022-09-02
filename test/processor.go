package splitbatchprocessor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var _ component.Processor = (*splitBatch)(nil)

type splitBatch struct {
	next consumer.Traces
}

func newSplitBatch(next consumer.Traces) *splitBatch {
	return &splitBatch{next: next}
}

func (s *splitBatch) ConsumeTraces(ctx context.Context, batch ptrace.Traces) error {
	return nil
}

func (s *splitBatch) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: true}
}

func (s *splitBatch) Start(_ context.Context, host component.Host) error {
	return nil
}

func (s *splitBatch) Shutdown(context.Context) error {
	return nil
}
