package postgres

// TODO: add conns limits
type DbConfig struct {
	PGReadNodes  []string `yaml:"pg_read_nodes"`
	PGWriteNodes []string `yaml:"pg_write_nodes"`
}
