# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from productpb import product_pb2 as productpb_dot_product__pb2


class ProductInfoStub(object):
    """Missing associated documentation comment in .proto file"""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.addProduct = channel.unary_unary(
                '/productpb.ProductInfo/addProduct',
                request_serializer=productpb_dot_product__pb2.Product.SerializeToString,
                response_deserializer=productpb_dot_product__pb2.ProductID.FromString,
                )
        self.getProduct = channel.unary_unary(
                '/productpb.ProductInfo/getProduct',
                request_serializer=productpb_dot_product__pb2.ProductID.SerializeToString,
                response_deserializer=productpb_dot_product__pb2.Product.FromString,
                )
        self.addAllProducts = channel.stream_unary(
                '/productpb.ProductInfo/addAllProducts',
                request_serializer=productpb_dot_product__pb2.Product.SerializeToString,
                response_deserializer=productpb_dot_product__pb2.ProductIDs.FromString,
                )
        self.getAllProducts = channel.unary_stream(
                '/productpb.ProductInfo/getAllProducts',
                request_serializer=productpb_dot_product__pb2.Empty.SerializeToString,
                response_deserializer=productpb_dot_product__pb2.Product.FromString,
                )


class ProductInfoServicer(object):
    """Missing associated documentation comment in .proto file"""

    def addProduct(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def getProduct(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def addAllProducts(self, request_iterator, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def getAllProducts(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ProductInfoServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'addProduct': grpc.unary_unary_rpc_method_handler(
                    servicer.addProduct,
                    request_deserializer=productpb_dot_product__pb2.Product.FromString,
                    response_serializer=productpb_dot_product__pb2.ProductID.SerializeToString,
            ),
            'getProduct': grpc.unary_unary_rpc_method_handler(
                    servicer.getProduct,
                    request_deserializer=productpb_dot_product__pb2.ProductID.FromString,
                    response_serializer=productpb_dot_product__pb2.Product.SerializeToString,
            ),
            'addAllProducts': grpc.stream_unary_rpc_method_handler(
                    servicer.addAllProducts,
                    request_deserializer=productpb_dot_product__pb2.Product.FromString,
                    response_serializer=productpb_dot_product__pb2.ProductIDs.SerializeToString,
            ),
            'getAllProducts': grpc.unary_stream_rpc_method_handler(
                    servicer.getAllProducts,
                    request_deserializer=productpb_dot_product__pb2.Empty.FromString,
                    response_serializer=productpb_dot_product__pb2.Product.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'productpb.ProductInfo', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ProductInfo(object):
    """Missing associated documentation comment in .proto file"""

    @staticmethod
    def addProduct(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/productpb.ProductInfo/addProduct',
            productpb_dot_product__pb2.Product.SerializeToString,
            productpb_dot_product__pb2.ProductID.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def getProduct(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/productpb.ProductInfo/getProduct',
            productpb_dot_product__pb2.ProductID.SerializeToString,
            productpb_dot_product__pb2.Product.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def addAllProducts(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_unary(request_iterator, target, '/productpb.ProductInfo/addAllProducts',
            productpb_dot_product__pb2.Product.SerializeToString,
            productpb_dot_product__pb2.ProductIDs.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def getAllProducts(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/productpb.ProductInfo/getAllProducts',
            productpb_dot_product__pb2.Empty.SerializeToString,
            productpb_dot_product__pb2.Product.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)
