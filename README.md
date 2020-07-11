# Grower

- Raspberry Pi
- Prometheus
- Kubernetes
- Event Sourcing


## Initialize the repository


1. Run the `init.sh` file to download the dependencies 
> If you are using the devcontainer setup you probably can skip this step.
> However, it may be a solution if you have problems generating code from protos
> if the error is related to generation of messages validation functions

```console
make init
```