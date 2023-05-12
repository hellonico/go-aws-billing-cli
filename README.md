Easy AWS Cost Query (CLI) in Go

... because the (non-existing) official doc was somehow a bit hard to figure out.

# Usage

```bash
Usage of ./awsbillingcli:
  -a string
    	aws profile name (default "default")
  -end string
    	end date
  -f string
    	Filter by services. Use , to separate; one of Amazon Route 53, AmazonCloudWatch, Amazon Route 53...
  -g string
    	group by dimension, one of: AZ, INSTANCE_TYPE, LEGAL_ENTITY_NAME, INVOICING_ENTITY, LINKED_ACCOUNT, OPERATION, PLATFORM, PURCHASE_TYPE, SERVICE, TENANCY, RECORD_TYPE, and USAGE_TYPE (default "LINKED_ACCOUNT")
  -help
    	print usage
  -m int
    	how many months back in time
  -metrics string
    	Metrics. Default metric is: UnblendedCost (default "UnblendedCost")
  -start string
    	start date. if this is set, month is ignored
```

@2023 - hellonico at gmail dot com