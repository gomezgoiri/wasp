---
keywords:
- ISCP
- Wasp-cli
- Configuration
- Goshimmer
- command line
description: How to configure the wasp-cli. Requirements and configuration parameters.
image: /img/logo/WASP_logo_dark.png
---

# Configuring wasp-cli

Here we describe step by step instructions how to use wasp-cli to interact with Wasp nodes on the Goshimmer network.

## Requirements

After going through the instructions on [Running a node](./running-a-node.md), you should have the wasp-cli binary available in your system.

## Configuration

You can create a basic default configuration by running `wasp-cli init`. This will create a configuration file named `wasp-cli.json` in the current directory.

Then, we need to tell `wasp-cli` the location of the Goshimmer node and the
committee of Wasp nodes:

```shell
$ wasp-cli set goshimmer.api 127.0.0.1:8080

$ wasp-cli set wasp.0.api 127.0.0.1:9090
$ wasp-cli set wasp.0.nanomsg 127.0.0.1:5550
$ wasp-cli set wasp.0.peering 127.0.0.1:4000

$ wasp-cli set wasp.1.api 127.0.0.1:9091
$ wasp-cli set wasp.1.nanomsg 127.0.0.1:5551
$ wasp-cli set wasp.1.peering 127.0.0.1:4001

...
```

Alternatively, you can edit the `wasp-cli.json` file and include the desired server locations:

- the goshimmer api address:

  ```json
    "goshimmer": {
      "api": "127.0.0.1:8080",
      "faucetpowtarget": -1
    },
  ```

- The API/nanomsg/peering address for each Wasp node:

  ```json
  "wasp": {
      "0": {
        "api": "127.0.0.1:9090",
        "nanomsg": "127.0.0.1:5550",
        "peering": "127.0.0.1:4000"
      },
      "1": {
        ...
      },
    }
  ```