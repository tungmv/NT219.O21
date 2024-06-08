OPENSSL CONFIGURATION AND CHECK CERTIFICATES
======================

To check if the certificates are valid, you have to first configure OpenSSL

## CONFIGURE OPENSSL

Navigate to wherever OpenSSL is installed on your machine.
Run:
```sh
    which openssl
```

on Ubuntu to check the root directory of OpenSSL, navigate to said directory and open `openssl.cnf` with a text editor of choice with elevated privilege

```sh
sudo vi openssl.cnf
```

change/add the content to the two below:

```plaintext
[provider_sect]
default = default_sect
oqsprovider = oqsprovider_sect
[oqsprovider_sect]
activate = 1
```

**do the same with the global openssl.cnf file**, typically located in:

```sh
/etc/ssl/openssl.cnf
```
for the best experience (optional)

## RUN OPENSSL COMMANDS TO CHECK THE CERTIFICATES

Given that I ran the following openssl commands to generate those files above:
```sh
openssl req -x509 -new -newkey dilithium3 -keyout dilithium3_CA.key -out dilithium3_CA.crt -nodes  -days 365 -config /usr/openssl/openssl.cnf
openssl genpkey -algorithm dilithium3 -out dilithium3_srv.key
openssl req -new -newkey dilithium3 -keyout dilithium3_srv.key -out dilithium3_srv.csr -nodes  -config /usr/openssl/openssl.cnf
openssl x509 -req -in dilithium3_srv.csr -out dilithium3_srv.crt -CA dilithium3_CA.crt -CAkey dilithium3_CA.key -CAcreateserial -days 365
```

note that your openssl.cnf can vary

Open a terminal, run `openssl s_server` with parameters of choice, for example in my case
```sh
openssl s_server -cert dilithium3_srv.crt -key dilithium3_srv.key -CAfile dilithium3_CA.crt -www -tls1_3 -groups kyber768:kyber1024
```

Open a second terminal, run 'openssl s_client' with custom parameters, for example in my case

```sh
openssl s_client -groups kyber768
```

Read the response on the client terminal closely. If everything is OK, you will see this:

```console
...
-----END CERTIFICATE-----
subject=C=VN, ST=Dong Nai, L=Bien Hoa, O=UIT, OU=ATTT2022, CN=22521115.wuaze.com, emailAddress=22521115@gm.uit.edu.vn
issuer=C=VN, ST=Dong Nai, L=Bien Hoa, O=UIT, OU=ATTT2022, CN=22521115.wuaze.com, emailAddress=22521115@gm.uit.edu.vn
---
No client certificate CA names sent
Peer signature type: dilithium3
---
SSL handshake has read 16169 bytes and written 1619 bytes
Verification: OK
---
New, TLSv1.3, Cipher is TLS_AES_256_GCM_SHA384
Server public key is 192 bit
This TLS version forbids renegotiation.
Compression: NONE
Expansion: NONE
No ALPN negotiated
Early data was not sent
Verify return code: 0 (ok)
---
---
...
```

If you encounter error, roll to the end of the response and check for verify return code error, with the most common being 
* error 19: self-signed certificate: you missed the CA.crt file. There are 2 .crt files with the CA one have to be included after `-CAfile` flag in your `openssl` commands (assuming you ran `openssl verify` earlier and got OK)
* error 21: unable to verify the first certificate: you missed a certificate as arguments of the `openssl` commands.

Good mf luck. Anyway tungminhv still haven't done shit. We burnin in hell with this one.