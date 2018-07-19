resource "aws_ecs_task_definition" "go-gen-ssl" {
  family                   = "go-gen-ssl"
  container_definitions    = "${file("task-definitions/go-gen-ssl.json")}"
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = "arn:aws:iam::728160576949:role/ECSTaskExecutionRole"
  network_mode             = "awsvpc"
  cpu                      = 256
  memory                   = 512
}
