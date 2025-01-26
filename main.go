package main

import (
	"fmt"
)

// https://midimagic.sgc-hosting.com/quadrafon.htm
// https://midimagic.sgc-hosting.com/quadmath.htm
// SQ :  https://www.4channelsound.com/encode.htm

func encodeSQ(lf, rf, lb, rb complex128) (complex128, complex128) {
	// SQ matrix-based encoding
	const alpha = 0.71
	j := complex(0, 1) // 90° phase-shift

	LT := lf + alpha*rb - j*alpha*lb
	RT := rf - alpha*lb + j*alpha*rb

	return LT, RT
}

func decodeSQ(LT, RT complex128) (complex128, complex128, complex128, complex128) {
	const alpha = 0.71
	j := complex(0, 1) // 90° phase-shift

	lf := LT
	rf := RT
	lb := j*alpha*LT - alpha*RT
	rb := alpha*LT - j*alpha*RT
	return lf, rf, lb, rb

}

/*
// QS : https://en.wikipedia.org/wiki/QS_Regular_Matrix
// Fonction pour l'encodage SQ : Quadriphonique -> Stéréo
func encodeQS(lf, rf, lb, rb complex128) (complex128, complex128) {
	// Encodage basé sur la matrice SQ
	const alpha = 0.924
	const beta = 0.383
	j := complex(0, 1) // 90° phase-shift

	//L := k*lf + 0.7*k*lb + 0.7*rb
	//R := j*rf - 0.7*lb + 0.7*j*rb

	L := alpha*lf + beta*rf + alpha*j*lb + beta*j*rb
	R := beta*lf + alpha*rf - beta*j*lb - alpha*j*rb

	return L, R
}


// Fonction pour le décodage SQ : Stéréo -> Quadriphonique
func decodeQS(L, R complex128) (complex128, complex128, complex128, complex128) {
	const alpha = 0.924
	const beta = 0.383
	j := complex(0, 1) // 90° phase-shift

	// Décodage : on résout les équations inverses
	lf := alpha*L + beta*R
	rf := beta*L + alpha*R
	lb := -j*alpha*L + j*beta*R
	rb := j*beta*L - j*alpha*R
	return lf, rf, lb, rb
}


*/

func main() {

	// Example of continuous quadraphonic signals (Left Front, Right Front, Left Back, Right Back)
	lf := complex(1, 0)   // Avant gauche  / left front
	rf := complex(1, 0)   // Avant droit   / right front
	lb := complex(0.5, 0) // Arrière gauche / left back
	rb := complex(0.5, 0) // Arrière droit / rigt back

	// Print the original values
	fmt.Printf("Original values: lf = %v, rf = %v, lb = %v, rb = %v\n", lf, rf, lb, rb)

	// Two-channel stereo encoding 4-2
	LT, RT := encodeSQ(lf, rf, lb, rb)
	fmt.Printf("Encoded Stereo Channels: LT = %v, RT = %v\n", LT, RT)

	// Decoding 2-4
	lfDecoded, rfDecoded, lbDecoded, rbDecoded := decodeSQ(LT, RT)
	fmt.Printf("Decoded channels: lf = %v, rf = %v, lb = %v, rb = %v\n", lfDecoded, rfDecoded, lbDecoded, rbDecoded)
}
