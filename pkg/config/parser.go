package main

import (
// "github.com/goccy/go-yaml"
)

type STP struct {
	RSTPEnabled      bool   `yaml:"rstpEnabled" json:"rstpEnabled"`
	STPGuard         string `yaml:"STPGuard" json:"STPGuard"`
	STPPortFastTrunk bool   `yaml:"STPPortFastTrunk" json:"STPPortFastTrunk"`
}

type SwitchPort struct {
	PortID       int    `yaml:"portID" json:"portID"`
	Name         string `yaml:"x" json:"x"`
	Enabled      bool   `yaml:"enabled" json:"enabled"`
	Type         string `yaml:"type" json:"type"`
	VLAN         int    `yaml:"vlan" json:"vlan"`
	VoiceVLAN    int    `yaml:"voiceVlan" json:"voiceVlan"`
	AllowedVlans string `yaml:"allowedVlans" json:"allowedVlans"`
	STP          STP
}

type Switch struct {
	Name        string `yaml:"name" json:"name"`
	SwitchPorts [48]SwitchPort
}
