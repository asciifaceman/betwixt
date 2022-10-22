package conf

// Conf is the primary configuration wrapper for
// serializing to and from disk
type Conf struct {
}

// AwsConfiguration defines variables used to launch
// a betwixt AWS instance
type AwsConfiguration struct {
	KeypairName        string    `json:"keypairName"`
	InstanceType       string    `json:"instanceType"`
	IamInstanceProfile string    `json:"iamInstanceProfile"`
	Tags               []*AwsTag `json:"tags"`
	SubnetID           string    `json:"subnetId"`
	Region             string    `json:"region"`
	AMI                string    `json:"ami"`
	SecurityGroups     []string  `json:"securityGroups"`
	SSHUsername        string    `json:"sshUsername"`
	SSHPrivateKeyPath  string    `json:"sshPrivateKeyPath"`
	//BlockDeviceMapping
}

// AwsTag defines a tag pair for an ec2 instance
type AwsTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
