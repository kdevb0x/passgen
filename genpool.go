package main

import ()

// GenPool is a collection pool for for async generators
func GenPool(sndr ...<-chan []byte) error {

}

func asyncRandomL() <-chan []byte {
	lChan := make(chan []byte, len(randomL), 10)

	var genfunc = func() []byte {

	}
}

func asyncRandomU() <-chan []byte {

}

func asyncRandomN() <-chan []byte {

}

func asyncRandomS() <-chan []byte {

}
