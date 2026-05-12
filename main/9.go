package main

import (
	"fmt"
	"time"
)

type config struct {
	timeout time.Duration
	retry int
	baseURL string
}

type Service struct {
	config config
}

type Option func(*config)

func WithTimeOut(d time.Duration) Option {
	return func(c *config) {
		c.timeout = d
	}
}

func WithRetries(r int) Option {
	return func(c *config) {
		c.retry = r
	}
}

func WithBaseURL(url string) Option {
	return func(c *config) {
		c.baseURL = url
	}
}

// default is set, because we aren't passing any parameters to set the config struct fields
func newDefaultConfig() (*config) {
	return &config{}
}

func NewService(opt ...Option) (s *Service) {
	cfg := newDefaultConfig()

	for _, opt := range opt {
		opt(cfg)
	}
	return &Service{ config: *cfg}
}

func _9() {
	service1 := NewService(
		WithTimeOut(2 * time.Second), 
		WithRetries(3), 
		WithBaseURL("https://api.service1.com"))

	service2 := NewService(
		WithRetries(3), 
		WithBaseURL("https://api.service2.com")) // since each service is created with different parameters, in order to have a generic constructor function, we need to wrap the parameters into a common thing so that we could use variadic constructor in which order of parameters and the number of parameters wont matter

	fmt.Printf("%+v\n", service1)
	fmt.Printf("%+v\n", service2)
	fmt.Println(c)
}
