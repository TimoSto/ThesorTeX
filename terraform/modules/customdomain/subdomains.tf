resource "aws_route53_zone" "subdomain" {
  count = length(var.subdomains)
  name  = "${var.subdomains[count.index]}.${var.domain_name}"
}

resource "aws_route53_record" "ns_subdomain" {
  count   = length(var.subdomains)
  name    = aws_route53_zone.subdomain[count.index].name
  records = aws_route53_zone.subdomain[count.index].name_servers
  ttl     = 60
  type    = "NS"
  zone_id = data.aws_route53_zone.public.zone_id
}
