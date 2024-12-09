package config

import (
	"bytes"
	"fmt"
)

type Postgres struct {
	HostProp     string `yaml:"host"`
	PortProp     string `yaml:"port"`
	UserProp     string `yaml:"user"`
	PasswordProp string `yaml:"password"`
	DatabaseProp string `yaml:"database"`
}

func (r *Postgres) Host() string {
	return loadStringProp(r.HostProp)
}

func (r *Postgres) Port() int {
	return loadIntProp(r.PortProp)
}

func (r *Postgres) User() string {
	return loadStringProp(r.UserProp)
}

func (r *Postgres) Password() string {
	return loadStringProp(r.PasswordProp)
}

func (r *Postgres) Database() string {
	return loadStringProp(r.DatabaseProp)
}

func (r *Postgres) Addr() string {
	return fmt.Sprintf("%s:%d", r.Host(), r.Port())
}

func (r *Postgres) DSN() string {
	dsn := bytes.NewBufferString(fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s",
		r.Host(),
		r.Port(),
		r.User(),
		r.Password(),
		r.Database(),
		"disable",
		"Europe/London",
	))

	return dsn.String()
}
