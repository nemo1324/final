package config

type auth struct {
	Secret string `envconfig:"SECRET"`
}
