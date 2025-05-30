package instrumentation

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// TODO Update metric names according to: https://prometheus.io/docs/instrumenting/writing_exporters/#naming
const (
	NameSpace                                      = "content_sources"
	PulpConnectivity                               = "pulp_connectivity"
	HttpStatusHistogram                            = "http_status_histogram"
	RepositoriesTotal                              = "repositories_total"
	RepositoryConfigsTotal                         = "repository_configs_total"
	PublicRepositories36HourIntrospectionTotal     = "public_repositories_36_hour_introspection_total"
	PublicRepositoriesWithFailedIntrospectionTotal = "public_repositories_with_failed_introspection_total"
	CustomRepositories36HourIntrospectionTotal     = "custom_repositories_36_hour_introspection_total"
	MessageLatency                                 = "message_latency"
	MessageResultTotal                             = "message_result_total"
	OrgTotal                                       = "org_total"
	TaskStats                                      = "task_stats"
	TaskStatsLabelPendingCount                     = "task_stats_pending_count"
	TaskStatsLabelOldestWait                       = "task_stats_oldest_wait"
	TaskStatsLabelAverageWait                      = "task_stats_average_wait"
	RHReposSnapshotNotCompletedInLast36HoursCount  = "rh_repos_snapshot_not_completed_in_last_36_hour_count"
	TaskPendingTimeAverageByType                   = "task_pending_time_average_by_type"
	TemplatesCount                                 = "templates_count"
	TemplatesUseLatestCount                        = "templates_use_latest_count"
	TemplatesUseDateCount                          = "templates_use_date_count"
	TemplatesUpdatedInLast24HoursCount             = "templates_updated_in_last_24_hour_count"
	TemplatesAgeAverage                            = "templates_age_average"
	CertificateExpiryDays                          = "certificate_expiry_days"
)

type Metrics struct {
	HttpStatusHistogram prometheus.HistogramVec

	// Custom metrics
	PulpConnectivity                               prometheus.Gauge
	RepositoriesTotal                              prometheus.Gauge
	RepositoryConfigsTotal                         prometheus.Gauge
	PublicRepositories36HourIntrospectionTotal     prometheus.GaugeVec
	PublicRepositoriesWithFailedIntrospectionTotal prometheus.Gauge
	CustomRepositories36HourIntrospectionTotal     prometheus.GaugeVec
	MessageResultTotal                             prometheus.CounterVec
	MessageLatency                                 prometheus.Histogram
	TaskStats                                      prometheus.GaugeVec
	OrgTotal                                       prometheus.Gauge
	RHReposSnapshotNotCompletedInLast36HoursCount  prometheus.Gauge
	TaskPendingTimeAverageByType                   prometheus.GaugeVec
	TemplatesCount                                 prometheus.Gauge
	TemplatesUseLatestCount                        prometheus.Gauge
	TemplatesUseDateCount                          prometheus.Gauge
	TemplatesUpdatedInLast24HoursCount             prometheus.Gauge
	TemplatesAgeAverage                            prometheus.Gauge
	CertificateExpiryDays                          prometheus.GaugeVec
	reg                                            *prometheus.Registry
}

