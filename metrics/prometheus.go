// Copyright 2017 the lostromos Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// https://prometheus.io/docs/practices/naming/ is what we are basing naming conventions off of.
var (
	// CreateFailures is a metric for the number of failures to create a release
	CreateFailures = prometheus.NewCounter(prometheus.CounterOpts{
		Help:      "The number of failed create events",
		Name:      "create_error_total",
		Namespace: "releases",
	})

	// CreatedReleases is a metric for the number of releases created by this operator
	CreatedReleases = prometheus.NewCounter(prometheus.CounterOpts{
		Help:      "The number of successful create events",
		Name:      "create_total",
		Namespace: "releases",
	})

	// DeleteFailures is a metric for the number of times a delete by this operator
	DeleteFailures = prometheus.NewCounter(prometheus.CounterOpts{
		Help:      "The number of failed delete events",
		Name:      "delete_error_total",
		Namespace: "releases",
	})

	// DeletedReleases is a metric for the number of releases deleted by this operator
	DeletedReleases = prometheus.NewCounter(prometheus.CounterOpts{
		Help:      "The number of successful delete events",
		Name:      "delete_total",
		Namespace: "releases",
	})

	// ManagedReleases is a metric of the number of current releases managed by the operator
	ManagedReleases = prometheus.NewGauge(prometheus.GaugeOpts{
		Help:      "The number of releases managed by this operator",
		Name:      "total",
		Namespace: "releases",
	})

	// UpdateFailures is a metric for the number of times an update operation failed by this operator
	UpdateFailures = prometheus.NewCounter(prometheus.CounterOpts{
		Help:      "The number of failed update events",
		Name:      "update_error_total",
		Namespace: "releases",
	})

	// UpdatedReleases is a metric for the number of releases updated successfully
	UpdatedReleases = prometheus.NewCounter(prometheus.CounterOpts{
		Help:      "The number of successful update events",
		Name:      "update_total",
		Namespace: "releases",
	})

	// TotalEvents is a metric for the number of events that have been handled by this operator
	TotalEvents = prometheus.NewCounter(prometheus.CounterOpts{
		Help:      "The number of events (create/delete/updates) processed by this operator",
		Name:      "events_total",
		Namespace: "releases",
	})
)

func init() {
	prometheus.MustRegister(CreatedReleases)
	prometheus.MustRegister(CreateFailures)
	prometheus.MustRegister(DeletedReleases)
	prometheus.MustRegister(DeleteFailures)
	prometheus.MustRegister(ManagedReleases)
	prometheus.MustRegister(UpdatedReleases)
	prometheus.MustRegister(UpdateFailures)
	prometheus.MustRegister(TotalEvents)
}
