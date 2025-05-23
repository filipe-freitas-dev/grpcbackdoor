rm *.pem

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=BR/ST=Sao Paulo/L=Sao Paulo/O=Casa/OU=Tecnologia/CN=*.domain.com/emailAddress=email@email.com"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-cert.pem -subj "/C=BR/ST=Sao Paulo/L=Sao Paulo/O=Office/OU=Computer/CN=*.domain.com/emailAddress=email1@email.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in server-cert.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in ca-cert.pem -noout -text
