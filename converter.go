package restclient

import "io"

type Encoder interface {
	Encode(o interface{}) (int64, error)
}

type Decoder interface {
	Decode(o interface{}) (int64, error)
}

type Converter interface {
	CreateEncoder(io.Writer) Encoder
	CreateDecoder(io.Reader) Decoder

	CanEncode(o interface{}, mediaType MediaType) bool
	CanDecode(o interface{}, mediaType MediaType) bool
	SupportMediaType() []MediaType
}
