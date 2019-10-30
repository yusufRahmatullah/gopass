package gopass

import (
	"testing"
)

func TestGenPass(t *testing.T) {
	cases := map[string]string{
		"":             "#4&14*ZYr-*34Cc#4I_Qm^!N*!UdW^^zhn4T9p&1m$Q8bCRxGgB$bEN-F69%^WO#",
		"::":           "TE6c5mr-25@!-n0rVM!1Fc*!@^T2@V#c4476v78Hl8$3Z%3Q3z0e8^i3Mev9Y*S*",
		"basic::basic": "_6_83J#^R$w7$nO%ERpc$#3!Ejn_3f4nw%fJRB2r8p^5&_32*#34yM!_&-A-scPo",
		"lala::yeye":   "*d7C&nS#5i9p16-22r&70GHYM@5#LQ%7D^02kmK_08!uLmF#-03w2-yWFYsjnr%5",
		"LongEnoughMasterKey::long enough purpose":                         "5OLui5F^*U27Z&cH1!11wl0Ds8U*!2w$j-Na$9e7Pi$*-w@sc1%J@MMZMf#16N@W",
		"TE6c5mr-25@!-n0rVM!1Fc*!@^T2@V#c4476v78Hl8$3Z%3Q3z0e8^i3Mev9Y*S*": "Nix65xl6nuF_6Z*!%3!6Tf8n%D1o6V-*ra&C!cp08#49V9ye6w*8J8o9FS3Xm!0-",
	}
	for seed, exp := range cases {
		got := GenPass(seed)
		if exp != got {
			t.Errorf("exp = %v, got = %v", exp, got)
		}
	}
}

func TestGenPin(t *testing.T) {
	cases := map[string]string{
		"":             "227416",
		"::":           "676176",
		"basic::basic": "098609",
		"lala::yeye":   "072187",
		"LongEnoughMasterKey::long enough purpose":                         "755554",
		"TE6c5mr-25@!-n0rVM!1Fc*!@^T2@V#c4476v78Hl8$3Z%3Q3z0e8^i3Mev9Y*S*": "652339",
	}
	for seed, exp := range cases {
		got := GenPin(seed)
		if exp != got {
			t.Errorf("exp = %v, got = %v", exp, got)
		}
	}
}
