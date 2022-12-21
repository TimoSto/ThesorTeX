locals {
  email = "thesortex.contact@gmail.com"
}

resource "aws_sns_topic" "monitoring_topic" {
  name         = "ThesorTeX-Monitoring"
  display_name = "Monitoring for budget and cloudwatch alarms"
  delivery_policy = jsonencode({
    "http" : {
      "defaultHealthyRetryPolicy" : {
        "minDelayTarget" : 20,
        "maxDelayTarget" : 20,
        "numRetries" : 3,
        "numMaxDelayRetries" : 0,
        "numNoDelayRetries" : 0,
        "numMinDelayRetries" : 0,
        "backoffFunction" : "linear"
      },
      "disableSubscriptionOverrides" : false,
      "defaultThrottlePolicy" : {
        "maxReceivesPerSecond" : 1
      }
    }
  })
}

resource "aws_sns_topic_subscription" "topic_email_subscription" {
  topic_arn = aws_sns_topic.monitoring_topic.arn
  protocol  = "email"
  endpoint  = local.email
  confirmation_timeout_in_minutes = 30
}