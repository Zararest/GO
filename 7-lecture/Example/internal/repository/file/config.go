package file

// FileRepositoryConfig contains base for file operations.
type FileRepositoryConfig struct {
	// BasePath is path to storage dir (base dir).
	BasePath string `yaml:"base-path"`
}
