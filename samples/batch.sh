#!/usr/bin/env bash

ACCOUNTS="billing nishida"
SERVICES="LINKED_ACCOUNT SERVICE"
OUTPUTDIR=output/
rm -fr $OUTPUTDIR
for account in ${ACCOUNTS}; do
   for service in ${SERVICES}; do 
      mkdir -p output/${service}
      # echo $service
      ./awsbillingcli -o $OUTPUTDIR/${service}/${account}.csv -a ${account} -start 3 -g ${service} 
   done
   # ./awsbillingcli -a ${account} -start 2 -g SERVICE -f "Amazon Elastic Compute Cloud - Compute,EC2 - Other" -o output/${account}
   # ./awsbillingcli -a ${account} -start 2 -g LINKED_ACCOUNT -f "Amazon Elastic Compute Cloud - Compute,EC2 - Other" -o output/${account}
   
done
