import grpc
from concurrent import futures
import seleniumserver_pb2
import seleniumserver_pb2_grpc
from seleniumserver_pb2 import HtmlArgs
from selenium import webdriver
import time

class ChromeDriver:
    def __init__(self):
        opt = webdriver.ChromeOptions()
        opt.add_argument("chromedriver")
        #opt.add_argument("--headless")
        self.driver = webdriver.Chrome(opt)

    def end_driver(self):
        self.driver.quit()

    def get_html(self, url):
        self.driver.get(url)
        time.sleep(5)
        return self.driver.page_source


chrome_driver = ChromeDriver()

class SeleniumServicer(seleniumserver_pb2_grpc.SeleniumServerServicer):
    def GetHtml(self, request, context):
        url = request.Url

        global chrome_driver

        response = seleniumserver_pb2.HtmlResponse(Html=chrome_driver.get_html(url))
        return response
    

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

    seleniumserver_pb2_grpc.add_SeleniumServerServicer_to_server(SeleniumServicer(), server)

    server.add_insecure_port('[::]:61337')

    server.start()
    server.wait_for_termination()
    global chrome_driver

    chrome_driver.end_driver()



serve()

