### generate certs:

``
openssl.exe req -x509 -newkey rsa:4096 -sha256 -nodes -keyout tls.key -out tls.crt -subj "/CN=nimble.127.0.0.1.nip.io" -days 365
``


Now create secret

``
kubectl create secret tls nip --cert=./tls.crt --key=./tls.key
`` 