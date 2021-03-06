// Copyright 2016 The prometheus-operator Authors
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

package alertmanager

import (
	"github.com/coreos/prometheus-operator/pkg/client/monitoring/v1alpha1"

	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/client-go/tools/cache"
)

var (
	descAlertmanagerSpecReplicas = prometheus.NewDesc(
		"prometheus_operator_alertmanager_spec_replicas",
		"Number of expected Alertmanager replicas for the Alertmanager object.",
		[]string{
			"namespace",
			"alertmanager",
		}, nil,
	)
)

type alertmanagerCollector struct {
	store cache.Store
}

func NewAlertmanagerCollector(s cache.Store) *alertmanagerCollector {
	return &alertmanagerCollector{store: s}
}

// Describe implements the prometheus.Collector interface.
func (c *alertmanagerCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- descAlertmanagerSpecReplicas
}

// Collect implements the prometheus.Collector interface.
func (c *alertmanagerCollector) Collect(ch chan<- prometheus.Metric) {
	for _, p := range c.store.List() {
		c.collectAlertmanager(ch, p.(*v1alpha1.Alertmanager))
	}
}

func (c *alertmanagerCollector) collectAlertmanager(ch chan<- prometheus.Metric, a *v1alpha1.Alertmanager) {
	ch <- prometheus.MustNewConstMetric(descAlertmanagerSpecReplicas, prometheus.GaugeValue, float64(*a.Spec.Replicas), a.Namespace, a.Name)
}
