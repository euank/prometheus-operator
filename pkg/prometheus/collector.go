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

package prometheus

import (
	"github.com/coreos/prometheus-operator/pkg/client/monitoring/v1alpha1"

	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/client-go/tools/cache"
)

var (
	descPrometheusSpecReplicas = prometheus.NewDesc(
		"prometheus_operator_prometheus_spec_replicas",
		"Number of expected Prometheus replicas for the Prometheus object.",
		[]string{
			"namespace",
			"prometheus",
		}, nil,
	)
)

type prometheusCollector struct {
	store cache.Store
}

func NewPrometheusCollector(s cache.Store) *prometheusCollector {
	return &prometheusCollector{store: s}
}

// Describe implements the prometheus.Collector interface.
func (c *prometheusCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- descPrometheusSpecReplicas
}

// Collect implements the prometheus.Collector interface.
func (c *prometheusCollector) Collect(ch chan<- prometheus.Metric) {
	for _, p := range c.store.List() {
		c.collectPrometheus(ch, p.(*v1alpha1.Prometheus))
	}
}

func (c *prometheusCollector) collectPrometheus(ch chan<- prometheus.Metric, p *v1alpha1.Prometheus) {
	ch <- prometheus.MustNewConstMetric(descPrometheusSpecReplicas, prometheus.GaugeValue, float64(*p.Spec.Replicas), p.Namespace, p.Name)
}
