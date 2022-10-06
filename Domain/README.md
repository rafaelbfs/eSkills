## Generating Code

``` shell
 protoc -I=Proto --go_out=. --go-grpc_out=. Proto/skills.proto
```

## Testing MongoDB

For the unit tests. a MongoDB Atlas account was used, you can create one for free at
https://www.mongodb.com/cloud/atlas/ though, obviously, with several limitations it's 
enough for testing. In it, you can choose to connect to the database 
with username/password or through certificates which can be generated and downloaded

then you create a *local.env* file at this same directory git does not 
add any of the .env files to the repository as they are linked to my private MongoDB 
account but you can look at *sample.dotenvfile.env* to see how it's used

