package go_snark

import (
	"fmt"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

// Circuit defines a simple circuit
type CircuitCubic struct {
	// struct tags on a variable is optional
	// default uses variable name and secret visibility.
	X frontend.Variable `gnark:"x"`
	Y frontend.Variable `gnark:",public"`
}

// Define declares the circuit constraints
// x**3 + x**2 + 5(x) + 5 == y
func (circuit *CircuitCubic) Define(api frontend.API) error {
	x3 := api.Mul(circuit.X, circuit.X, circuit.X)
	x2 := api.Mul(circuit.X, circuit.X)
	x1 := api.Mul(circuit.X, 5)
	api.AssertIsEqual(circuit.Y, api.Add(x3, x2, x1, 5))
	return nil
}

func Snark() {
	// step1. instantiate circuit
	var cubicCircuit CircuitCubic
	r1cs, err := frontend.Compile(ecc.BN254, r1cs.NewBuilder, &cubicCircuit)
	if err != nil {
		fmt.Errorf("err")
		return
	}

	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		return
	}

	// step2. export the groth16.VerifyingKey as a solidity smart contract.
	var fileName = "./verify/verifyCubicEqual.sol"

	_, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	var solidityFile, _ = os.Create(fileName)
	vk.ExportSolidity(solidityFile)

	// step3. generate witness and prove
	assignment := &CircuitCubic{
		X: frontend.Variable(2),
		Y: frontend.Variable(27),
	}

	witness, _ := frontend.NewWitness(assignment, ecc.BN254)

	proof, err := groth16.Prove(r1cs, pk, witness)
	if err != nil {
		return
	}

	// step4. generate public witness and verify
	validPublicWitness, _ := frontend.NewWitness(assignment, ecc.BN254, frontend.PublicOnly())
	err = groth16.Verify(proof, vk, validPublicWitness)
	if err != nil {
		fmt.Printf("verification failed\n")
		return
	}
	fmt.Printf("verification succeded\n")
}
