import grpc
import seleniumserver__pb2

def startServer():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))