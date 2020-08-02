### go-sshclient

**Usage**:

```go
go run sshclient.go -u <username> -p -A <Ip Address> 
Enter password: 
```

**Saving the auth data in a json file**:

```go
go run sshclient.go -u <username> -p -A <Ip Address> -j
```

**Use the json file to auth:**

```go
go run sshclient.go -c auth.json
```



**Flag**:

**-u** ---> define username

**-p** ---> enter password

**-A** ---> define ip address

**-P** ---> define server port (default: 22)

**-j** ---> save the auth config in a json file with encoded password

**-c** --->  load the json config file for auth 