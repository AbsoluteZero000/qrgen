# QRCodeGen CLI

A lightweight, command-line QR code generator tool for encoding text into QR codes. This utility is simple to use, allowing you to quickly generate a QR code with a single command.

## Features
- Generate QR codes directly from the command line.
- Minimal dependencies and fast execution.
- Supports alphanumeric text and symbols.

## Installation
Clone the repository and build the tool:

```bash
git clone <repository-url>
cd qrgen
go build -o qrgen
```

Ensure the binary is in your PATH for easy access:

```bash
sudo mv qrgen /usr/local/bin
```

## Usage
```bash
qrgen [text to encode]
```

### Example:
```bash
qrgen "Hello, World!"
```

This command will generate a QR code representing the text "Hello, World!" and output it to your terminal or save it as a file (based on your implementation).

## How It Works
This tool follows the basic principles of QR code generation outlined in the [Creating a QR Code step-by-step](https://www.nayuki.io/page/creating-a-qr-code-step-by-step) guide. Understanding the structure of QR codes ensures that the tool generates valid and efficient codes.

## Contributing
Feel free to submit issues or pull requests. All contributions are welcome!

## License
This project is licensed under the MIT License.
