package zwave

type SerialApiCapabilitiesResponse struct {
	CommandId           byte
	ApplicationVersion  byte
	ApplicationRevision byte
	Manufacturer1       byte
	Manufacturer2       byte
	ProductType1        byte
	ProductType2        byte
	ProductId1          byte
	ProductId2          byte
	SupportedFunctions  []byte
}

func ParseSerialApiCapabilitiesResponse(payload []byte) *SerialApiCapabilitiesResponse {
	val := &SerialApiCapabilitiesResponse{
		CommandId:           payload[0],
		ApplicationVersion:  payload[1],
		ApplicationRevision: payload[2],
		Manufacturer1:       payload[3],
		Manufacturer2:       payload[4],
		ProductType1:        payload[5],
		ProductType2:        payload[6],
		ProductId1:          payload[7],
		ProductId2:          payload[8],
		SupportedFunctions:  payload[9:],
	}

	return val
}

func (n *SerialApiCapabilitiesResponse) GetSupportedFunctions() []uint8 {
	supportedFunctions := []uint8{}
	var funcId uint8 = 1

	var i, j uint8
	for i = 0; i < uint8(len(n.SupportedFunctions)); i++ {
		for j = 0; j < 8; j++ {
			if isBitSet(n.SupportedFunctions[i], j) {
				supportedFunctions = append(supportedFunctions, funcId)
			}

			funcId++
		}
	}

	return supportedFunctions
}
