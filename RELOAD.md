#Step 1: Replace reloaded code

Directory tree of this archive:

```
ROOT
 |-- src
 |    |-- checkoutservice
 |    |    |-- genproto
 |    |    |    |-- demo.pb.go
 |    |    |    |-- demo.pb.reload.go (+)
 |    |    |-- Dockerfile
 |    |    |-- Gopkg.toml
 |    |    |-- main.go
 |    |-- frontend
 |    |    |-- genproto
 |    |    |    |-- demo.pb.go
 |    |    |    |-- demo.pb.reload.go (+)
 |    |    |-- Dockerfile
 |    |    |-- Gopkg.toml
 |    |    |-- handlers.go
 |    |-- productcatalogservice
 |    |    |-- genproto
 |    |    |    |-- demo.pb.go
 |    |    |    |-- demo.pb.reload.go (+)
 |    |    |-- Dockerfile
 |    |    |-- Gopkg.toml
 |    |    |-- server.go
 |    |-- shippingservice
 |         |-- genproto
 |         |    |-- demo.pb.go
 |         |    |-- demo.pb.reload.go (+)
 |         |-- Dockerfile
 |         |-- Gopkg.toml
 |         |-- server.go
 |-- reload
 |    |-- build.sh
 |    |-- dep.sh
 |    |-- genproto.sh
 |    |-- reload.sh
 |-- build.sh
```
1. Replace the original files under **src** with the files above, (files following with **(+)** are newly added).
2. Copy **.sh** files under **reload** to each service.
3. Delete **vendor** in **.dockerignore** for each service.
4. Execute **build.sh** to manage dependencies.
5. If you want to regenerate **.proto** file, execute **genproto.sh**.

# Step 2: Launch the services

Execute:

```shell
skaffold dev
```

#Step 3: Use client collect tracing data

1. Download code from: https://github.com/AleckDarcy/reload.git.
2. Run test case in **reload/core/client/core/client_test.go**.