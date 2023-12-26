"""
Install necessary packages with the following commands:
poetry init # create new venv

poetry add grpcio
poetry add grpcio-tools

Generate the gRPC stubs with the following command:
python -m grpc_tools.protoc -I ./proto --python_out=./proto --pyi_out=./proto --grpc_python_out=./proto ./proto/demo_grpc.proto
"""
import sys

# Fix protoc relative import issues
sys.path.append("./proto")

import grpc

# generated gRPC stubs
import proto.demo_grpc_pb2
import proto.demo_grpc_pb2_grpc


def run():
    # Create a gRPC channel to connect to the server
    with grpc.insecure_channel('localhost:10000') as channel:
        # Create a stub for the gRPC service
        stub = proto.demo_grpc_pb2_grpc.DemoServiceStub(channel)

        # Make a gRPC request
        message = proto.demo_grpc_pb2.HelloRequest(name="Python")
        response: proto.demo_grpc_pb2.HelloResponse = stub.SayHello(message)

        # Process the response
        print(response.message)  # Replace with your response handling logic

        # Adder request
        message = proto.demo_grpc_pb2.AdderRequest(a=10, b=20)
        response: proto.demo_grpc_pb2.AdderResponse = stub.Adder(message)

        print(response.result)

        # Server streaming
        message = proto.demo_grpc_pb2.HelloRequest(name="Call")
        for char in stub.StringToChar(message):
            # char is type of proto.demo_grpc_pb2.CharResponse
            print(chr(char.char))

        # request-side streaming
        chars = ['p', 'y', 't', 'h', 'o', 'n']

        # create a iterator of CharRequest
        def msg_iter(chars):
            for char in chars:
                msg = proto.demo_grpc_pb2.CharRequest(char=ord(char))
                yield msg

        response: proto.demo_grpc_pb2.HelloResponse = stub.CharToString(msg_iter(chars))
        print(response.message)

        # bidirectional streaming
        responses = stub.AllCharUpper(msg_iter(chars))
        for response in responses:
            response: proto.demo_grpc_pb2.CharResponse
            print(chr(response.char))


if __name__ == '__main__':
    run()
