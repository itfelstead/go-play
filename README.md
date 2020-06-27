# go-play

Just a personal area for playing with go.

## Workflow Approach:

For now, for local dev, just mount the source dir in the git repo, and mount it into a golang container for building and running.
Edit source either inside the container (vi) or on host with Visual Studio Code.  
Alternatively, do simple experiments in the online playground at play.golang.org.

*Reminder: When PC finally gets Windows v2004, try this:  https://code.visualstudio.com/docs/remote/containers*.  

**Golang Container:** based on https://hub.docker.com/_/golang
- installs GO
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



