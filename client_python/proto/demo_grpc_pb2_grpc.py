# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import demo_grpc_pb2 as demo__grpc__pb2


class DemoServiceStub(object):
    """what RPC methods will this service have
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SayHello = channel.unary_unary(
                '/demo_proto.DemoService/SayHello',
                request_serializer=demo__grpc__pb2.HelloRequest.SerializeToString,
                response_deserializer=demo__grpc__pb2.HelloResponse.FromString,
                )
        self.StringToChar = channel.unary_stream(
                '/demo_proto.DemoService/StringToChar',
                request_serializer=demo__grpc__pb2.HelloRequest.SerializeToString,
                response_deserializer=demo__grpc__pb2.CharResponse.FromString,
                )
        self.Adder = channel.unary_unary(
                '/demo_proto.DemoService/Adder',
                request_serializer=demo__grpc__pb2.AdderRequest.SerializeToString,
                response_deserializer=demo__grpc__pb2.AdderResponse.FromString,
                )


class DemoServiceServicer(object):
    """what RPC methods will this service have
    """

    def SayHello(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def StringToChar(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Adder(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DemoServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'SayHello': grpc.unary_unary_rpc_method_handler(
                    servicer.SayHello,
                    request_deserializer=demo__grpc__pb2.HelloRequest.FromString,
                    response_serializer=demo__grpc__pb2.HelloResponse.SerializeToString,
            ),
            'StringToChar': grpc.unary_stream_rpc_method_handler(
                    servicer.StringToChar,
                    request_deserializer=demo__grpc__pb2.HelloRequest.FromString,
                    response_serializer=demo__grpc__pb2.CharResponse.SerializeToString,
            ),
            'Adder': grpc.unary_unary_rpc_method_handler(
                    servicer.Adder,
                    request_deserializer=demo__grpc__pb2.AdderRequest.FromString,
                    response_serializer=demo__grpc__pb2.AdderResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'demo_proto.DemoService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class DemoService(object):
    """what RPC methods will this service have
    """

    @staticmethod
    def SayHello(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/demo_proto.DemoService/SayHello',
            demo__grpc__pb2.HelloRequest.SerializeToString,
            demo__grpc__pb2.HelloResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def StringToChar(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/demo_proto.DemoService/StringToChar',
            demo__grpc__pb2.HelloRequest.SerializeToString,
            demo__grpc__pb2.CharResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Adder(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/demo_proto.DemoService/Adder',
            demo__grpc__pb2.AdderRequest.SerializeToString,
            demo__grpc__pb2.AdderResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
