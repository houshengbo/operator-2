# Knative Operator Development

Most of the same tools required for
[Knative Serving development](https://github.com/knative/serving/blob/master/DEVELOPMENT.md)
are required for the operator, too.

You'll probably need to
[install Istio](https://istio.io/latest/docs/setup/getting-started/). For the e2e
tests to pass, you only need its CRD's.

All the integration/upgrade tests are now written for Istio, though we are seeking for the extension in the future.

You should clone this repo to `$GOPATH/src/knative.dev/operator`. All commands
below are relative to that path.

## To install the operator:

```
ko apply -f config/
```

To verify the installation:

```
kubectl get deployment
```

You will get the following result, indicating your installation is fine:

```
NAME               READY   UP-TO-DATE   AVAILABLE   AGE
knative-operator   1/1     1            1           17m
```

You can check the log via the command:

```
kubectl logs -f deploy/knative-operator
```

## To run the unit tests:

```
go test -v ./...
```

## To run the integration tests:

First, install the Knative Operator. The integration tests use two environment variables: `TEST_NAMESPACE` for Knative
Serving tests and `TEST_EVENTING_NAMESPACE` for Knative Eventing tests.

```
export TEST_NAMESPACE=knative-serving
export TEST_EVENTING_NAMESPACE=knative-eventing
```

You can choose any names, but the Knative Serving and Knative Eventing should have different namespaces.

Create the namespaces:

```
kubectl create namespace $TEST_NAMESPACE
kubectl create namespace $TEST_EVENTING_NAMESPACE
```

All the integration tests are tagged with `e2e`. Run the integration tests for Knative Serving and Eventing:

```
go test -v -tags=e2e -count=1 ./test/e2e
```

You should get all the tests passed, which means your installation is successful. If you run into any issues, log your
issues [here](https://github.com/knative/operator/issues).

## To run the upgrade tests:

The upgrade tests have taken everything into account. You do not even need to install Knative Operator or Istio before
running the tests. Make sure you have set two of these environment variables: `TEST_NAMESPACE` and `TEST_EVENTING_NAMESPACE`.

```
export TEST_NAMESPACE=knative-serving
export TEST_EVENTING_NAMESPACE=knative-eventing
```

IMPORTANT: Before running ANY upgrade tests, run the following commands to clean up the existing resources:

```
kubectl delete KnativeServing --all -n $TEST_NAMESPACE
kubectl delete KnativeEventing --all -n $TEST_EVENTING_NAMESPACE
```

Run the upgrade tests for the Serving and Eventing CRs:

```
go test -v -tags=upgrade -count=1 ./test/upgrade
```
