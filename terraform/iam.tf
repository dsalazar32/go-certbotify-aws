resource "aws_iam_role" "certbot-cloudwatchevents-role" {
  name = "${var.aws_resource_prefix}ECSEventsRole"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "eventRolePolicy",
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": {
        "Service": "events.amazonaws.com"
      }
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "certbot-execution-policy-attach" {
  role = "${aws_iam_role.certbot-cloudwatchevents-role.name}"
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceEventsRole"
}

resource "aws_iam_role" "certbot-ecs-taskexecution-role" {
  name = "${var.aws_resource_prefix}ECSTaskExecutionRole"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "taskExecutionRolePolicy",
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": {
        "Service": "ecs-tasks.amazonaws.com"
      }
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "certbot-ecs-taskexecution-policy" {
  role = "${aws_iam_role.certbot-ecs-taskexecution-role.id}"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "taskExecutionPolicy",
      "Action": [
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:ecs:${var.aws_region}:${var.aws_account_number}:task-definition/${var.aws_task_definition}"
    },
    {
      "Sid": "eventRolePolicy",
      "Action": [
        "iam:GetRole",
        "iam:PassRole"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:iam::${var.aws_account_number}:role/${var.aws_resource_prefix}ECSEventsRole"
    },
    {
      "Sid": "restrictActionsToNamespace",
      "Action": [
        "s3:*"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::${var.aws_resource_prefix}*"
    },
    {
      "Sid": "getReadAccessToRoute53Resource",
      "Effect": "Allow",
      "Action": [
          "route53:ListHostedZones",
          "route53:GetChange"
      ],
      "Resource": [
          "*"
      ]
    },
    {
      "Sid": "restrictAccessToHostedZone",
      "Effect" : "Allow",
      "Action" : [
          "route53:ChangeResourceRecordSets"
      ],
      "Resource" : [
          "arn:aws:route53:::hostedzone/Z3NZSDRMZFX0NL",
          "arn:aws:route53:::hostedzone/ZZRVAMU29NNB1"
      ]
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "certbot-ecs-cloudwatch-policy-attach" {
  role = "${aws_iam_role.certbot-ecs-taskexecution-role.name}"
  policy_arn = "arn:aws:iam::aws:policy/CloudWatchEventsFullAccess"
}
