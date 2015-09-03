Expose hardware info using JSON/REST and HTML Front-End.

**FrontEnd:**

http://myserver.example.com:5050

**JSON endpoints:**

```
/json
/cpu/json
/disks/json
/lvm/json
/lvm/log_vols/json
/lvm/phys_vols/json
/lvm/vol_grps/json
/memory/json
/mounts/json
/network/interfaces/json
/network/json
/network/json
/network/routes/json
/opsys/json
/pci/json
/sysctl/json
```

**Example:**

```bash
curl http://myserver.example.com:5050/cpu/json
```

# Usage

```bash
Usage:
  peekaboo [OPTIONS]

Application Options:
  -v, --verbose       Verbose
      --version       Version
  -b, --bind-addr=    Bind to address (0.0.0.0)
  -p, --port=         Port (5050)
  -s, --static-dir=   Static content (static)
  -t, --template-dir= Templates (templates)

Help Options:
  -h, --help          Show this help message
```

# Setup Go build environment

```bash
yum install golang
mkdir ~/go
export GOPATH=~/go
export PATH=$GOPATH/bin:$PATH
go get github.com/constabulary/gb/...
```

## Build

```bash
make
sudo bin/peekaboo -s src/github.com/mickep76/peekaboo/static -t src/github.com/mickep76/peekaboo/templates
```

## Build RPM

```bash
yum install rpm-build
make rpm
sudo rpm -i peekaboo-<version>-<release>.rpm
sudo systemctl enable peekaboo
sudo systemctl start peekaboo
```

# Change port or bind address

```bash
vi /etc/systemd/system/peekaboo.service
```

Add -b bind address, defaults to "0.0.0.0". For port add -p, defaults to 5050.

```
ExecStart=/usr/bin/peekaboo -s /var/lib/peekaboo/static -t /var/lib/peekaboo/templates -b <bind addr.> -p <port>
```

Reload SystemD and then restart Peekaboo.

```bash
systemctl daemon-reload
systemctl restart peekaboo
```
