package devops_scripts

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// GetEC2InstanceType returns the type of the specified EC2 instance
func GetEC2InstanceType(ctx context.Context, instanceID string) (string, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")}, nil)
	if err!= nil {
		return "", err
	}

	ec2Client := ec2.New(sess)

	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	}

	output, err := ec2Client.DescribeInstancesWithContext(ctx, input)
	if err!= nil {
		return "", err
	}

	if len(output.Reservations) == 0 {
		return "", fmt.Errorf("instance %s not found", instanceID)
	}

	return *output.Reservations[0].Instances[0].InstanceType, nil
}

// GetProcessID returns the process ID of the specified process
func GetProcessID(processName string) (int, error) {
	processes, err := os.ReadDir("/proc")
	if err!= nil {
		return 0, err
	}

	for _, process := range processes {
		if process.Name() == processName {
			id, err := strconv.Atoi(process.Name())
			if err!= nil {
				return 0, err
			}
			return id, nil
		}
	}

	return 0, fmt.Errorf("process %s not found", processName)
}

// GetSSHKey returns the SSH key of the specified user
func GetSSHKey(username string) (string, error) {
	sshKeys, err := os.ReadDir("/home/" + username + "/.ssh")
	if err!= nil {
		return "", err
	}

	for _, key := range sshKeys {
		if key.Name() == "id_rsa" || key.Name() == "id_ed25519" {
			return "/home/" + username + "/.ssh/" + key.Name(), nil
		}
	}

	return "", fmt.Errorf("SSH key not found for user %s", username)
}

// RetryFunction retries a function up to a specified number of times with a specified delay between retries
func RetryFunction(ctx context.Context, fn func() error, retries int, delay time.Duration) error {
	for i := 0; i < retries; i++ {
		err := fn()
		if err == nil {
			return nil
		}
		if i < retries-1 {
			log.Printf("Error %d/%d: %s, retrying in %s\n", i+1, retries, err, delay)
			time.Sleep(delay)
		}
	}
	return err
}

// PanicHandler catches panics and logs the error
func PanicHandler() {
	if r := recover(); r!= nil {
		log.Println("Recovered in", r)
		debug.PrintStack()
		os.Exit(1)
	}
}

func IsStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if a == b {
			return true
		}
	}
	return false
}

func StringsInSlice(strings []string, list []string) bool {
	for _, str := range list {
		if!IsStringInSlice(str, strings) {
			return false
		}
	}
	return true
}

func StringsNotInSlice(strings []string, list []string) bool {
	for _, str := range list {
		if IsStringInSlice(str, strings) {
			return false
		}
	}
	return true
}

func StringsSubset(strings []string, list []string) bool {
	for _, str := range list {
		if!IsStringInSlice(str, strings) {
			return false
		}
	}
	return true
}

func StringsSuperset(strings []string, list []string) bool {
	for _, str := range list {
		if!IsStringInSlice(str, strings) {
			return false
		}
	}
	return true
}