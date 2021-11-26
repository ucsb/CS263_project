import requests

'''
#requests made to inventory-service-database

URL = "http://localhost"
PORT = 8085

#example of creating or updating a doughnut
response = requests.post(f"{URL}:{PORT}/doughnuts/save", json={"id": "Glazed", "name": "Glazed", "price": 2.0, "inventory": 220})
#example of deleting a doughnut
response = requests.delete(f"{URL}:{PORT}/doughnuts/delete/Old%20Fashion")
'''

#requests made to checkout-service
URL = "http://192.168.0.176"
PORT = 8084

response = requests.post(f"{URL}:{PORT}/", json={'name': 'Glazed', 'quantity': 2})
print(response)