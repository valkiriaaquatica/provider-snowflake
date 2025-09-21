# Provider SNOWFLAKE

`provider-snowflake` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
SNOWFLAKE API.

## KNOW ISSUES
- As long you can check in the provider docs https://registry.terraform.io/providers/snowflakedb/snowflake/latest/docs, there is a /Preview and a /Stable drop-down menu. This Crossplane provider is just focused on the /Stable one, so the resources that exist in /Preview will *not* be handle.

## Installation (make sure you have Crossplane before installed in your cluster)

- Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/valkiriaaquatica/provider-snowflake):
```
up ctp provider install xpkg.upbound.io/valkiriaaquaticamendi/provider-snowflake:v0.1.0
```

- Declarative installation 
  ```
  cat <<EOF | kubectl apply -f -
  apiVersion: pkg.crossplane.io/v1
  kind: Provider
  metadata:
    name: provider-snowflake
  spec:
    package: xpkg.upbound.io/valkiriaaquaticamendi/provider-snowflake:v0.1.0
  EOF
  ```
  or
  ```
  kubectl apply -f examples/install.yaml
  ```
  Now create the seecret with your Snowflake credentials, filling the secret and apply it
  ```
  vi examples/providerconfig/secret.yaml.tmpl
  kubectl apply -f examples/providerconfig/secret.yaml.tmpl
  ```
  Then create the Provider configuration using that secret
  ```
  kubectl apply -f examples/providerconfig/providerconfig.yaml
  ```
  In the folder examples/ and examples-generated/ you can have multiple examples to quick create. If you have any interesting example to add, feel free to contribute. examples/ folder is based on more testes examples while the examples-generated/ wrap the examples from Terraform docs  into Yamls.

  Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

  You can see the API reference [here](https://doc.crds.dev/github.com/valkiriaaquatica/provider-snowflake).

## Developing

### (Optional) -> Intitializate the devbox environment
  If you have devbox or want to work with it, it makes life easier for packages like go, do the following:
  ```console
  cd devbox/
  devbox install
  devbox shell
  ```

  1. Run the generator

    ```bash
    make generate
    ```
  2. Run the image  against an existant Kubernetes cluster that has Crossplane already installed
    Install the CRDs in the cluster:
      ```
      kubectl apply -f package/crds/
      ```
  3. Run the image  against an existant Kubernetes cluster that has Crossplane already installed
      and test it works well
    ```bash
    make run
    ```
  4. Run the tests
      and test it works well
    ```bash
    make test
    ```
  4. Run the local docker build image
      and test it works well
    ```bash
    make build
    ```
  ---ild
  ```


## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/valkiriaaquatica/provider-snowflake/issues).

