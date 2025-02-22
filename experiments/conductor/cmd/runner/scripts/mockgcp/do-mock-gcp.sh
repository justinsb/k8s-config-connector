#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

export WORKDIR=~/kccai/work1/
export BRANCH_NAME=resource-dataproc-clusters
export LOG_DIR=/tmp/conductor/resource-dataproc-clusters/

export GCLOUD_COMMAND="gcloud dataproc clusters"


./01-generate-script.sh

export RUN_TEST=mockdataproc/testdata/cluster/crud

./02-run-script-real-gcp.sh 

export PROTO_PACKAGE=google.cloud.dataproc.v1


./03a-add-to-makefile.sh

export SERVICE=dataproc
export RESOURCE=cluster
export PROTO_SERVICE=google.cloud.dataproc.v1.ClusterController
export PROTO_MESSAGE=google.cloud.dataproc.v1.Cluster
export HTTP_HOST=dataproc.googleapis.com

./03-implement-mocks.sh


./04-run-script-mockgcp.sh
