package provisioning

import "github.com/github/ietf-cms/protocol"

func ParseMobileProvision(data []byte) (*MobileProvision, error) {
	content, err := protocol.ParseContentInfo(data)
	if err != nil {
		return nil, err
	}
	signedData, err := content.SignedDataContent()
	if err != nil {
		return nil, err
	}
	buf, err := signedData.EncapContentInfo.DataEContent()
	if err != nil {
		return nil, err
	}
	return NewMobileProvision(buf), nil
}
