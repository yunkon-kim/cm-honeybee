package infra

import (
	"errors"
	"github.com/cloud-barista/cm-honeybee/driver/network"
	modelNet "github.com/cloud-barista/cm-honeybee/model/network"
	"github.com/jollaman999/utils/logger"
)

func GetNetworkInfo() (modelNet.Network, error) {
	var n modelNet.Network
	var err error

	n.NetworkInterfaces, err = network.GetNICs()
	if err != nil {
		errMsg := "NIC: " + err.Error()
		logger.Println(logger.DEBUG, true, errMsg)

		return n, errors.New(errMsg)
	}

	n.Netfilter, err = network.GetNetfilterList()
	if err != nil {
		errMsg := "NETFILTER: " + err.Error()
		logger.Println(logger.DEBUG, true, errMsg)

		return n, errors.New(errMsg)
	}

	n.Bonding, err = network.GetBondingInfo()
	if err != nil {
		errMsg := "BONDING: " + err.Error()
		logger.Println(logger.DEBUG, true, errMsg)

		return n, errors.New(errMsg)
	}

	return n, nil
}
