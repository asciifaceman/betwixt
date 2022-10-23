package lifecycle

import (
	"fmt"

	"github.com/asciifaceman/betwixt/betwixt/conf"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type AWSRemote struct {
	Config *conf.AwsConfiguration
}

func NewAWSRemote() *AWSRemote {
	return &AWSRemote{}
}

func (a *AWSRemote) LoadConfig(conf *conf.LifecycleConfig) error {
	a.Config = conf.AWS
	return nil
}

func (a *AWSRemote) ec2() (*ec2.EC2, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(a.Config.Region),
	})
	if err != nil {
		return nil, err
	}

	return ec2.New(sess), nil
}

func (a *AWSRemote) Launch() error {
	fmt.Printf("Launching instance in %s\n", a.Config.Region)
	fmt.Println("not really this is just a test")
	//ec2, err := a.ec2()
	//if err != nil {
	//	return err
	//}
	return nil
}

func (a *AWSRemote) Destroy() error {
	return nil
}
