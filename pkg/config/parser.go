package config

import (
// "github.com/goccy/go-yaml"
)

type STP struct {
	RSTPEnabled      bool   `yaml:"rstpEnabled" json:"rstpEnabled"`
	STPGuard         string `yaml:"stpGuard" json:"stpGuard"`
	STPPortFastTrunk bool   `yaml:"stpPortFastTrunk" json:"stpPortFastTrunk"`
}

type SwitchPort struct {
	PortID       int    `yaml:"portId" json:"portId"`
	Name         string `yaml:"name" json:"name"`
	Enabled      bool   `yaml:"enabled" json:"enabled"`
	Type         string `yaml:"type" json:"type"`
	VLAN         int    `yaml:"vlan" json:"vlan"`
	VoiceVLAN    int    `yaml:"voiceVlan" json:"voiceVlan"`
	AllowedVlans string `yaml:"allowedVlans" json:"allowedVlans"`
	STP          STP    `yaml:"stp" json:"stp"`
}

type Switch struct {
	Name        string       `yaml:"name" json:"name"`
	SwitchPorts []SwitchPort `yaml:"switchports" json:"switchport"`
}

type NetworkState struct {
	Version        string `yaml:"version" json:"version"`
	OrganizationID string `yaml:"organization_id" json:"organization_id"`
	NetworkID      string `yaml:"network_id" json:"network_id"`
	Switches       []Switch
}
