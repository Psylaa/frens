package shared

type Privacy string

const (
	PrivacyPublic  Privacy = "public"
	PrivacyPrivate Privacy = "private"
)

type DataType string

const (
	DataTypeUser   DataType = "user"
	DataTypePost   DataType = "post"
	DataTypeFollow DataType = "follow"
	DataTypeToken  DataType = "token"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)
