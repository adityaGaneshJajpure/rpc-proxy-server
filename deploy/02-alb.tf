resource "aws_security_group" "alb" {
  name   = "proxy-alb-sg"
  vpc_id = module.vpc.vpc_id
  description = "Allow inbound traffic from port 80 and 443, to the ALB"
 
  ingress {
   protocol         = "tcp"
   from_port        = 80
   to_port          = 80
   cidr_blocks      = ["0.0.0.0/0"]
   ipv6_cidr_blocks = ["::/0"]
  }
 
  ingress {
   protocol         = "tcp"
   from_port        = 443
   to_port          = 443
   cidr_blocks      = ["0.0.0.0/0"]
   ipv6_cidr_blocks = ["::/0"]
  }
 
  egress {
   protocol         = "-1"
   from_port        = 0
   to_port          = 0
   cidr_blocks      = ["0.0.0.0/0"]
   ipv6_cidr_blocks = ["::/0"]
  }
}

resource "aws_lb" "proxy_lb" {
  name               = "proxy-lb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.alb.id]
  subnets            = module.vpc.public_subnets
 }

resource "aws_alb_target_group" "proxy_target_group" {
  name        = "proxy-tg"
  port        = 80
  protocol    = "HTTP"
  vpc_id      = module.vpc.vpc_id
  target_type = "ip"
 
  health_check {
   healthy_threshold   = "3"
   interval            = "30"
   protocol            = "HTTP"
   matcher             = "200"
   timeout             = "3"
   path                = "/_healthz"
   unhealthy_threshold = "2"
  }
} 

resource "aws_alb_listener" "http" {
  load_balancer_arn = aws_lb.main_lb.id
  port              = 80
  protocol          = "HTTP"
 
  default_action {
    target_group_arn = aws_alb_target_group.proxy_target_group.id
    type             = "forward"
  }
}
