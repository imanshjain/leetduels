import pytest
import requests
import websocket

API_KEY = "AIzaSyDArMH5J6UBVyuAluYqkjgjRI6o1riFJKM"
url = f"https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key={API_KEY}"

@pytest.fixture(scope="module", autouse=False)
def get_session():

    payload = {
        "email": "test@example.com",
        "password": "Test_User_001",
        "returnSecureToken": True
    }

    response = requests.post(url, json=payload)
    data = response.json()
    print(data)

    return data

@pytest.fixture(scope="module", autouse=False)
def get_ws(get_session):

    ws = websocket.create_connection("ws://localhost:8080/ws/initMessage?idToken=" + get_session.get("idToken"))
    yield ws
    ws.close()