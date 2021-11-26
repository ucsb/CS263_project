from flask import Flask, request
import py_eureka_client.eureka_client as eureka_client
import requests

rest_port = 8084
eureka_client.init(eureka_server="http://localhost:8761/eureka",
                   app_name="checkout-service",
                   instance_port=rest_port)

app = Flask(__name__)
@app.route("/", methods=['GET'])
def get():
    return "hello world"

if __name__ == "__main__":
    app.run(host='0.0.0.0', port = rest_port)