// Copyright (c) 2021 Kien Nguyen-Tuan <kiennt2609@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config is the top-level configuration.
type Config struct {
	Auth Auth `yaml:"auth"`
	// Format is the output format. Markdown table by default.
	Format string `yaml:"format"`
	// Topic is the awesome main topic
	// For example, you want to create a Golang language awesome
	// currated list, the `topic` is `go|golang`
	// Checkout the github topics: https://github.com/topics
	Topic string `yaml:"topic"`
	// SubTopics is the list of sub sections in the awesome list.
	SubTopics []string `yaml:"sub_topics"`
	// OutputFile is the path of README file.
	OutputFile string `yaml:"output_file"`
}

// Auth is Github authentication configuration.
type Auth struct {
	// Username is Github username
	Username string `yaml:"username"`
	// Password is Github password
	Password string `yaml:"password"`
	// OTP is one-time password for users with two-factor auth enabled
	OTP string
	// AccessToken is Github person API token:
	// https://github.com/blog/1509-personal-api-tokens
	// If this field is specified, other fields (username, password & otp)
	// will be ignored
	AccessToken string `yaml:"access_token"`
}

// DefaultConfig is the default top-level configuration.
var DefaultConfig = Config{
	OutputFile: "README.md",
}

// UnmarshalYAML implements the yaml.Unmarshaler interface
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = DefaultConfig
	// We want to set c to the defaults and then overwrite it with the input.
	// To make unmarshal fill the plain data struct rather than calling UnmarshalYAML
	// again, we have to hide it using a type indirection.
	type plain Config
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}
	return nil
}

// String represents Configuration instance as string.
func (c *Config) String() string {
	b, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Sprintf("<error creating config string: %s>", err)
	}
	return string(b)
}

// Load parses the YAML input s into a Config.
func Load(s string) (*Config, error) {
	cfg := &Config{}
	err := yaml.UnmarshalStrict([]byte(s), cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// LoadFile parses the given YAML file into a Config.
func LoadFile(filename string) (*Config, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	cfg, err := Load(string(content))
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
