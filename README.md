# Zero-Knowledge SNARK Example

This repository contains a simple example of using zero-knowledge proofs with the `gnark` library to prove knowledge of a witness for a cubic equation and then verify the proof. It demonstrates the process of setting up a circuit, compiling it, generating proofs, and verifying them using zk-SNARKs.

## Requirements

- Go (Golang)
- `github.com/consensys/gnark-crypto/ecc`
- `github.com/consensys/gnark/backend/groth16`
- `github.com/consensys/gnark/frontend`
- `github.com/consensys/gnark/frontend/cs/r1cs`

## Usage

1. Clone this repository:

```shell
git clone git@github.com:raadhshenshahhaseeb/go-snark.git
```

2. Build and run the `Snark` function:

```shell
go mod tidy
go run cmd/main.go
```

This will compile the circuit, generate a proof, and verify it using zk-SNARKs. The program will output whether the verification succeeded or failed.

## Code Explanation

The code consists of the following parts:

- `CircuitCubic`: A struct that defines the circuit constraints for the cubic equation.
- `Define`: A function that declares the circuit constraints based on the provided equation.
- `Snark`: A function that demonstrates the entire process of circuit instantiation, proof generation, and verification.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

This code is based on the `gnark` library example. Credit goes to the `gnark` development team for providing the foundation for this example.

---
