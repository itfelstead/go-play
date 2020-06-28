# go-play

Just a personal area for playing with go.

# Workflow Approach:

For now, for local dev, just mount the source dir from the git repo into a golang container for building and running.  
Edit source and run either inside the container (vi) or on host with Visual Studio Code via its `Remote Development / ms-vscode-remote.vscode-remote-extensionpack`feature.  
Alternatively, do simple experiments in the online playground at play.golang.org.

## Using VS Code with the dev container:

### Option 1) VS Code Remote Containers (WSL2)

*TODO*  
*Reminder: When PC finally gets Windows v2004, try this: https://code.visualstudio.com/docs/remote/containers*

**pros:**
- *tbd*

**cons:**
- *tbd*

### Option 2a) VS Code 'Attach to running container...' (custom container, no devcontainer.json)

The idea is to allow my local VS Code to connect to my development container.  
To do this;

- Enable the `Remote Development / ms-vscode-remote.vscode-remote-extensionpack` extension.
- Run the dev container
- Enable the go-extension
- Install the tools that the go extension required.
    This could be done within VS Code while connected to the container, but I've added the tools to the Dockerfile so they will already be present.
```    
        go get -u github.com/ramya-rao-a/go-outline
        go get -u github.com/uudashr/gopkgs/cmd/gopkgs
        go get -u github.com/stamblerre/gocode
        go get -u github.com/rogpeppe/godef
        go get -u github.com/sqs/goreturns
        go get -u golang.org/x/lint/golint
```
- Open a remote window (click the green "><" in the bottom left of your host's VS Code window), select `Remote-Containers: Attach to running container..`, select your running container.

**pros:**
- it is workable

**cons:**
- Issues with approacgithub.com/stamblerre/gocode; VS Code insists it isn't there.
- Pre-installing VS Code server and extensions is not that well supported (https://github.com/microsoft/vscode-remote-release/issues/1718)


### Option 2b) VS Code 'Attach to running container...' (VS Code generated container)

*TODO - look into devcontainer.json in general too: https://code.visualstudio.com/docs/remote/containers#_creating-a-devcontainerjson-file*

**pros:**
- *tbd*

**cons:**
- *tbd*

## Golang Dev Container: based on https://hub.docker.com/_/golang
- installs GO
- installs the go tools required by the VS Code extention
- sets up env;
    - sets $GOPATH (/go) (i.e. where go projects are to be located)
    - adds $GOPATH/bin to our $PATH

```
    cd <repo which holds the golangdev Dockerfile>
    docker build . --tag golangdev:latest

    (optional i want it in the remote repo: https://hub.docker.com/u/itfelstead) 
    docker tag golangdev:latest itfelstead/golangdev:latest
    docker push itfelstead/golangdev:latest

    cd <source dir of interest>
    docker run -it --rm -v ${PWD}:/go/src/app --name gotest golangdev
        $ go run helloworld.go
        or
        $ go build helloworld.go
        $ ./helloworld
```



