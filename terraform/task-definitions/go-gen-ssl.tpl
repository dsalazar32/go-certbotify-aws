[
  {
    "name": "go-gen-ssl",
    "image": "dsalazar32/go-gen-ssl:latest",
    "cpu": 10,
    "memory": 100,
    "entrypoint": [ "go-gen-ssl" ],
    "environment": [
      {
        "name": "AWS_REGION",
        "value": "${aws_region}"
      },
      {
        "name": "AWS_ECS_CLUSTER_NAME",
        "value": "${aws_ecs_cluster_name}"
      },
      {
        "name": "AWS_ECS_TASK_NAME",
        "value": "go-gen-ssl-${aws_ecs_task_suffix}"
      }
    ],
    "command": [
      "generate",
      "-email",
      "${email}",
      "-d",
      "${domain}",
      "-s3",
      "-auto-renew"
    ],
    "logConfiguration": {
      "logDriver": "awslogs",
      "options": {
        "awslogs-group": "/dsalazar32/go-gen-ssl/logs",
        "awslogs-region": "${aws_region}"
      }
    }
  }
]
