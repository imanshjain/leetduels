import pytest
import websocket
import json

def test_init_message(get_session, get_ws):
    
    payload = {
        "idToken": get_session.get("idToken")
    }

    get_ws.send(json.dumps(payload))
    response = get_ws.recv()
    data = json.loads(response)

    payload = data.get("payload", {})
    assert payload.get("message") == "Welcome to LeetDuel!"
