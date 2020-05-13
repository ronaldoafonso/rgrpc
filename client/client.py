
import time
import grpc

import productpb.product_pb2 as pb
import productpb.product_pb2_grpc as pbgrpc


products = [
    {"name": "Product 1", "description": "Desc Product 1", "price": 1.0},
    {"name": "Product 2", "description": "Desc Product 2", "price": 2.0}
]

def run():
    channel = grpc.insecure_channel('server_productinfo_1:50051')
    stub = pbgrpc.ProductInfoStub(channel)

    for product in products:
        p = pb.Product(name=product["name"],
                       description=product["description"],
                       price=product["price"])
        response = stub.addProduct(p)
        print("Added {0}- Response: {1}".format(product["name"], response))

        productID = pb.ProductID(value=response.value)
        productInfo = stub.getProduct(productID)
        print("Get productID: {0}".format(productInfo))


if __name__ == '__main__':
    run()
