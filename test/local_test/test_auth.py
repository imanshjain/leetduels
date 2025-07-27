import pytest

def test_session(get_session):
    assert "idToken" in get_session, "Session should contain idToken"