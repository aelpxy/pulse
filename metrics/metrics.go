package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// connection metrics
	ConnectionsActive = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pulse_connections_active",
		Help: "Number of active WebSocket connections",
	}, []string{"app_key"})

	ConnectionsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "pulse_connections_total",
		Help: "Total number of WebSocket connections established",
	}, []string{"app_key"})

	ConnectionsRejected = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "pulse_connections_rejected_total",
		Help: "Total number of rejected connections",
	}, []string{"app_key", "reason"})

	// channel metrics
	ChannelsActive = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "pulse_channels_active",
		Help: "Number of active channels",
	})

	ChannelSubscriptions = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "pulse_channel_subscriptions_total",
		Help: "Total number of channel subscriptions",
	}, []string{"app_key", "channel_type"})

	// message metrics
	MessagesPublished = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "pulse_messages_published_total",
		Help: "Total number of messages published",
	}, []string{"app_key", "event_type"})

	MessagesSent = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "pulse_messages_sent_total",
		Help: "Total number of messages sent to clients",
	}, []string{"app_key"})

	MessageErrors = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "pulse_message_errors_total",
		Help: "Total number of message errors",
	}, []string{"app_key", "error_type"})

	// HTTP API metrics
	HTTPRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "pulse_http_requests_total",
		Help: "Total number of HTTP requests",
	}, []string{"endpoint", "method", "status"})

	HTTPRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "pulse_http_request_duration_seconds",
		Help:    "HTTP request latencies in seconds",
		Buckets: prometheus.DefBuckets,
	}, []string{"endpoint", "method"})

	// app metrics
	AppsLoaded = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "pulse_apps_loaded",
		Help: "Number of apps loaded from configuration",
	})

	// performance metrics
	MessageLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "pulse_message_latency_seconds",
		Help:    "Message processing latency in seconds",
		Buckets: []float64{.001, .005, .01, .025, .05, .1, .25, .5, 1},
	}, []string{"app_key", "event_type"})
)

func GetChannelType(channelName string) string {
	if len(channelName) > 8 && channelName[:8] == "private-" {
		return "private"
	}

	if len(channelName) > 9 && channelName[:9] == "presence-" {
		return "presence"
	}

	return "public"
}
