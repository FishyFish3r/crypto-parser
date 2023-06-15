import grpc
from concurrent import futures
import seleniumserver_pb2
import seleniumserver_pb2_grpc
from seleniumserver_pb2 import HtmlArgs

import google.protobuf.empty_pb2

class SeleniumServicer(seleniumserver_pb2_grpc.SeleniumServerServicer):
    def GetHtml(self, request, context):
        url = request.Url

        return google.protobuf.empty_pb2.Empty()
    

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

    seleniumserver_pb2_grpc.add_SeleniumServerServicer_to_server(SeleniumServicer(), server)

    server.add_insecure_port('[::]:61337')

    server.start()
    server.wait_for_termination()

serve()