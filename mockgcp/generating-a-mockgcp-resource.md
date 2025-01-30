
1. Define some variables that describe the command and API you want to mock:

```
export REPO_ROOT="$(git rev-parse --show-toplevel)"
export GCLOUD_COMMAND="gcloud workflows"
export API="workflows"
export RESOURCE="instance"
```
   
1. Ask codebot to generate the mockgcp test script:

```
cd ${REPO_ROOT}/mockgcp/
cat ${REPO_ROOT}/dev/tools/controllerbuilder/cmd/codebot/examples/mockgcp-test-case/1-create-test.md | \
    envsubst 'GCLOUD_COMMAND=${GCLOUD_COMMAND}' | \
    codebot
```

1. Ask codebot to run the script and capture the output:

```
cd ${REPO_ROOT}/mockgcp/
# Hack to make sure the gcloud access token is current
gcloud auth print-access-token > /dev/null
cat ${REPO_ROOT}/dev/tools/controllerbuilder/cmd/codebot/examples/mockgcp-test-case/2-record-realgcp-output.md | \
    envsubst 'API=${API};RESOURCE=${RESOURCE}' | \
    codebot
```

TODO: We assume that we created `mock${API}/testdata/${RESOURCE}/crud/script.yaml`; it would be much better to
have step 1 output the path in a structured way, and then feed that value into future steps (somehow!)
