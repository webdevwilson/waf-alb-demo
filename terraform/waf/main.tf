data "aws_wafregional_web_acl" "waf_rule" {

}

resource aws_wafregional_web_acl my_waf {
  name        = "fortinet-owasp-top10"
  metric_name = "fortinetowasptop10"

  default_action {
    type = "ALLOW"
  }

  rule {
    priority = 1
    rule_id  = "7474b078-6abd-408a-a2e4-d70e242b610a"

    override_action {
      type = "NONE"
    }

    type = "GROUP"
  }
}
