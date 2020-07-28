# Pepper

## Compilation

To build, test and compile a cloud-native package using Docker, run [`make image`](https://gitlab.com/sorleone/pepper/-/blob/master/Makefile#L3).

## Execution

To run the image locally run [`make run`](https://gitlab.com/sorleone/pepper/-/blob/master/Makefile#L6). Otherwise you can hit with a request (see [`req.sh`](https://gitlab.com/sorleone/pepper/-/blob/master/req.sh)) the automated deployment endpoint https://pepper-sorleone.cloud.okteto.net/receipt, updated for each Git push by the CI/CD pipeline. In case you'll find the service to be unreachable, it will be due to an automated sleep functionality of the free Kubernetes service I am using, let me know and I will restart it for you.
