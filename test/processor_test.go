package splitbatchprocessor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

func TestSplitDifferentTracesIntoDifferentBatches(t *testing.T) {
	// we have 1 ResourceSpans with 1 ILS and two traceIDs, resulting in two batches
	inBatch := ptrace.NewTraces()
	rs := inBatch.ResourceSpans().AppendEmpty()

	// the first ILS has two spans
	ils := rs.ScopeSpans().AppendEmpty()
	library := ils.Scope()
	library.SetName("first-library")
	firstSpan := ils.Spans().AppendEmpty()
	firstSpan.SetName("first-batch-first-span")
	firstSpan.SetTraceID(pcommon.NewTraceID([16]byte{1, 2, 3, 4}))
	secondSpan := ils.Spans().AppendEmpty()
	secondSpan.SetName("first-batch-second-span")
	secondSpan.SetTraceID(pcommon.NewTraceID([16]byte{2, 3, 4, 5}))

	// test
	out := SplitTraces(inBatch)

	// verify
	assert.Len(t, out, 2)

	// first batch
	firstOutILS := out[0].ResourceSpans().At(0).ScopeSpans().At(0)
	assert.Equal(t, library.Name(), firstOutILS.Scope().Name())
	assert.Equal(t, firstSpan.Name(), firstOutILS.Spans().At(0).Name())

	// second batch
	secondOutILS := out[1].ResourceSpans().At(0).ScopeSpans().At(0)
	assert.Equal(t, library.Name(), secondOutILS.Scope().Name())
	assert.Equal(t, secondSpan.Name(), secondOutILS.Spans().At(0).Name())
}
