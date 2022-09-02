package splitbatchprocessor

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// SplitTraces returns one ptrace.Traces for each trace in the given ptrace.Traces input. Each of the resulting ptrace.Traces contains exactly one trace.
func SplitTraces(batch ptrace.Traces) []ptrace.Traces {
	// for each span in the resource spans, we group them into batches of rs/ils/traceID.
	// if the same traceID exists in different ils, they land in different batches.
	var result []ptrace.Traces

	for i := 0; i < batch.ResourceSpans().Len(); i++ {
		rs := batch.ResourceSpans().At(i)

		for j := 0; j < rs.ScopeSpans().Len(); j++ {
			// the batches for this ILS
			batches := map[pcommon.TraceID]ptrace.ResourceSpans{}

			ils := rs.ScopeSpans().At(j)
			for k := 0; k < ils.Spans().Len(); k++ {
				span := ils.Spans().At(k)
				key := span.TraceID()

				// for the first traceID in the ILS, initialize the map entry
				// and add the singleTraceBatch to the result list
				if _, ok := batches[key]; !ok {
					trace := ptrace.NewTraces()
					newRS := trace.ResourceSpans().AppendEmpty()
					// currently, the ResourceSpans implementation has only a Resource and an ILS. We'll copy the Resource
					// and set our own ILS
					rs.Resource().CopyTo(newRS.Resource())

					newILS := newRS.ScopeSpans().AppendEmpty()
					// currently, the ILS implementation has only an InstrumentationLibrary and spans. We'll copy the library
					// and set our own spans
					ils.Scope().CopyTo(newILS.Scope())
					batches[key] = newRS

					result = append(result, trace)
				}

				// there is only one instrumentation library per batch
				tgt := batches[key].ScopeSpans().At(0).Spans().AppendEmpty()
				span.CopyTo(tgt)
			}
		}
	}

	return result
}
