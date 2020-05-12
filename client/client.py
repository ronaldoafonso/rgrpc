
import time
import grpc

import productpb.product_pb2 as pb
import productpb.product_pb2_grpc as pbgrpc


def run():
    channel = grpc.insecure_channel('server_productinfo_1:50051')
    stub = pbgrpc.ProductInfoStub(channel)

    product1 = pb.Product(name = "Product 1",
                          description = "Desc Product 1",
                          price = 1.0)
    response = stub.addProduct(product1)
    print("Added product1 - Response:", response)

    productID1 = pb.ProductID(value = response.value)
    productInfo = stub.getProduct(productID1)
    print("Get productID1: ", productInfo)


if __name__ == '__main__':
    run()
