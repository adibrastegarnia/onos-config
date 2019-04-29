# gNOI Device simulator

This is a docker VM that runs gNOI implementation supporting the following services:

* *Certificate management* that can be used for certificate installation and rotation. 

## Create the docker image
```bash
docker build -t gnoisim -f Dockerfile .
```

## Run 3 devices with Docker Compose
Run 3 devices with 
```bash
docker-compose up
```

This gives an output like:
```bash
Creating gnoisim_devicesim2_1 ... 
Creating gnoisim_devicesim3_1 ... 
Creating gnoisim_devicesim1_1 ... 
Creating gnoisim_devicesim2_1
Creating gnoisim_devicesim3_1
Creating gnoisim_devicesim1_1 ... done
Attaching to gnoisim_devicesim2_1, gnoisim_devicesim3_1, gnoisim_devicesim1_1
devicesim3_1  | gNOI running on localhost:50004
devicesim2_1  | gNOI running on localhost:50003
devicesim3_1  | I0429 14:36:58.565531      10 gnoi_target.go:66] No credentials, setting Bootstrapping state.
devicesim2_1  | I0429 14:36:58.380245      11 gnoi_target.go:66] No credentials, setting Bootstrapping state.
devicesim3_1  | I0429 14:36:58.566199      10 gnoi_target.go:48] Starting gNOI server.
devicesim1_1  | gNOI running on localhost:50002
devicesim2_1  | I0429 14:36:58.380885      11 gnoi_target.go:48] Starting gNOI server.
devicesim1_1  | I0429 14:37:00.703256      10 gnoi_target.go:66] No credentials, setting Bootstrapping state.
devicesim1_1  | I0429 14:37:00.704000      10 gnoi_target.go:48] Starting gNOI server.
```

> This uses port mapping to make the devices available to gNOI clients and is the
> only option for running on Mac or Windows, and can be used with Linux too.

### Running on Linux
If you are fortunate enough to be using Docker on Linux, then you can use the
above method __or__ the alternative:

```bash
docker-compose -f docker-compose-linux.yml up
```

This will use the fixed IP addresses 172.25.0.11, 172.25.0.12, 172.25.0.13 for
device1-3. An entry must still be placed in your /etc/hosts file for all 3 like:
```bash
172.25.0.11 device1.opennetworking.org
172.25.0.12 device2.opennetworking.org
172.25.0.13 device3.opennetworking.org
```

> This uses a custom network 'simnet' in Docker and is only possible on Docker for Linux.
> If you are on Mac or Windows it is __not possible__ to route to User Defined networks,
> so the port mapping technique must be used.

> It is not possible to use the name mapping of the docker network from outside
> the cluster, so either the entries have to be placed in /etc/hosts or on some
> DNS server


## Run a single docker container
If you just want to run a single device, it is not necessary to run docker-compose.
It can be done just by docker directly, and can be handy for troubleshooting.
```bash
docker run --env "GNOI_TARGET=localhost" --env "GNOI_PORT=50002" -p "50002:50002" gnoisim
```
To stop it use "docker kill"


