This code will take a snapshot with raspistill and upload it to dropbox.

You will need a config file with some parameters:

    package main

    const (
    		clientid = "..."
    		clientsecret = "..."
    		token = "..."
    )

Those parameters are given by dropbox when you register an app.

Then you can run this with

    go run config.go upload.go

### Install go on Raspberry PI

1. Download the latest version from [http://dave.cheney.net/unofficial-arm-tarballs] with wget

2. Move that file to /usr/local (use sudo mv)

3. Untar it with sudo tar -xzf

4. Create the file /etc/profile.d/go.sh and put the following lines in it:

    export PATH=/usr/local/go/bin:$PATH
    export GOPATH="$HOME/your-workspace-dir/"
