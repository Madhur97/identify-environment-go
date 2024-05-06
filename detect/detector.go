package detect

// Detector defines the interface for fetching identifier for an environment
type Detector interface {
	GetIdentifiers() ([]string, error)
}
