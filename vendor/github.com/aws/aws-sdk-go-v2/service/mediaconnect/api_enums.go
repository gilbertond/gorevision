// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconnect

type Algorithm string

// Enum values for Algorithm
const (
	AlgorithmAes128 Algorithm = "aes128"
	AlgorithmAes192 Algorithm = "aes192"
	AlgorithmAes256 Algorithm = "aes256"
)

func (enum Algorithm) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Algorithm) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type KeyType string

// Enum values for KeyType
const (
	KeyTypeSpeke     KeyType = "speke"
	KeyTypeStaticKey KeyType = "static-key"
)

func (enum KeyType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum KeyType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Protocol string

// Enum values for Protocol
const (
	ProtocolZixiPush Protocol = "zixi-push"
	ProtocolRtpFec   Protocol = "rtp-fec"
	ProtocolRtp      Protocol = "rtp"
	ProtocolZixiPull Protocol = "zixi-pull"
	ProtocolRist     Protocol = "rist"
)

func (enum Protocol) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Protocol) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SourceType string

// Enum values for SourceType
const (
	SourceTypeOwned    SourceType = "OWNED"
	SourceTypeEntitled SourceType = "ENTITLED"
)

func (enum SourceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SourceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Status string

// Enum values for Status
const (
	StatusStandby  Status = "STANDBY"
	StatusActive   Status = "ACTIVE"
	StatusUpdating Status = "UPDATING"
	StatusDeleting Status = "DELETING"
	StatusStarting Status = "STARTING"
	StatusStopping Status = "STOPPING"
	StatusError    Status = "ERROR"
)

func (enum Status) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Status) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
