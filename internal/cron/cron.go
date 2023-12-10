// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package cron

import (
	"gitea.ysicing.net/cloud/pangu/internal/prom"
	"github.com/ergoapi/util/zos"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

var Cron *Client

type Client struct {
	client *cron.Cron
}

func New() *Client {
	return &Client{client: cron.New()}
}

func (c *Client) Start() {
	logrus.Info("start cron tasks")
	c.client.Start()
	c.Default()
}

func (c *Client) Add(spec string, cmd func()) (int, error) {
	id, err := c.client.AddFunc(spec, cmd)
	if err != nil {
		return 0, err
	}
	logrus.Infof("add cron: %v", id)
	return int(id), nil
}

func (c *Client) Remove(id int) {
	c.client.Remove(cron.EntryID(id))
}

func (c *Client) Default() {
	logrus.Info("add default cron")
	id, err := c.Add("@every 30s", func() {
		logrus.Debug(zos.GetHostname())
		prom.CronRunTimesCounter.WithLabelValues("default_cron").Inc()
	})
	if err != nil {
		logrus.Errorf("add default cron error: %s", err)
		return
	}
	logrus.Infof("add default cron [%d] success", id)
}

func (c *Client) Stop() {
	logrus.Info("stop cron tasks")
	c.client.Stop()
}

func (c *Client) List() []cron.Entry {
	return c.client.Entries()
}

func init() {
	Cron = New()
}
