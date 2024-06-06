# Do an MMH NT219.O21 
## TODO
- encrypt/decrypt kyber using openquantumsafe's openssl - github.com/open-quantum-safe/openssl - github.com/open-quantum-safe/liboqs
- encrypt/decrypt kyber using openquantumsafe's openssl 
- establish keys of symmetric-key systems in higher-level protocols (TLS)
- Resources
    <!--https://github.com/open-quantum-safe/openssl - https://github.com/open-quantum-safe/liboqs ` -->
    - https://github.com/open-quantum-safe/liboqs-python.git - https://github.com/pyca/pyopenssl.git
    - https://github.com/GiacomoPope/kyber-py.git
    - https://cryptopedia.dev/posts/kyber/ - algorithm details
## DOCS
- install liboqs-py (replace --parallel option match the number of available cores on your system)
``` shell
mkdir oqs && cd oqs
git clone --depth=1 https://github.com/open-quantum-safe/liboqs-python
cd liboqs-python
pip install .
# run example 
python3 liboqs-python/examples/kem.py
python3 liboqs-python/examples/sig.py
python3 liboqs-python/examples/rand.py
```