package validator

const (
	// KvpTypeRegexString Used to check the valid of the KVP type
	KvpTypeRegexString = "(?i)^(general|raw|id|pii|id.external|result|error).[a-z]+$"
	// Base64RegexString Used to check the valid base64 string
	Base64RegexString = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
)
