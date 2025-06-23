#!/usr/bin/env python3
"""
Simple test client for the Distributed Key-Value Database
This script tests the basic functionality of the server
"""

import socket
import time

def test_basic_operations():
    """Test basic key-value operations"""
    try:
        # Connect to the server
        client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        client.connect(('localhost', 6000))

        # Read initial message
        response = client.recv(1024).decode().strip()
        print(f"Server says: {response}")

        # Test SIGNUP
        print("\n=== Testing SIGNUP ===")
        client.send(b"SIGNUP\n")
        response = client.recv(1024).decode().strip()
        print(f"Server says: {response}")

        client.send(b"testuser\n")
        response = client.recv(1024).decode().strip()
        print(f"Server says: {response}")

        client.send(b"testpass\n")
        response = client.recv(1024).decode().strip()
        print(f"Server says: {response}")

        # Test SET operation
        print("\n=== Testing SET Operation ===")
        client.send(b"SET key1 value1\n")
        response = client.recv(1024).decode().strip()
        print(f"SET key1 value1 -> {response}")

        client.send(b"SET key2 value2\n")
        response = client.recv(1024).decode().strip()
        print(f"SET key2 value2 -> {response}")

        # Test GET operation
        print("\n=== Testing GET Operation ===")
        client.send(b"GET key1\n")
        response = client.recv(1024).decode().strip()
        print(f"GET key1 -> {response}")

        client.send(b"GET key2\n")
        response = client.recv(1024).decode().strip()
        print(f"GET key2 -> {response}")

        # Test HAS operation
        print("\n=== Testing HAS Operation ===")
        client.send(b"HAS key1\n")
        response = client.recv(1024).decode().strip()
        print(f"HAS key1 -> {response}")

        client.send(b"HAS nonexistent\n")
        response = client.recv(1024).decode().strip()
        print(f"HAS nonexistent -> {response}")

        # Test DELETE operation
        print("\n=== Testing DELETE Operation ===")
        client.send(b"DELETE key1\n")
        response = client.recv(1024).decode().strip()
        print(f"DELETE key1 -> {response}")

        client.send(b"GET key1\n")
        response = client.recv(1024).decode().strip()
        print(f"GET key1 (after delete) -> {response}")

        client.close()
        print("\n=== Test completed successfully! ===")

    except Exception as e:
        print(f"Error: {e}")

def test_multiple_users():
    """Test multiple user sessions"""
    print("\n=== Testing Multiple Users ===")

    # User 1
    client1 = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client1.connect(('localhost', 6000))
    client1.recv(1024)  # Initial message

    client1.send(b"SIGNUP\n")
    client1.recv(1024)
    client1.send(b"user1\n")
    client1.recv(1024)
    client1.send(b"pass1\n")
    client1.recv(1024)

    # User 2
    client2 = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client2.connect(('localhost', 6000))
    client2.recv(1024)  # Initial message

    client2.send(b"SIGNUP\n")
    client2.recv(1024)
    client2.send(b"user2\n")
    client2.recv(1024)
    client2.send(b"pass2\n")
    client2.recv(1024)

    # User 1 sets data
    client1.send(b"SET user1_key user1_value\n")
    response = client1.recv(1024).decode().strip()
    print(f"User1 SET -> {response}")

    # User 2 sets data
    client2.send(b"SET user2_key user2_value\n")
    response = client2.recv(1024).decode().strip()
    print(f"User2 SET -> {response}")

    # User 1 tries to get user2's data (should fail)
    client1.send(b"GET user2_key\n")
    response = client1.recv(1024).decode().strip()
    print(f"User1 trying to GET user2's key -> {response}")

    # User 1 gets their own data
    client1.send(b"GET user1_key\n")
    response = client1.recv(1024).decode().strip()
    print(f"User1 GET their own key -> {response}")

    client1.close()
    client2.close()
    print("Multiple user test completed!")

if __name__ == "__main__":
    test_basic_operations()
    time.sleep(1)
    test_multiple_users()
