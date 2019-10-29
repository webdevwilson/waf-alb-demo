data "aws_wafregional_web_acl" "web_acl" {
  name = var.waf_rule_set
}

// create an ALB
resource "aws_alb" "my_alb" {
  name = "waf-alb"
}

// attach waf to alb
resource "aws_wafregional_web_acl_association" "waf_association" {
  web_acl_id   = data.aws_wafregional_web_acl.web_acl.id
  resource_arn = aws_alb.my_alb.arn
}