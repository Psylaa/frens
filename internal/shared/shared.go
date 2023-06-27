package shared

type Privacy string

const (
	PrivacyPublic  Privacy = "public"
	PrivacyPrivate Privacy = "private"
)

type DataType string

const (
	DataTypeUser DataType = "user"
	DataTypePost DataType = "post"
)
