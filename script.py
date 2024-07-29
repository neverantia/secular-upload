import requests

url = 'http://localhost:8080/upload'

file_path = "go.mod" 


with open(file_path, 'rb') as file:
    files = {'file': file}
    
    response = requests.post(url, files=files)
    
    print(response.text)
