esource "aws_vpc" "main" {
  cidr_block           = var.vpc_cidr
  enable_dns_support   = true
  enable_dns_hostnames = true

  ipv6_cidr_block           = var.ipv6_cidr
  assign_generated_ipv6_cidr_block = true

  tags = {
    Name = var.vpc_name
  }
}

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.main.id
  tags = {
    Name = var.igw_name
  }
}

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }
  route {
    ipv6_cidr_block = "::/0"
    gateway_id      = aws_internet_gateway.igw.id
  }
  tags = {
    Name = var.public_rt_name
  }
}

resource "aws_subnet" "public" {
  count = 3

  vpc_id                  = aws_vpc.main.id
  cidr_block              = cidrsubnet(aws_vpc.main.cidr_block, 8, count.index)
  ipv6_cidr_block         = aws_vpc.main.ipv6_cidr_block_association[count.index].ipv6_cidr_block
  map_public_ip_on_launch = true
  availability_zone       = data.aws_availability_zones.available.names[count.index]

  tags = {
    Name = "${var.public_subnet_name_prefix}-${count.index + 1}"
  }
}

resource "aws_subnet" "private" {
  count = 3

  vpc_id            = aws_vpc.main.id
  cidr_block        = cidrsubnet(aws_vpc.main.cidr_block, 8, count.index + 3)
  ipv6_cidr_block   = aws_vpc.main.ipv6_cidr_block_association[count.index + 3].ipv6_cidr_block
  availability_zone = data.aws_availability_zones.available.names[count.index]

  tags = {
    Name = "${var.private_subnet_name_prefix}-${count.index + 1}"
  }
}
