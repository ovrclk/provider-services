package rest

import (
	cltypes "github.com/akash-network/provider/cluster/types/v1beta3"
)

type LeasedIPStatus struct {
	Port         uint32 `json:"port"`
	ExternalPort uint32 `json:"external_port"`
	Protocol     string `json:"protocol"`
	IP           string `json:"ip"`
}

type LeaseStatus struct {
	Services       map[string]*cltypes.ServiceStatus        `json:"services"`
	ForwardedPorts map[string][]cltypes.ForwardedPortStatus `json:"forwarded_ports"` // Container services that are externally accessible
	IPs            map[string][]LeasedIPStatus              `json:"ips"`
}
