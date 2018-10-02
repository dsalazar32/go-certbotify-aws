data "template_file" "dsalazar_io" {
  template = "${file("task-definitions/go-gen-ssl.tpl")}"
  vars {
    domain = "*.dsalazar.io"
    aws_region = "us-east-1"
    aws_ecs_cluster_name = "Automata"
    aws_ecs_task_suffix = "dsalazar_io"
    email = "david@dsalazar.io"
  }
}

resource "aws_ecs_task_definition" "go-gen-ssl-dsalazar_io" {
  family = "go-gen-ssl-dsalazar_io"
  container_definitions = "${data.template_file.dsalazar_io.rendered}"
  task_role_arn = "arn:aws:iam::${var.aws_account_number}:role/${var.aws_resource_prefix}ECSTaskExecutionRole"
}
