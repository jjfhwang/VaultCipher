# VaultCipher

A Go library for securely encrypting and decrypting data using HashiCorp Vault as a key management system (KMS).

## Description

VaultCipher provides a simple and secure way to integrate encryption into Go applications by leveraging HashiCorp Vault's transit secrets engine. This library handles the complexities of interacting with Vault, allowing developers to focus on encrypting and decrypting data without managing encryption keys directly. By using Vault as a KMS, VaultCipher helps improve security posture by centralizing key management, auditing key usage, and enabling key rotation. It simplifies the process of integrating encryption into applications, ensuring that sensitive data is protected at rest and in transit.

## Features

*   **Encryption with Vault Transit Engine:** Seamlessly encrypt data using Vault's Transit Secrets Engine, ensuring secure encryption and decryption operations.
*   **Key Rotation Support:** Automatically handles key rotation as configured in Vault, preventing the need for manual key management.
*   **Authenticated Encryption:** Uses authenticated encryption modes (e.g., AES-GCM) to provide both confidentiality and integrity of the encrypted data.
*   **Simplified API:** Provides a straightforward API for encrypting and decrypting data, minimizing the learning curve and simplifying integration.
*   **Error Handling:** Robust error handling to gracefully manage potential issues when interacting with Vault, such as network errors or permission issues.

## Installation

To install VaultCipher, you need to have Go installed and configured. Use the following command:



This command will download and install the VaultCipher library and its dependencies.

Before using VaultCipher, ensure you have a Vault server running and configured with a Transit Secrets Engine. You will also need a Vault token with appropriate permissions to access the Transit Secrets Engine.

## Usage

Here are some examples of how to use the VaultCipher library:

First, create a new `VaultCipher` client:



**Explanation:**

*   **Import Statements:** Includes necessary import statements, including the Vault API client and the VaultCipher library.
*   **Vault Client Initialization:** Initializes a Vault client with the Vault server address and token.  Replace placeholders with actual values.
*   **VaultCipher Initialization:** Creates a new `VaultCipher` instance, specifying the Vault client and the Transit Secrets Engine path.
*   **Encryption:** Encrypts the plaintext data using the `Encrypt` method.
*   **Decryption:** Decrypts the ciphertext using the `Decrypt` method.
*   **Error Handling:** Includes comprehensive error handling to manage potential issues during Vault interactions.

## Contributing

We welcome contributions to VaultCipher! To contribute, please follow these guidelines:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Implement your changes and write tests.
4.  Ensure all tests pass.
5.  Submit a pull request with a clear description of your changes.

Please adhere to the Go coding style and ensure that your code is well-documented.

## License

VaultCipher is licensed under the MIT License. See the LICENSE file for details.