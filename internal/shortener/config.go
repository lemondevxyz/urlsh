package shortener

type Config struct {
	Length     int    `validate:"required,min=1" mapstructure:"length"`
	Characters string `validate:"required" mapstructure:"characters"`
}
