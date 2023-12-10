// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package prom

import "github.com/prometheus/client_golang/prometheus"

const (
	ns = "pangu"
)

var (
	CronRunTimesCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: ns,
		Name:      "cron_run_time",
		Help:      "cron执行次数",
	}, []string{"name"})
)

func init() {
	prometheus.MustRegister(CronRunTimesCounter)
}
