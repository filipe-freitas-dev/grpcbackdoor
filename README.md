# grpcbackdoor

**grpcbackdoor** is a proof-of-concept backdoor built using **gRPC**, designed for educational and cybersecurity research purposes. It showcases how remote access and command execution can be implemented over a secure and efficient communication channel.

## 🛠️ What It Does

This tool allows remote command execution by establishing a single persistent connection between the client and the server. Once connected, commands can be sent and responses received through an encrypted gRPC stream.

## 🔧 How It Works

- 🧩 **gRPC with Protocol Buffers**: All communication between client and server is defined and structured using Protocol Buffers, ensuring efficient serialization and deserialization.
- 🔒 **End-to-End Encryption**: All data transmitted is encrypted, providing a secure communication channel that prevents eavesdropping and tampering.
- 📡 **HTTP/2 Transport Layer**: Communication is conducted over HTTP/2, which provides multiplexing, header compression, and lower latency.
- 🔁 **Streaming**: Instead of opening a new connection for each command, a single gRPC stream is used, allowing continuous, bidirectional communication.

## 📁 Project Structure

- `client/`: Contains the source code for the gRPC client (backdoor agent).
- `server/`: Contains the source code for the gRPC server (controller).
- `.proto`: Protocol Buffers definition file for RPC services and messages.

## 🚀 Technologies Used

- [gRPC](https://grpc.io/): Modern, high-performance RPC framework.
- [Protocol Buffers](https://developers.google.com/protocol-buffers): Interface description language used by gRPC.
- [Go](https://golang.org/): Programming language used for development.
- HTTP/2: Transport protocol for all RPC traffic.

## ⚠️ Legal Disclaimer

This project is intended **only** for **educational** and **authorized penetration testing** purposes. Unauthorized use of this software to compromise systems without consent is illegal and unethical. The authors assume no responsibility for misuse.

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

🔗 [Visit the Repository](https://github.com/filipe-freitas-dev/grpcbackdoor)

