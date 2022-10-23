package conf

// AwsState defines runtime details required to
// identify existing instances at a project level
type AwsState struct {
	InstanceID string `json:"instanceID"`
}

func (a *AwsState) SetInstanceID(instanceID string) {
	a.InstanceID = instanceID
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

func (a *AwsConfiguration) MergeLocal(local *LocalAwsConfiguration) error {
	if local == nil {
		return nil
	}

	// TODO: Make this smarter, mergo won't work because the structs are different
	// but we don't want to hard code every merge (maybe just tags)
	// but I don't know what to use to merge two like but different structs
	// Might have to write a reflection for this and just accept the pain
	if local.KeypairName != "" {
		a.KeypairName = local.KeypairName
	}
	if local.InstanceType != "" {
		a.InstanceType = local.InstanceType
	}
	if local.IamInstanceProfile != "" {
		a.IamInstanceProfile = local.IamInstanceProfile
	}
	if len(local.Tags) > 0 {
		for _, localTag := range local.Tags {
			a.Tags = append(a.Tags, localTag)
		}
	}
	if local.SubnetID != "" {
		a.SubnetID = local.SubnetID
	}
	if local.Region != "" {
		a.Region = local.Region
	}
	if local.AMI != "" {
		a.AMI = local.AMI
	}
	if len(local.SecurityGroups) > 0 {
		a.SecurityGroups = local.SecurityGroups
	}
	if local.SSHUsername != "" {
		a.SSHUsername = local.SSHUsername
	}
	if local.SSHPrivateKeyPath != "" {
		a.SSHPrivateKeyPath = local.SSHPrivateKeyPath
	}

	return nil

}

// LocalAwsConfiguration defines variables used to launch
// a betwixt AWS instance
type LocalAwsConfiguration struct {
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

// AwsTag defines a tag pair for an ec2 instance
type AwsTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
