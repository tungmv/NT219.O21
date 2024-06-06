
import oqs
import ssl
import socket

# Client side (normally in a separate script)
with oqs.KeyEncapsulation('Kyber1024') as client_kem:
    client_public_key = client_kem.generate_keypair()
    encrypted_key, shared_secret_client = client_kem.encap_secret(server_public_key)
    shared_secret_server = server_kem.decap_secret(encrypted_key)

    context = ssl.SSLContext(ssl.PROTOCOL_TLS)
    context.set_ciphers('AES256-GCM-SHA384')
    context.load_verify_locations(cafile='ca.crt')

    client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client_socket.connect(('localhost', 8443))
    with context.wrap_socket(client_socket, server_hostname='localhost') as tls_conn:
        tls_conn.sendall(b"Hello, server!")
        data = tls_conn.recv(1024)
        print(f"Received: {data}")
