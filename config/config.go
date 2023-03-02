package config

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// auto
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql Pgsql `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
}
