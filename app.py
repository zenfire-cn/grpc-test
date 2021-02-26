import grpc
import helloworld_pb2
import helloworld_pb2_grpc

def run():
  with grpc.insecure_channel('127.0.0.1:6544') as channel:
    stub = helloworld_pb2_grpc.HelloServiceStub(channel)

    response = stub.SayHello(helloworld_pb2.HelloReq(name="python"))
    print("Client received: " + str(response))

    response = stub.Add(helloworld_pb2.AddReq(a=10, b=10))
    print("Client received: " + str(response))
 
if __name__ == '__main__':
  run()