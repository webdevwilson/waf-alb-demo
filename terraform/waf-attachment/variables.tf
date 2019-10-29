variable "waf_rule_set" {
  type        = "string"
  default     = "fortinet-owasp-top10"
  description = "WAF Rules to use"
}