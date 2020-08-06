# Own guide for setting u protobuf in golang

I was working with Kafka and Watermill library in a recent project. In the projects, I passed data between producers and consumers by serializing the data using JSON format. After further discussions, I realised that there are many alternatives (Gob, XML, Protobuf). In this short repository, I explore the process of setting up Protobufs and try to see if this is a good alternative.

## Setting up Protobufs for Go lang

_Plug-ins required_
- github.com/golang/protobuf
- github.com/golang/protobuf/proto
- github.com/golang/protobuf/protoc-gen-go

(for proto compilers and other libraries for code)

### 1: Write your .proto file
With reference to **person.proto**:

- Set syntax, package
- Set go_package
    - ```option go_package = "./pb";```   
    - Path to store the output proto go file
    
### 2: Generate .go file from .proto file
With reference to **pb directory**

- .go files contains code for the struct, as well as getters and setters
- Generate using the protoc command
```
$ protoc --go_out=. *.proto
```

### 3: Use the .go output file to create structs and serialise it
With reference to **main.go**

- Create struct form the pb package
    - ``	shane := &pb.Person{
        		Name: "Shane",
        		Age:  21,
        		Emails: emails,
        	}``
        	
        	
<hr>
<br>

## Notes on .proto files
- Data types
    - possibility of enums etc.
    - Nested messages(struct)
    - Array of items (repeated)
          
          
```
message Email { 
                      string primary =1;
                  }
                
                  repeated Email emails = 3;
```
<hr>
<br>

_Written by Shane Lim Xuan Yu_

_References_ :
- https://github.com/protocolbuffers/protobuf/tree/master/examples
- (tutorial) https://tutorialedge.net/golang/go-protocol-buffer-tutorial/