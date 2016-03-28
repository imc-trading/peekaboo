<p align="center">
  <img src="img/logo.png" width="50%">
</p>

Expose hardware info using a REST API written in Go and a Front-End written in AngularJS.


## FrontEnd

```
http://myserver.example.com:5050
```

## BackEnd

### Endpoints

```
/api/v1/network/interfaces
/api/v1/network/routes
/api/v1/system
/api/v1/system/os
/api/v1/system/cpu
/api/v1/system/memory
/api/v1/system/sysctls (Linux only)
/api/v1/storage/disks (Linux only)
/api/v1/storage/mounts (Linux only)
/api/v1/storage/lvm/physvols (Linux only)
/api/v1/storage/lvm/logvols (Linux only)
/api/v1/storage/lvm/volgrps (Linux only)
/api/v1/docker
/api/v1/docker/containers
/api/v1/docker/images
```

### Options

**envelope**

Include an envelope with error, cache and status information.

```
?envlope=true
```

**Example:**

```json
{
  "cache": {
    "lastUpdated": "...",
    "timeoutSec": 0,
    "fromCache": false
  },
  "data": {
  },
  "error": [],
  "status": 200
}
```

**refresh**

Refresh cache.

```
?refresh=true
```

**indent**

Don't indent JSON.

```
?indent=false
```

**Example:**

```bash
curl http://myserver.example.com:5050/api/v1/system
```

# Usage

```bash
Peekaboo

Usage:
  peekaboo daemon [--debug] [--bind=<addr>] [--static=<dir>]
  peekaboo list
  peekaboo get <hardware-type>
  peekaboo -h | --help
  peekaboo --version

Commands:
  daemon               Start as a daemon serving HTTP requests.
  list                 List hardware names available.
  get                  Return information about hardware.

Arguments:
  hardware-type        Name of hardware to return information about.

Options:
  -h --help            Show this screen.
  --version            Show version.
  -d --debug           Debug.
  -b --bind=<addr>     Bind to address and port. [default: 0.0.0.0:5050]
  -s --static=<dir>    Directory for static content. [default: static]
```

# Setup Go on Linux

```bash
sudo yum install -y golang
mkdir ~/go
export GOPATH=~/go
export PATH=$GOPATH/bin:$PATH
go get github.com/constabulary/gb/...
```

## Build and run

```bash
gb build
sudo bin/peekaboo daemon -d
```

## Build RPM

Fiest make sure you have Docker configured.

```bash
make rpm
```

## Change configuration

```bash
systemctl stop peekaboo
vi /etc/sysconfig/peekaboo
```

Change port to 8080.

**Example:**

```
OPTIONS="--bind 0.0.0.0:8080"
```

Reload SystemD and then restart Peekaboo.

```bash
systemctl start peekaboo
```

## Install using Brew on Mac OS X

```bash
brew tap mickep76/funk-gnarge
brew install peekaboo
ln -sfv /usr/local/opt/peekaboo/*.plist ~/Library/LaunchAgents
launchctl load ~/Library/LaunchAgents/homebrew.mxcl.peekaboo.plist
```

## Change configuration

```bash
launchctl unload ~/Library/LaunchAgents/homebrew.mxcl.peekaboo.plist
vi ~/Library/LaunchAgents/homebrew.mxcl.peekaboo.plist
```

Change port to 8080.

**Example:**

```
...
    <string>--bind</string>
    <string>0.0.0.0:8080</string>
..
```

Restart Peekaboo.

```bash
launchctl load ~/Library/LaunchAgents/homebrew.mxcl.peekaboo.plist
```
