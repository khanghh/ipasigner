package provisioning

import (
	"bytes"
	"crypto/x509"
	"time"

	"howett.net/plist"
)

type MobileProvision struct {
	AppIDName                   string    `plist:"AppIDName"`
	ApplicationIdentifierPrefix []string  `plist:"ApplicationIdentifierPrefix"`
	CreationDate                time.Time `plist:"CreationDate"`
	Platform                    []string  `plist:"Platform"`
	IsXcodeManaged              bool      `plist:"IsXcodeManaged"`
	DeveloperCertificates       [][]byte  `plist:"DeveloperCertificates"`
	DEREncodeProfile            []byte    `plist:"DER-Encode-Profile"`
	Entitlements                struct {
		ApsEnvironment        string   `plist:"aps-environment"`
		ApplicationIDentifier string   `plist:"application-identifier"`
		KeychainAccessGroups  []string `plist:"keychain-access-groups"`
		GetTaskAllow          bool     `plist:"get-task-allow"`
		TeamIdentifier        string   `plist:"com.apple.developer.team-identifier"`
	} `plist:"Entitlements"`
	ExpirationDate       time.Time `plist:"ExpirationDate"`
	Name                 string    `plist:"Name"`
	ProvisionsAllDevices bool      `plist:"ProvisionsAllDevices"`
	ProvisionedDevices   []string  `plist:"ProvisionedDevice"`
	TeamIdentifier       []string  `plist:"TeamIdentifier"`
	TeamName             string    `plist:"TeamName"`
	TimeToLive           int       `plist:"TimeToLive"`
	UUID                 string    `plist:"UUID"`
	Version              int       `plist:"Version"`
}

func (p *MobileProvision) GetDeveloperCertificates() ([]*x509.Certificate, error) {
	certificates := make([]*x509.Certificate, len(p.DeveloperCertificates))

	for i := range p.DeveloperCertificates {
		certificate, err := x509.ParseCertificate(p.DeveloperCertificates[i])

		if err != nil && certificate == nil {
			return nil, err
		}

		certificates[i] = certificate
	}

	return certificates, nil
}

func (p *MobileProvision) IsProvisionedDevice(udid string) bool {
	return p.ProvisionsAllDevices || func(udid string) bool {
		for _, id := range p.ProvisionedDevices {
			if id == udid {
				return true
			}
		}

		return false
	}(udid)
}

func (p *MobileProvision) IsExpired(time time.Time) bool {
	return time.After(p.CreationDate) && time.Before(p.ExpirationDate)
}

func NewMobileProvision(content []byte) *MobileProvision {
	buf := bytes.NewReader(content)
	var data MobileProvision
	decoder := plist.NewDecoder(buf)
	decoder.Decode(&data)

	return &data
}
