# Ginkgo Experiments

## Example

1. Run the `eventually` experiments using one node,

    ```shell
    GINKGO_CMD="ginkgo -v --nodes=1 eventually" make
    ```

1. Run the `teamcity` experiments using two nodes, enabling both teamcity and junit ginkgo reporters.

    ```shell
    GINKGO_REPORTERS="junit,teamcity" GINKGO_CMD="ginkgo -v --nodes=2 teamcity" make
    ```
