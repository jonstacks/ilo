# ILO

A library and command line tool written in golang for interacting with
Integrated Lights Out(iLO) devices.

## Introduction

This is a port of my existing python library
[ilo-utils](https://github.com/jonstacks/ilo-utils) to golang so that:

1. Get familiar with golang.
2. Reduce the need for runtime libraries and having python versions installed.
   This is important because those running the command might not be used to
   installing and configuring Python. Distributing as a binary is probably a
   better strategy.
3. Possibly provide better performance. Performance is pretty good right now,
   this is just one of those things that would be extra nice.

## Porting

In order to port over the existing cli, I will be doing the following:

1. Writing an iLO test server in go. This will allow me to write more go code,
   and to have something that will behave like an iLO server.
2. Verify that the existing CLI works against this test server by spinning up
   several VMs and having them run these test servers.
3. Write the new `ilo` package and `ilo-sweep` command in go.

### Test Server

Since I no longer have access to a Server with ILO. This project contains a test
server. This test server allows us to test the `ilo-sweep` command by running
the test server on multiple VMS. This server will also listen on port 17988 so
we can test if that port is open from a port sweeping capability. A Vagrantfile
has also been included so that we can easily spin up 3 test boxes with `vagrant
up`. You can then ssh into each of the boxes and start the ilo-server with:

    ```sh
    cd /vagrant/cmd/ilo-server/ && sudo ./ilo-server
    ```
