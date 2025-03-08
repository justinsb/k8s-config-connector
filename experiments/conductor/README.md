mkdir -p ~/kccai

pushd ~/kccai
git clone https://github.com/GoogleCloudPlatform/k8s-config-connector.git
popd


RUNNER="go run . runner --branch-conf=branches-all.yaml --logging-dir ${HOME}/kccai/logs --for-resources-regex=compute.* --branch-repo=${HOME}/kccai/k8s-config-connector"

```
# 1 - [Validate] Repo directory and metadata
${RUNNER} --command=1
```


```
# 2 - [Branch] Create the local github branches from the metadata
${RUNNER} --command=2
```

```
# 3 - [Branch] Delete the local github branches from the metadata
${RUNNER} --command=3
```

```
# 4 - [Project] Enable GCP APIs for each branch
${RUNNER} --command=4
```

```
# 5 - [Generated] Read the specific type of generated files in each github branch
${RUNNER} --command=5
```

```
# 6 - [Generated] Write the specific type of files from all_scripts.yaml to each github branch
${RUNNER} --command=6
```

```
# 10 - [Mock] Create script.yaml for mock gcp generation in each github branch
${RUNNER} --command=10
```


2025/03/08 11:03:53     4 - [Project] Enable GCP APIs for each branch
2025/03/08 11:03:53     5 - [Generated] Read the specific type of generated files in each github branch
2025/03/08 11:03:53     6 - [Generated] Write the specific type of files from all_scripts.yaml to each github branch
2025/03/08 11:03:53     10 - [Mock] Create script.yaml for mock gcp generation in each github branch
2025/03/08 11:03:53     11 - [Mock] Create _http.log for mock gcp generation in each github branch
2025/03/08 11:03:53     12 - [Mock] Generate mock Service and Resource go files in each github branch
2025/03/08 11:03:53     13 - [Mock] Add service to mock_http_roundtrip.go in each github branch
2025/03/08 11:03:53     14 - [Mock] Add proto to makefile in each github branch
2025/03/08 11:03:53     15 - [Mock] Run mockgcptests on generated mocks in each github branch
2025/03/08 11:03:53     20 - [CRD] Generate Types and Mapper for each branch
2025/03/08 11:03:53     21 - [CRD] Generate CRD for each branch
2025/03/08 11:03:53     22 - [Fuzzer] Generate fuzzer for each branch -->

