# Teleport Tacos

Teleport Tacos is demo app I wrote for Scylla Summit 2023. However, I do have plans for it to be an app used in future demos. Shout out to [Katan Coding](https://t.co/QaN1cAzDmu?amp=1) for the inspiration on the Hex Architecture in Go.

## Building the Images

There are go-task targets for build the images. 

```sh
task build-client && task build-server
```