import requests

# Define the URL of the upload endpoint
url = 'http://localhost:8080/upload'

# Path to the file you want to upload
file_path = "go.mod" 


# Open the file in binary mode
with open(file_path, 'rb') as file:
    # Define the files parameter for the POST request
    files = {'file': file}
    
    # Make the POST request to upload the file
    response = requests.post(url, files=files)
    
    # Print the server response
    print(response.text)