// See: https://consoledot.pages.redhat.com/docs/dev/platform-documentation/understanding-slo.html
// See: https://prometheus.io/docs/tutorials/understanding_metric_types/#types-of-metrics
func NewMetrics(reg *prometheus.Registry) *Metrics {
	if reg == nil {
		panic("reg cannot be nil")
	}
	metrics := &Metrics{
		reg: reg,
		PulpConnectivity: promauto.With(reg).NewGauge(prometheus.GaugeOpts{
			Namespace: NameSpace,
			Name:      PulpConnectivity,
			Help:      "Status of pulp connection",
		}),
		HttpStatusHistogram: *promauto.With(reg).NewHistogramVec(prometheus.HistogramOpts{
			Namespace: NameSpace,
			Name:      HttpStatusHistogram,
			Help:      "Duration of HTTP requests",
			Buckets:   prometheus.DefBuckets,
		}, []string{"status", "method", "path"}),

		MessageLatency: promauto.With(reg).NewHistogram(prometheus.HistogramOpts{
			Namespace: NameSpace,
			Name:      MessageLatency,
			Help:      "Time to pickup task messages",
			//                        1m  5m   30m   1h    2h    3h     5h     10h
			Buckets: []float64{.5, 1, 60, 300, 1800, 3600, 7200, 10800, 18000, 36000},
		}),
		MessageResultTotal: *promauto.With(reg).NewCounterVec(prometheus.CounterOpts{
			Namespace:   NameSpace,
			Name:        MessageResultTotal,
			Help:        "Result of task messages",
			ConstLabels: nil,
		}, []string{"state"}),
		RepositoriesTotal: promauto.With(reg).NewGauge(prometheus.GaugeOpts{
			Namespace: NameSpace,
			Name:      RepositoriesTotal,
			Help:      "Number of repositories",
		}),
		RepositoryConfigsTotal: promauto.With(reg).NewGauge(prometheus.GaugeOpts{
			Namespace: NameSpace,
			Name:      RepositoryConfigsTotal,
			Help:      "Number of repository configurations",
		}),
		TaskStats: *promauto.With(reg).NewGaugeVec(prometheus.GaugeOpts{
			Namespace: NameSpace,
			Name:      TaskStats,
			Help:      "Stats around Tasks",
		}, []string{"label"}),
		PublicRepositories36HourIntrospectionTotal: *promauto.With(reg).NewGaugeVec(prometheus.GaugeOpts{
			Namespace: NameSpace,
			Name:      PublicRepositories36HourIntrospectionTotal,
			Help:      "Breakdown of public repository count by those that attempted introspection and those that missed introspection.",
		}, []string{"status"}),
		PublicRepositoriesWithFailedIntrospectionTotal: promauto.With(reg).NewGauge(prometheus.GaugeOpts{
			Namespace: NameSpace,
			Name:      PublicRepositoriesWithFailedIntrospectionTotal,
			Help:      "Number of repositories with failed introspection",
		}),
		CustomRepositories36HourIntrospectionTotal: *promauto.With(reg).NewGaugeVec(prometheus.GaugeOpts{
			Namespace: NameSpace,
			Name:      CustomRepositories36HourIntrospectionTotal,
			Help:      "Breakdown of custom repository count by those that attempted introspection and those that missed introspection.",
		}, []string{"status"}),
		OrgTotal: promauto.With(reg).NewGauge(prometheus.GaugeOpts{
			Namespace: NameSpace,
			Name:      OrgTotal,
			Help:      "Number of organizations with at least one repository.",
		}),
		RHReposSnapshotNotCompletedInLast36HoursCount: promauto.With(reg).NewGauge(prometheus.GaugeOpts{
			Namespace: NameSpace,
			Name:      RHReposSnapshotNotCompletedInLast36HoursCount,
			Help:      "Number of Red Hat repositories that haven't had successful snapshot task in the last 36 hours.",
		}),
		TaskPendingTimeAverageByType: *promauto.With(reg).NewGaugeVec(prometheus.GaugeOpts{
			Namespace: NameSpace,
			Name:      TaskPendingTimeAverageByType,
			Help:      "Average pending time of the tasks by their type.",
		}, []string{"task_type"}),
		TemplatesCount: promauto.With(reg).NewGauge(prometheus.GaugeOpts{
			Namespace: NameSpace,
			Name:      TemplatesCount,
			Help:      "Total number of templates.",
		}),
		TemplatesUseLatestCount: promauto.With(reg).NewGauge(prometheus.GaugeOpts{
			Namespace: NameSpace,
			Name:      TemplatesUseLatestCount,
			Help:      "Number of templates which are set to use latest.",
		}),
		TemplatesUseDateCount: promauto.With(reg).NewGauge(prometheus.GaugeOpts{
			Namespace: NameSpace,
			Name:      TemplatesUseDateCount,
			Help:      "Number of templates which have a set date.",
		}),
		TemplatesUpdatedInLast24HoursCount: promauto.With(reg).NewGauge(prometheus.GaugeOpts{
			Namespace: NameSpace,
			Name:      TemplatesUpdatedInLast24HoursCount,
			Help:      "Number of templates that have been updated in the last 24 hours.",
		}),
		TemplatesAgeAverage: promauto.With(reg).NewGauge(prometheus.GaugeOpts{
			Namespace: NameSpace,
			Name:      TemplatesAgeAverage,
			Help:      "Average age (days between the set template date and now) of templates.",
		}),
		CertificateExpiryDays: *promauto.With(reg).NewGaugeVec(prometheus.GaugeOpts{
			Namespace: NameSpace,
			Name:      CertificateExpiryDays,
			Help:      "Number of days until the certificate expires by the certificate label",
		}, []string{"certificate_label"}),
	}

	reg.MustRegister(collectors.NewBuildInfoCollector())

	return metrics
}

func (m *Metrics) RecordMessageResult(success bool) {
	status := "failed"
	if success {
		status = "success"
	}
	if m != nil {
		m.MessageResultTotal.With(prometheus.Labels{"state": status}).Inc()
	}
}
func (m *Metrics) RecordMessageLatency(msgTime time.Time) {
	diff := time.Since(msgTime)
	m.MessageLatency.Observe(diff.Seconds())
}

func (m Metrics) Registry() *prometheus.Registry {
	return m.reg
}
