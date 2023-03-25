#!/bin/bash


if [ -z "$1" ]; then 
    echo "Error: missing articate package as 1st input"
    echo "Usage: "
    echo "  $0 ARTIFACT_PACKAGE TAG"

    exit 1

fi

if [ ! -f "$1" ] ; then  
   echo "Error: artifcact $1 does not exists"
   exit 1

fi

artifcat_path=$1
artifact=$(basename $artifcat_path)

if [ -z "$2" ]; then
    echo "Error: missing tag  as 2nd input"
    echo "Usage: "
    echo "  $0 $1  TAG"

    exit 1

fi

TAG=$2


echo "Checking Signature for version: ${TAG}"
cosign verify-blob \
  --certificate "https://github.com/vdbulcke/vault-token-monitor/releases/download/${TAG}/${artifact}.pem" \
  --signature "https://github.com/vdbulcke/vault-token-monitor/releases/download/${TAG}/${artifact}.sig"  \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com  \
  --certificate-identity  "https://github.com/vdbulcke/vault-token-monitor/.github/workflows/release.yaml@refs/tags/${TAG}"  \
  ${artifcat_path}