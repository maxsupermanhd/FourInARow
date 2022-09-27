package main

var randSeed = uint64(0)
var randGamma = uint64(0x00281E2DBA6606F3)
var randGammaPrime = uint64((1 << 56) - 5)

// var randNextSplit = uint64(0)
var randDoubleUlp = float64(1.0 / (1 << 53))

func rndMix56(z uint64) uint64 {
	z = (uint64(z^(z>>33)) * uint64(0xff51afd7ed558ccd)) & 0x00FFFFFFFFFFFFFF
	z = (uint64(z^(z>>33)) * uint64(0xc4ceb9fe1a85ec53)) & 0x00FFFFFFFFFFFFFF
	return z ^ (z >> 33)
}

func rndSetSeed(seed uint64) {
	randSeed = seed
	s := randGamma
	if s >= randGammaPrime {
		s -= randGammaPrime
	}
	randGamma = rndMix56(s) + 13
	// randNextSplit = s
}

func rndMix64(z uint64) uint64 {
	z = uint64(z^(z>>33)) * uint64(0xff51afd7ed558ccd)
	z = uint64(z^(z>>33)) * uint64(0xc4ceb9fe1a85ec53)
	return z ^ (z >> 33)
}

func rndUpdate(s, g uint64) uint64 {
	p := s * g
	if p >= s {
		return p
	} else {
		if p >= uint64(0x800000000000000D) {
			return p - 13
		} else {
			return (p - 13) + g
		}
	}
}

func rndNextRaw64() uint64 {
	randSeed = rndUpdate(randSeed, randGamma)
	return randSeed
}

func rndNextLong() uint64 {
	return rndMix64(rndNextRaw64())
}

func rndNextDouble() float64 {
	return float64(rndNextLong()>>11) * randDoubleUlp
}

func bRND(i int) float64 {
	return rndNextDouble()
}
