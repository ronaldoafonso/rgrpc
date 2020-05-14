
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

    # Add products and get them one at a time
    for product in products:
        p = pb.Product(name=product["name"],
                       description=product["description"],
                       price=product["price"])
        response = stub.addProduct(p)
        print("Added {0} - Response: {1}".format(product["name"], response))

        productID = pb.ProductID(value=response.value)
        productInfo = stub.getProduct(productID)
        print("Get productID: {0}".format(productInfo))

    # Add some more products
    def sendProducts():
        products = [
            {"name": "Product 3", "description": "Desc Product 3", "price": 3.0},
            {"name": "Product 4", "description": "Desc Product 4", "price": 4.0}
        ]
        for product in products:
            yield pb.Product(name=product["name"],
                             description=product["description"],
                             price=product["price"])

    productIDs = stub.addAllProducts(sendProducts())
    print("-----------------------")
    for productID in productIDs.productIDs:
        print("ID: {0}.".format(productID.value))

    # Get all products
    print("-----------------------")
    print("Get All Products")
    for product in stub.getAllProducts(pb.Empty()):
        print("ID: {0}.".format(product.id))
        print("Name: {0}.".format(product.name))
        print("Desc: {0}.".format(product.description))
        print("Price: {0}.".format(product.price))

if __name__ == '__main__':
    run()
