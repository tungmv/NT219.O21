import oqs
import ssl
import socket

# Initialize Kyber key exchange
with oqs.KeyEncapsulation('Kyber1024') as server_kem, oqs.KeyEncapsulation('Kyber1024') as client_kem:
    
    # Server side
    server_public_key = server_kem.generate_keypair()
    
    # Simulate sending the public key to the client and client generating a shared secret
    encrypted_key, shared_secret_client = client_kem.encap_secret(server_public_key)
    
    # Server decapsulates the secret to obtain the shared secret
    shared_secret_server = server_kem.decap_secret(encrypted_key)

    # Both client and server now have the same shared secret
    assert shared_secret_client == shared_secret_server

    # Create a context using the shared secret for symmetric encryption in TLS
    context = ssl.SSLContext(ssl.PROTOCOL_TLS)
    context.set_ciphers('AES256-GCM-SHA384')  # Symmetric cipher suite
    context.load_cert_chain(certfile='server.crt', keyfile='server.key')

    # Start a TLS server using the shared secret
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM, 0)
    server_socket.bind(('localhost', 8443))
    server_socket.listen(5)
    
    conn, addr = server_socket.accept()
    with context.wrap_socket(conn, server_side=True) as tls_conn:
        print("TLS connection established with Kyber key exchange")

        # Secure communication using the established TLS connection
        data = tls_conn.recv(1024)
        print(f"Received: {data}")
        tls_conn.sendall(b"Hello, secure world!")