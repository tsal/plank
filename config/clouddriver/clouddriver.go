package clouddriver

// Clouddriver mirrors clouddriver.yaml files on disk
type Clouddriver struct {
	DockerRegistry struct {
		Enabled  bool            `json:"enabled,omitempty" mapstructure:"enabled"`
		Accounts []DockerAccount `json:"accounts,omitempty" mapstructure:"accounts"`
	} `json:"dockerRegistry,omitempty" mapstructure:"dockerRegistry"`

	Kubernetes struct {
		Enabled  bool                `json:"enabled,omitempty" mapstructure:"enabled"`
		Accounts []KubernetesAccount `json:"accounts,omitempty" mapstructure:"accounts"`
	} `json:"kubernetes,omitempty" mapstructure:"kubernetes"`

	AWS struct {
		Enabled           bool         `json:"enabled,omitempty" mapstructure:"enabled"`
		DefaultAssumeRole string       `json:"defaultAssumeRole,omitempty" mapstructure:"defaultAssumeRole"`
		DefaultRegions    []AWSRegion  `json:"defaultRegions,omitempty" mapstructure:"defaultRegions"`
		Accounts          []AWSAccount `json:"accounts,omitempty" mapstructure:"accounts"`
	} `json:"aws,omitempty" mapstructure:"aws"`

	GCP struct {
		Enabled  bool         `json:"enabled,omitempty" mapstructure:"enabled"`
		Accounts []GCPAccount `json:"accounts,omitempty" mapstructure:"accounts"`
	} `json:"google,omitempty" mapstructure:"google"`

	Artifacts struct {
		Github struct {
			Enabled bool                  `json:"enabled,omitempty" mapstructure:"enabled"`
			Account GithubArtifactAccount `json:"accounts,omitempty" mapstructure:"accounts"`
		} `json:"github,omitempty" mapstructure:"github"`
	} `json:"artifacts,omitempty" mapstructure:"artifacts"`
}

// GithubArtifactAccount settings
type GithubArtifactAccount struct {
	Name                 string `json:"name,omitempty" mapstructure:"name"`
	Username             string `json:"username,omitempty" mapstructure:"username"`
	Password             string `json:"password,omitempty" mapstructure:"password"`
	UsernamePasswordFile string `json:"usernamePasswordFile,omitempty" mapstructure:"usernamePasswordFile"`
	Token                string `json:"token,omitempty" mapstructure:"token"`
	TokenFile            string `json:"tokenFile,omitempty" mapstructure:"tokenFile"`
}

// DockerAccount settings
type DockerAccount struct {
	Name         string   `json:"name,omitempty" mapstructure:"name"`
	Username     string   `json:"username,omitempty" mapstructure:"username"`
	PasswordFile string   `json:"passwordFile,omitempty" mapstructure:"passwordFile"`
	Address      string   `json:"address,omitempty" mapstructure:"address"`
	Repositories []string `json:"repositories,omitempty" mapstructure:"repositories"`
}

// KubernetesAccount settings
type KubernetesAccount struct {
	Name             string           `json:"name,omitempty" mapstructure:"name"`
	Namespaces       []string         `json:"namespaces,omitempty" mapstructure:"namespaces"`
	KubeconfigFile   string           `json:"kubeconfigFile,omitempty" mapstructure:"kubeconfigFile"`
	ProviderVersion  string           `json:"providerVersion,omitempty" mapstructure:"providerVersion"`
	ServiceAccount   bool             `json:"serviceAccount,omitempty" mapstructure:"serviceAccount"`
	DockerRegistries []DockerRegistry `json:"dockerRegistries,omitempty" mapstructure:"dockerRegistries"`
}

// DockerRegistry settings
type DockerRegistry struct {
	AccountName string `json:"accountName,omitempty" mapstructure:"accountName"`
}

// AWSAccount settings
type AWSAccount struct {
	Name      string      `json:"name,omitempty" mapstructure:"name"`
	AccountID string      `json:"accountId,omitempty" mapstructure:"accountId"`
	Regions   []AWSRegion `json:"regions,omitempty" mapstructure:"regions"`
}

// AWSRegion settings
type AWSRegion struct {
	Name string `json:"name,omitempty" mapstructure:"name"`
}

// GCPAccount settings
type GCPAccount struct {
	Name     string `json:"name,omitempty" mapstructure:"name"`
	Project  string `json:"project,omitempty" mapstructure:"project"`
	JSONPath string `json:"jsonPath,omitempty" mapstructure:"jsonPath"`
}
