package applications

import (
	"github.com/armory/plank/client"
	"time"
	"log"
)

// DefaultService to use when no service is received.
var DefaultService *Service

var defaultOrcaURL = "http://armory-orca:8083"
var defaultFront50URL = "http://armory-front50:8080"

type postGetter interface {
	Post(path string, contentType client.ContentType, body, dest interface{}) error
	Get(path string, dest interface{}) error
}

// Service for interacting with the applications API.
type Service struct {
	client     postGetter
	orcaURL    string
	front50URL string
	pollTime   time.Duration
	log        log.Logger
}

// Option for configuring a service.
type Option func(*Service) error

// NewService for checking permissions.
func NewService(options ...Option) *Service {
	defClient, _ := client.New()
	s := &Service{
		client:   defClient,
		orcaURL:  defaultOrcaURL,
		pollTime: 7 * time.Minute,
	}
	for _, option := range options {
		// TODO: handle errors
		option(s)
	}
	return s
}

// PollTime to be used when polling. Useful in integration tests.
func PollTime(t time.Duration) Option {
	return func(s *Service) error {
		s.pollTime = t
		return nil
	}
}

// OrcaURL option to change the URL used to talk to Orca.
func OrcaURL(url string) Option {
	return func(s *Service) error {
		s.orcaURL = url
		return nil
	}
}

// Front50URL option to change the URL used to talk to Gate.
func Front50URL(url string) Option {
	return func(s *Service) error {
		s.front50URL = url
		return nil
	}
}
