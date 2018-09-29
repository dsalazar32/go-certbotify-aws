# TODO: Work on task execution role
# Follow manual setup in `ECSTaskExecutionRole`
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

resource "aws_iam_role_policy_attachment" "certbot-execution-policy-attach" {
  role       = "${aws_iam_role.certbot-cloudwatchevents-role.name}"
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceEventsRole"
}