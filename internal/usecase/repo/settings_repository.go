package repo

type SettingsRepository interface {
	Get(key string) (string, error)
	Set(key, value string) error
}
