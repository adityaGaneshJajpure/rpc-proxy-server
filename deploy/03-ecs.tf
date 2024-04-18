resource "aws_ecr_repository" "proxy_repository" {
  name                 = "proxy"
  image_tag_mutability = "MUTABLE"
}

resource "aws_ecs_cluster" "proxy_cluster" {
  name = "proxy-cluster"
}

resource "aws_ecs_task_definition" "proxy_task_definition" {
  family                   = "proxy-task-definition"
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  task_role_arn            = aws_iam_role.ecs_task_role.arn
  network_mode             = "awsvpc"
  cpu                      = 256
  memory                   = 512

  container_definitions = templatefile("containers/task.tpl.json",
    { CONTAINER_PORT              = var.container_port,
      REGION                      = var.region,
      LOG_GROUP                   = aws_cloudwatch_log_group.proxy_service_log_group.name,
      APP_NAME                    = var.app_name,
  })
}

resource "aws_cloudwatch_log_group" "proxy_service_log_group" {
  name = "proxy-service-log"
}

resource "aws_ecs_service" "proxy_service" {
  name                               = "proxy-service"
  cluster                            = aws_ecs_cluster.proxy_cluster.id
  task_definition                    = aws_ecs_task_definition.proxy_task_definition.arn
  desired_count                      = 1
  deployment_minimum_healthy_percent = 50
  deployment_maximum_percent         = 200
  launch_type                        = "FARGATE"
  scheduling_strategy                = "REPLICA"

  network_configuration {
    security_groups  = [aws_security_group.proxy_svc_sg.id]
    subnets          = module.vpc.private_subnets
  }

  load_balancer {
    target_group_arn = aws_alb_target_group.proxy_target_group.arn
    container_name   = var.app_name
    container_port   = var.container_port
  }
}
