### go-sshclient

**Build:**

```go
go build sshclient.go
```

**Usage**:

```bash
./sshclient -u <username> -p -A <Ip Address> 
Enter password: 
```

**Saving the auth data in a json file**:

```bash
./sshclient -u <username> -p -A <Ip Address> -j
```

**Use the json file to auth:**

```bash
./sshclient -c auth.json
```



**Example:**

```bash
./sshclient -u ex0dia -p -A 192.168.1.100 
```

```bash
./sshclient -u ex0dia -p -A 192.168.1.100 -P 2222
```

```bash
./sshclient -u ex0dia -p -A 192.168.1.100 -j
```

```bash
./sshclient -c auth.json
```





**Flag**:

**-u** ---> define username

**-p** ---> enter password

**-A** ---> define ip address

**-P** ---> define server port (default: 22)

**-j** ---> save the auth config in a json file with encoded password

**-c** --->  load the json config file for auth 