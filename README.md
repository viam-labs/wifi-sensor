# wifi-sensor

A sensor that detects wifi strength

## Usage

### 1. Build binary

If you clone this repository to the target environment where you run your Viam robot, then you can build a binary named `wifi` with:

```
go build -o wifi
```

Alternatively, if you want to build this a binary for a different target environment, please use the [viam canon tool](https://github.com/viamrobotics/canon).

### 2. Add to robot configuration

Copy the binary to in the environment where your Viam robot is running and add the following to your configuration:

```
  ...
  "modules": [
    ...,
    {
      "executable_path": "<path_to_binary>",
      "name": "wifi"
    }
    ...,
  ],
  ...
```

For more information on how to configure modular components, [see this example](https://docs.viam.com/services/slam/run-slam-cartographer/#step-1-add-your-rdiplar-as-a-modular-component).
