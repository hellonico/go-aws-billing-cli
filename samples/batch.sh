#!/usr/bin/env bash

ACCOUNTS=${1:-default}
DIMENSIONS="AZ LINKED_ACCOUNT SERVICE LEGAL_ENTITY_NAME INVOICING_ENTITY TENANCY OPERATION USAGE_TYPE RECORD_TYPE"

GRANULARITY=${2:-MONTHLY}
MONTH=${3:-3}

OUTPUTDIR=`echo output/billing_${GRANULARITY}_${MONTH}_months | tr '[:upper:]' '[:lower:]'`
echo $OUTPUTDIR

rm -fr $OUTPUTDIR

for account in ${ACCOUNTS}; do
   for DIMENSION in ${DIMENSIONS}; do 
      echo $account ">" $DIMENSION
      ./awsbillingcli -o $OUTPUTDIR/BYDIMENSION/${DIMENSION}/${account}.csv -a ${account} -gr $GRANULARITY -start $MONTH -g ${DIMENSION} 
      ./awsbillingcli -o $OUTPUTDIR/BYACCOUNT/${account}/${DIMENSION}.csv -a ${account} -gr $GRANULARITY -start $MONTH -g ${DIMENSION} 
   done
done
