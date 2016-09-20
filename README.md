


#Run directly
Open one terminal:
```bash
go run tcp-server.go
```

Open another terminal:
```bash
go run tcp-client.go
```

#Dockerized

```bash
docker build -t tanjinfu/go-tcp-server:v1 .
docker run -i -t -p 31069:31069 tanjinfu/go-tcp-server:v1
```