module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "3.0.0"

  name = "sandbox"
  cidr = var.vpc_cidr

  azs  = var.vpc_azs
  private_subnets = var.private_subnets
  public_subnets = var.public_subnets

  create_igw             = true
  enable_nat_gateway     = true
  create_egress_only_igw = true
  single_nat_gateway     = true
}

resource "aws_security_group" "proxy_svc_sg" {
  name   = "proxy-service-sg"
  vpc_id = module.vpc.vpc_id
  description = "Allow inbound access from the ALB only"
 
  ingress {
   protocol         = "tcp"
   from_port        = 80
   to_port          = 80
   security_groups = [aws_security_group.alb.id]
  }

  ingress {
   protocol         = "tcp"
   from_port        = 3000
   to_port          = 3000
   security_groups = [aws_security_group.alb.id]
  }
 
  egress {
   protocol         = "-1"
   from_port        = 0
   to_port          = 0
   cidr_blocks      = ["0.0.0.0/0"]
  }
}
