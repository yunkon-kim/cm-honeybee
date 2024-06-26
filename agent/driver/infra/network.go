package infra

import (
	"errors"
	"github.com/cloud-barista/cm-honeybee/agent/driver/network"
	modelNet "github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/model/onprem/network"
	"github.com/jollaman999/utils/logger"
	"sync"
)

var networkInfoLock sync.Mutex

func GetNetworkInfo() (modelNet.Network, error) {
	if !networkInfoLock.TryLock() {
		return modelNet.Network{}, errors.New("network info collection is in progress")
	}
	defer func() {
		networkInfoLock.Unlock()
	}()

	var n modelNet.Network
	var err error

	n.Host.NetworkInterface, err = network.GetNICs()
	if err != nil {
		errMsg := "NIC: " + err.Error()
		logger.Println(logger.DEBUG, true, errMsg)

		return n, errors.New(errMsg)
	}

	n.Host.Route, err = network.GetRoutes()
	if err != nil {
		errMsg := "ROUTES: " + err.Error()
		logger.Println(logger.DEBUG, true, errMsg)

		return n, errors.New(errMsg)
	}

	n.Host.FirewallRule, err = network.GetFirewallRules()
	if err != nil {
		errMsg := "FIREWALL RULE: " + err.Error()
		logger.Println(logger.DEBUG, true, errMsg)

		return n, errors.New(errMsg)
	}

	n.Host.DNS, err = network.GetDNS()
	if err != nil {
		errMsg := "DNS: " + err.Error()
		logger.Println(logger.DEBUG, true, errMsg)

		return n, errors.New(errMsg)
	}

	return n, nil
}
