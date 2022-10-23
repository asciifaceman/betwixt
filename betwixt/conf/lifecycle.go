package conf

type LifecycleConfig struct {
	AWS *AwsConfiguration
}

func NewLifecycleConfig(global *Global, local *Local) (*LifecycleConfig, error) {
	l := &LifecycleConfig{}

	// AWS
	l.AWS = global.AWS
	err := l.AWS.MergeLocal(local.AWS)
	if err != nil {
		return nil, err
	}

	return l, nil

}
