- [build simple dt(distributed transaction) trans with dtm(grpc)](#build-simple-dtdistributed-transaction-trans-with-dtmgrpc)
  - [target](#target)
  - [implement the following steps](#implement-the-following-steps)
  - [how to run](#how-to-run)


## build simple dt(distributed transaction) trans with dtm(grpc)

### target
- client => dtm => server  (=> = call)

### implement the following steps
1. proto
2. server
3. service
4. client
5. dtm

### how to run
1. run [dtm](https://dtm.pub) server
2. run grpc server
3. client call housed dtm interface 