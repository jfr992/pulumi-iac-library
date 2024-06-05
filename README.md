## Pulumi POC with github actions

## Description

This Pulumi project sets up a network, an Application Load Balancer (ALB), and an Auto Scaling Group (ASG) using custom Pulumi packages. Check the architecture diagram

## Project Structure

Network: Creates a Virtual Private Cloud (VPC) along with public and private subnets.
ALB: Creates an Application Load Balancer and a target group.
ASG: Creates an Auto Scaling Group with instances that are registered to the target group.

### Prerequisites

Before running this project, ensure you have the following installed:

Go
Pulumi
AWS CLI

### Architecture of webServer deployment:

![diagram](POC.drawio.png "diagram")

## Configuration 

The infrastructure configuration is managed through the infra.yaml file, and the user data script for the ASG instances is provided in the userdata.sh file.
```
vpc:
  name: vpc
  cidr_block: <vpc-cidr>

subnets:
  - cidr_block: <public-cidr>
    az: us-east-1a
    public: true

  - cidr_block: <private-cidr>
    az: us-east-1a
    public: false

  - cidr_block: <private-cidr>
    az: us-east-1b
    public: false

alb:
  port: 80
  allowed_cidrs:
    - 0.0.0.0/0

asg:
  name: "some-name"
  ami-id: "ami-123"
  instance-type: "t2.micro"
  min-size: 1
  max-size: 1
  desired-capacity: 1
  ports:
    - 80
```

### Running the Project locally


* Log-ing to aws console and create an IAM user with enough priviledges.
* Gather the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY and set them as env variables 
```
export AWS_ACCESS_KEY_ID="<AWS_ACCESS_KEY_ID>"
export AWS_SECRET_ACCESS_KEY="<AWS_SECRET_ACCESS_KEY>"

git clone https://github.com/yourusername/your-repo-name.git
cd your-repo-name

go mod tidy

pulumi up


```
### Outputs

The following outputs will be exported:

vpcID: The ID of the VPC created.
privateSubnetIds: The IDs of the private subnets created.
publicSubnetIds: The IDs of the public subnets created.
targetGroupArn: The ARN of the target group created.
securityGroupID: The ID of the security group created.

### GitHub Actions
This project uses GitHub Actions to preview the infrastructure changes on pull requests and updates the infrastructure when is merged . Check pulumi-preview.yml and pulumi-up.yml

## Author

Juan Felipe Reyes Marlés  