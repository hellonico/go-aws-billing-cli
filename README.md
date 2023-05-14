Easy AWS Cost Query (CLI) in Go

... because the (non-existing) official doc was somehow a bit hard to figure out.

# Usage

```bash
Usage of ./awsbillingcli:
  -a string
      aws profile name. can specify multiple. if not specify try to value from env variable AWS_PROFILE.
  -start string
      start date. can be either a date 2023-01-01, or an integer value for a number of month.
  -end string
      start date. can be either a date 2023-02-01, or an integer value for a number of month.
  -ft string
      Dimension to filter by. (Ex: SERVICE, LINKED_ACCOUNT etc... (default "SERVICE")
  -f string
      Filter by services. Use , to separate; one or more of (non-exhaustive):
      AWS CloudTrail
      AWS Config
      AWS Cost Explorer
      AWS Directory Service
      AWS Glue
      AWS Key Management Service
      AWS Lambda
      AWS Step Functions
      AWS Support (Developer)
      Amazon Chime
      Amazon Chime Dialin
      Amazon EC2 Container Registry (ECR)
      EC2 - Other
      Amazon Elastic Compute Cloud - Compute
      Amazon Elastic Load Balancing
      Amazon GuardDuty
      Amazon MQ
      Amazon Registrar
      Amazon Relational Database Service
      Amazon Route 53
      Amazon Simple Notification Service
      Amazon Simple Queue Service
      Amazon Simple Storage Service
      Amazon Virtual Private Cloud
      AmazonCloudWatch
  -g string
      group by dimension, one of:
      AZ
      INSTANCE_TYPE
      LEGAL_ENTITY_NAME
      INVOICING_ENTITY
      LINKED_ACCOUNT
      OPERATION
      PLATFORM
      PURCHASE_TYPE
      SERVICE
      TENANCY
      RECORD_TYPE
      USAGE_TYPE
       (default "LINKED_ACCOUNT")
  -gr string
      granularity. one of:
      MONTHLY
      DAILY
      YEARLY
       (default "MONTHLY")
  -m string
      Metrics. One or more of:
      AmortizedCost
      BlendedCost
      NetAmortizedCost
      NetUnblendedCost
      NormalizedUsageAmount
      UnblendedCost
      UsageQuantity
       (default "UnblendedCost")
  -o string
      output. empty for console output, or file path
  -help
      print usage
```

@2023 - hellonico at gmail dot com