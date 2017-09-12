package claim

import (
	"testing"
	"encoding/hex"
)

var raw_claims = []string{
"08011002225e0801100322583056301006072a8648ce3d020106052b8104000a03420004d015365a40f3e5c03c87227168e5851f44659837bcf6a3398ae633bc37d04ee19baeb26dc888003bd728146dbea39f5344bf8c52cedaf1a3a1623a0166f4a367",
"080110011ad7010801128f01080410011a0c47616d65206f66206c696665221047616d65206f66206c696665206769662a0b4a6f686e20436f6e776179322e437265617469766520436f6d6d6f6e73204174747269627574696f6e20342e3020496e7465726e6174696f6e616c38004224080110011a195569c917f18bf5d2d67f1346aa467b218ba90cdbf2795676da250000803f4a0052005a001a41080110011a30b6adf6e2a62950407ea9fb045a96127b67d39088678d2f738c359894c88d95698075ee6203533d3c204330713aa7acaf2209696d6167652f6769662a5c080110031a40c73fe1be4f1743c2996102eec6ce0509e03744ab940c97d19ddb3b25596206367ab1a3d2583b16c04d2717eeb983ae8f84fee2a46621ffa5c4726b30174c6ff82214251305ca93d4dbedb50dceb282ebcb7b07b7ac65",
"080110011ad7010801128f01080410011a0c47616d65206f66206c696665221047616d65206f66206c696665206769662a0b4a6f686e20436f6e776179322e437265617469766520436f6d6d6f6e73204174747269627574696f6e20342e3020496e7465726e6174696f6e616c38004224080110011a195569c917f18bf5d2d67f1346aa467b218ba90cdbf2795676da250000803f4a0052005a001a41080110011a30b6adf6e2a62950407ea9fb045a96127b67d39088678d2f738c359894c88d95698075ee6203533d3c204330713aa7acaf2209696d6167652f676966",
}


func TestDecodeClaims(t *testing.T) {
	for _, claim_hex := range(raw_claims) {
		claim := Claim{}
		err := claim.LoadFromHexString(claim_hex)
		if err != nil {
			t.Error(err)
		}
		err = claim.LoadFromHexString(claim_hex)
		if err.Error() != "already initialized" {
			t.Error(err)
		}
		serialized_hex, err := claim.SerializedHexString()
		if err != nil {
			t.Error(err)
		}
		if serialized_hex != claim_hex {
			t.Error("failed to re-serialize")
		}
	}
}

func TestStripSignature(t *testing.T) {
	claim_hex := raw_claims[1]
	claim := Claim{}
	err := claim.LoadFromHexString(claim_hex)
	if err != nil {
		t.Error(err)
	}
	no_sig, err := claim.SerializedNoSignature()
	if err != nil {
		t.Error(err)
	}
	if hex.EncodeToString(no_sig) != raw_claims[2] {
		t.Error("Failed to remove signature")
	}
}
