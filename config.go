package ossinspector

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	PackageTrustRule PackageTrustRule `yaml:"package_trust_rule"`
}
type AuthorRule struct {
	Age          int `yaml:"age"`
	Repos        int `yaml:"repos"`
	Followers    int `yaml:"followers"`
	Contributors int `yaml:"contributors"`
}
type RepoRule struct {
	Stars        int `yaml:"stars"`
	Forks        int `yaml:"forks"`
	Watchers     int `yaml:"watchers"`
	Contributors int `yaml:"contributors"`
	Commits      int `yaml:"commits"`
	LastRelease  int `yaml:"last_release"`
}
type CommitRule struct {
	LastCommitAge int `yaml:"last_commit_age"`
}
type PackageTrustRule struct {
	AuthorRules  []AuthorRule  `yaml:"author_rule"`
	RepoRules    []RepoRule    `yaml:"repo_rule"`
	CommitRules []CommitRule `yaml:"commit_rules"`
}

func NewConfig(path string) (*Config, error) {
	var config Config
	bytes, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
