# ARS - AWS Role Switcher

## Requirements

1. Have AWS credentials file located in `~/.aws/credentials`
2. MacOS or Linux (only tested on MacOS)
3. Use Google Chrome, Firefox, Safari or Brave Browser

## How to setup AWS assume roles

### 1. Create IAM policy on AWS console (on AWS account you want to access)

Login to the AWS account you want to access (not your origin account), then create IAM policy like below:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::<your-origin-aws-account-id>:root"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
```

### 2. Add IAM policy on AWS account (on your origin AWS account)

Login to your origin account, then add IAM policy to the user you want to give permission to assume the above role like below:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "sts:AssumeRole",
      "Resource": "arn:aws:iam::<aws-account-id-you-want-to-access>:role/<role-name>"
    }
  ]
}
```

### 3. Update AWS credentials file

Update `.aws/credentials` file

```ini
[my-aws-profile]
aws_access_key_id = XXXX
aws_secret_access_key = XXXX

[my-role]
role_arn = arn:aws:iam::<role-account-id>:role/<role-name>
source_profile = my-aws-profile
```

## Installation

1. Download the binary
    - OSX ? Download bin/osx/ars.tar.gz
    - Linux ? Download bin/linux/ars.tar.gz
2. Extract the file. **Tip**: If you're using terminal -> run `tar -xzvf ars.tar.gz`
3. Copy the extracted file to `usr/local/bin`. **Tip**: If you're using terminal -> run `cp ars /usr/local/bin`

## Usage

```bash
$ ars --help
Usage of ars:
  -ask-redirect-url
      Prompt redirect URL
  -browser string
      Browser to open AWS Console. Accepted values: firefox, chrome, safari, brave (default "chrome")
  -redirect-url string
      Redirect URL after login (default "https://<profile_region>.console.aws.amazon.com")
```
