# Containers from Scratch

A talk I gave for Meetup Rosenheim about how Linux contains are implemented by
Docker and other container runtimes.
Containers are often compared to virtual machines. In this talk I gave an
overview of when this comparison falls short.

## Use the container script

The container script starts a process in a new "container". Specifically it
uses the folder "rootfs" as a new root directory and sets the UTS, PID and NET
namespaces.
You can run it with: `go run container.go run [COMMAND]`.

E.g. `go run container.go run sh`
