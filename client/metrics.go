package client

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/ipld/edelweiss/services"
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
			errStr = metricsErrStr(err)
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

// metricsErrStr returns a string to use for recording metrics from an error.
// We shouldn't use the error string itself as that can result in high-cardinality metrics.
// For more specific root causing, check the logs.
func metricsErrStr(err error) string {
	if errors.Is(err, context.DeadlineExceeded) {
		return "DeadlineExceeded"
	}
	if errors.Is(err, context.Canceled) {
		return "Canceled"
	}
	if errors.Is(err, services.ErrSchema) {
		return "Schema"
	}

	var serviceErr *services.ErrService
	if errors.As(err, &serviceErr) {
		return "Service"
	}

	var protoErr *services.ErrProto
	if errors.As(err, &protoErr) {
		return "Proto"
	}

	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) {
		if dnsErr.IsNotFound {
			return "DNSNotFound"
		}
		if dnsErr.IsTimeout {
			return "DNSTimeout"
		}
		return "DNS"
	}

	var netErr net.Error
	if errors.As(err, &netErr) {
		if netErr.Timeout() {
			return "NetTimeout"
		}
		return "Net"
	}

	return "Other"
}
