package client

import (
	"context"
	"errors"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	defaultDurationDistribution = view.Distribution(0, 1, 2, 5, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000, 20000)

	measureDuration = stats.Float64("delegated_routing/duration", "The time to complete an entire request", stats.UnitMilliseconds)
	measureRequests = stats.Float64("delegated_routing/requests", "The number of requests made", stats.UnitDimensionless)

	keyName  = tag.MustNewKey("name")
	keyError = tag.MustNewKey("error")

	durationView = &view.View{
		Measure:     measureDuration,
		TagKeys:     []tag.Key{keyName, keyError},
		Aggregation: defaultDurationDistribution,
	}
	requestsView = &view.View{
		Measure:     measureRequests,
		TagKeys:     []tag.Key{keyName, keyError},
		Aggregation: view.Sum(),
	}

	DefaultViews = []*view.View{
		durationView,
		requestsView,
	}
)

// startMetrics begins recording metrics.
// The returned function flushes the metrics when called, recording metrics about the passed error.
func startMetrics(ctx context.Context, name string) (done func(err error)) {
	start := time.Now()

	return func(err error) {
		latency := time.Since(start)

		errStr := "None"
		if err != nil {
			logger.Warnw("received delegated routing error", "Error", err)
			if errors.Is(err, context.Canceled) {
				errStr = "Canceled"
			} else if errors.Is(err, context.DeadlineExceeded) {
				errStr = "DeadlineExceeded"
			} else {
				errStr = "Unknown"
			}
		}

		stats.RecordWithTags(ctx,
			[]tag.Mutator{
				tag.Upsert(keyName, name),
				tag.Upsert(keyError, errStr),
			},
			measureDuration.M(float64(latency.Milliseconds())),
			measureRequests.M(1),
		)
	}
}
