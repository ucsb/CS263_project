from flask import Flask, request
import flask
import py_eureka_client.eureka_client as eureka_client
import requests
import urllib

rest_port = 8084
eureka_client.init(eureka_server="http://localhost:8761/eureka",
                   app_name="checkout-service",
                   instance_port=rest_port)

app = Flask(__name__)

orders_list = []

@app.route("/", methods=['POST'])
def checkout():
    name = flask.request.json['name']
    quantity =flask.request.json['quantity']
    
    try:
        res = eureka_client.do_service("inventory-service-database", f"/doughnuts/checkout/{name}/{quantity}")
        if res == "Doughnut checked out successfully":
            orders_list.append(("Doughnut bought: " + name, "Quantity bought: " + str(quantity)))
            print(orders)
        return res
    except urllib.request.HTTPError as e:
        # If all nodes are down, a `HTTPError` will raise.
        return e

@app.route("/orders", methods=['GET'])  
def orders():
    output = ""
    for first, second in orders_list:
        output = output + first + " " + second + "<br />"
    return output

if __name__ == "__main__":
    app.run(host='0.0.0.0', port = rest_port)