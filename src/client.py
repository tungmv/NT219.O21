import oqs
from hashlib import sha256

# Initialize OQS
oqs.default_initialize()

# Create a Kyber keypair for client
client_sk = oqs.KeyEncapsulation("Kyber-512")
client_pk = client_sk.generate_keypair()
