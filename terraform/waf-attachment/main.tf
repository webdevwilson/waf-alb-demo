// Query for Web ACL ARN
data "aws_wafregional_web_acl" "web_acl" {
  name = "fortinet-owasp-top10"
}

// Create an ALB
resource "aws_alb" "my_alb" {
  name = "waf-alb"
}

// Attach ALB to WAF
resource "aws_wafregional_web_acl_association" "waf_association" {
  web_acl_id   = data.aws_wafregional_web_acl.web_acl.id
  resource_arn = aws_alb.my_alb.arn
}