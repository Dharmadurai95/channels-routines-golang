# main or default go route is responsible for executer our code line by line

## if you want to non blocking and do multiple thing concurrently we can use new go routine

## let say you have a function that responsible for making api func name will be checkstatus 
## by default we cannot call the checkstatus func before get the first checkstatus response 
## Here is the solution to call the checkstatus func concurrently  ---> go checkstatus()  using GO key word infront of the function we can create new go routine


## concurrency is like a event loop it will moniter the code at a same time but it will one sigle thread   (If once CPU)
## parallelism is multile thread executed at exact same time (if we have multiple CPU)

## if we want to communicate between two routine we have use CHANNEL IMPORTANT: if we create CHANNEL of type string then we cannot pass int value through this channel



