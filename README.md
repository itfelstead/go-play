# go-play

Just a personal area for playing with go.

# Workflow Approach:

Use the online playground at play.golang.org for trivial experiments.  
Use a sandbox container integrated with VS Code via the `Remote Development / ms-vscode-remote.vscode-remote-extensionpack` feature for main development.

The following will give us a sandbox dev container, with a bind mount of the local project repo files: i.e. we can edit/run/git happily in our container, while the source files remain on the host.  
The container in the repo is based on the VS Code 'try go' container, which provided a nice starting point (see later section for details on how it was created, and other options).

## Using the VS Code Dev Sandbox with go-play

- Run VS Code
- Click the 'Open a remote Window' button (lower left corner), and select `Remote-Containers: Open folder in container`
- Select the `go-play` repo folder on your local system  
    The container will now be created and run based on the `.devcontainer/`.
- You can now edit and build, and use source control  
    *Note: i had git permission issues; upgrade of docker (not git) seemed to cure.. maybe a restart would too?)*  
    To build via CLI: Open a terminal (ctrl+shift+'), cd src/, go build helloworld.go  
    To build via VS Code: F5 (Run->Start Debugging). But you'll need to alter launch.json to point to the correct executable (default is server.go)

## (FYI only) Notes on VS Code based approaches

The selected approach was `Remote-Containers: Open folder in container`, where the container was heavily based on VS Code's 'try go' sample container.  
This will be described first, followed by some detais on other approaches that were considered.  

### (Chosen approach) `Remote-Containers: Open folder in container`: Try based container
The following steps were taken to build the 'try' based dev container.  
*Note: The resulting config is in this repo already, so no need to redo!*
- Grab VS Code 'try' repo;
    `cd <somewhere else>; git clone https://github.com/microsoft/vscode-remote-try-go.git`
- Copy its `.devcontainer/*`, and `.vscode/*` files into our repo
    That gives us the ready to go container with tools.
- Modify to suit  
    - altered container 'name' in devcontainer.json. I left the user as 'vscode'.
    - you might want to alter launch.json as it assumes the app is 'server.go'

### Alternative VS Code based approaches
The following is a summary of some other VS Code `Remote Development / ms-vscode-remote.vscode-remote-extensionpack` based approaches that were considered, to act as a reminder for future dev.  
All require you to use the 'Open a remote Window' button (lower left corner).  
See https://code.visualstudio.com/docs/remote/containers for more information.
- `Remote-WSL:`:
    *Reminder: try this when PC finally gets Windows v2004*
- `Remote-Containers: Try a Sample...`:
    - open your project Folder
    - Try; will create/run a container and mount your Folder
    - HOWEVER... repo source is held in a docker volume... so not too convinient, but...
    - good for getting example `Dockerfile`, `devcontainer.json`, `launch.json` and `settings.json` for your own sandbox containers: code is here; https://github.com/microsoft/vscode-remote-try-go
- `Remote-Containers: Attach to running container..`:
    Fine if you have your container running independently of VS Code (e.g. you use it for other things too).  
    It will need to install the VS Code server etc.. though.  
    Pre-installing VS Code server and extensions is not that well supported (https://github.com/microsoft/vscode-remote-release/issues/1718)  
- `Remote-SSH:`:
    Allows VS Code to edit files on a remote host, and it works ok but not of use to me in my home setup.
    Also, can SSH to a remote docker container. I've not tried that, but again not of use in my current setup.