resource "aws_ecs_task_definition" "go-gen-ssl" {
  family = "go-gen-ssl"
  container_definitions = "${file("task-definitions/go-gen-ssl.json")}"
  task_role_arn = "arn:aws:iam::${var.aws_account_number}:role/${var.aws_resource_prefix}ECSTaskExecutionRole"
}
