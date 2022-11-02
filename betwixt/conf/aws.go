package conf

import "github.com/imdario/mergo"

// AwsTag defines a tag pair for an ec2 instance
type AwsTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// AwsConfiguration defines variables used to launch
// a betwixt AWS instance
type AwsConfiguration struct {
	KeypairName        string    `json:"keypairName,omitempty"`
	InstanceType       string    `json:"instanceType,omitempty"`
	IamInstanceProfile string    `json:"iamInstanceProfile,omitempty"`
	Tags               []*AwsTag `json:"tags,omitempty"`
	SubnetID           string    `json:"subnetId,omitempty"`
	Region             string    `json:"region,omitempty"`
	AMI                string    `json:"ami,omitempty"`
	SecurityGroups     []string  `json:"securityGroups,omitempty"`
	SSHUsername        string    `json:"sshUsername,omitempty"`
	SSHPrivateKeyPath  string    `json:"sshPrivateKeyPath,omitempty"`
	//BlockDeviceMapping
}

func NewAWSConfig() *AwsConfiguration {
	a := &AwsConfiguration{}
	a.InstanceType = "t2.micro"
	a.Tags = make([]*AwsTag, 0)
	a.Tags = append(a.Tags, &AwsTag{
		Key:   "betwixt",
		Value: "managed",
	})
	return a
}

// Merge merges the supplied *AwsConfiguration into the referenced AwsConfiguration
func (a *AwsConfiguration) Merge(mergingConfig *AwsConfiguration) error {
	return mergo.MergeWithOverwrite(a, mergingConfig)
}
