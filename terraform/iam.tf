resource "aws_iam_role_policy" "certbot-cloudwatchevents-policy" {
  name = "ECSEventsPolicy"
  role = "${aws_iam_role.certbot-cloudwatchevents-role.id}"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ecs:RunTask"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Effect": "Allow",
            "Action": "iam:PassRole",
            "Resource": [
                "arn:aws:iam::728160576949:role/ECSTaskExecutionRole"
            ]
        }
    ]
}
EOF
}

resource "aws_iam_role" "certbot-cloudwatchevents-role" {
  name = "ECSEventsRole"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
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