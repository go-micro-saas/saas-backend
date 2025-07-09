package global

const (
	KeyPrefix = "saas_backend_"
)

func Key(k string) string {
	return KeyPrefix + k
}
