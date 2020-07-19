# Pepper

## Compilation

To build, test and compile a cloud-native package using Docker, run [`make image`](https://gitlab.com/sorleone/pepper/-/blob/master/Makefile#L1).

## Execution

To run the image locally run [`make run`](https://gitlab.com/sorleone/pepper/-/blob/master/Makefile#L4). Otherwise you can hit with a request (see [`req.sh`](https://gitlab.com/sorleone/pepper/-/blob/master/req.sh)) the automated deployment endpoint https://pepper-sorleone.cloud.okteto.net/receipt, updated for each Git push by the CI/CD pipeline.
