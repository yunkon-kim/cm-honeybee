// Getting network interfaces for Linux

//go:build linux

package network

import (
	"github.com/cloud-barista/cm-honeybee/lib/routes"
	"github.com/cloud-barista/cm-honeybee/model/network"
	"github.com/shirou/gopsutil/v3/net"
)

func GetNICs() ([]network.NIC, error) {
	var nics []network.NIC

	interfaces, err := net.Interfaces()
	if err != nil {
		return nics, err
	}

	var defaultRoutes []routes.RouteStruct
	var allRoutes []routes.RouteStruct

	defaultRoutes, err = routes.GetLinuxRoutes(true)
	if err != nil {
		return nics, err
	}

	allRoutes, err = routes.GetLinuxRoutes(false)
	if err != nil {
		return nics, err
	}

	for _, i := range interfaces {
		var addresses []string
		var gateways []string
		var ros []network.Route

		for _, a := range i.Addrs {
			addresses = append(addresses, a.Addr)
		}

		for _, route := range defaultRoutes {
			if route.Interface == i.Name {
				gateways = append(gateways, route.NextHop)
			}
		}

		for _, route := range allRoutes {
			if route.Interface == i.Name {
				ros = append(ros, network.Route{
					Destination: route.Destination,
					Netmask:     route.Netmask,
					NextHop:     route.NextHop,
				})
			}
		}

		nics = append(nics, network.NIC{
			Interface: i.Name,
			Address:   addresses,
			Gateway:   gateways,
			Route:     ros,
			MAC:       i.HardwareAddr,
			MTU:       i.MTU,
		})
	}

	return nics, nil
}
