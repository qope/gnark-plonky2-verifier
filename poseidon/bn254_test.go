package poseidon

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
	gl "github.com/qope/gnark-plonky2-verifier/goldilocks"
)

type TestPoseidonBN254Circuit struct {
	In  [BN254_SPONGE_WIDTH]frontend.Variable
	Out [BN254_SPONGE_WIDTH]frontend.Variable
}

func (circuit *TestPoseidonBN254Circuit) Define(api frontend.API) error {
	poseidonChip := NewBN254Chip(api)
	output := poseidonChip.Poseidon(circuit.In)

	for i := 0; i < BN254_SPONGE_WIDTH; i++ {
		api.AssertIsEqual(
			output[i],
			circuit.Out[i],
		)
	}

	return nil
}

func TestPoseidonBN254(t *testing.T) {
	assert := test.NewAssert(t)

	testCaseFn := func(in [BN254_SPONGE_WIDTH]frontend.Variable, out [BN254_SPONGE_WIDTH]frontend.Variable) {
		circuit := TestPoseidonBN254Circuit{In: in, Out: out}
		witness := TestPoseidonBN254Circuit{In: in, Out: out}
		err := test.IsSolved(&circuit, &witness, ecc.BN254.ScalarField())
		assert.NoError(err)
	}

	testCases := [][2][]string{
		{
			{"0", "0", "0", "0"},
			{
				"5317387130258456662214331362918410991734007599705406860481038345552731150762",
				"17768273200467269691696191901389126520069745877826494955630904743826040320364",
				"19413739268543925182080121099097652227979760828059217876810647045303340666757",
				"3717738800218482999400886888123026296874264026760636028937972004600663725187",
			},
		},
		{
			{"0", "1", "2", "3"},
			{
				"6542985608222806190361240322586112750744169038454362455181422643027100751666",
				"3478427836468552423396868478117894008061261013954248157992395910462939736589",
				"1904980799580062506738911865015687096398867595589699208837816975692422464009",
				"11971464497515232077059236682405357499403220967704831154657374522418385384151",
			},
		},
		{
			{
				"21888242871839275222246405745257275088548364400416034343698204186575808495616",
				"21888242871839275222246405745257275088548364400416034343698204186575808495616",
				"21888242871839275222246405745257275088548364400416034343698204186575808495616",
				"21888242871839275222246405745257275088548364400416034343698204186575808495616",
			},
			{
				"13055670547682322550638362580666986963569035646873545133474324633020685301274",
				"19087936485076376314486368416882351797015004625427655501762827988254486144933",
				"10391468779200270580383536396630001155994223659670674913170907401637624483385",
				"17202557688472898583549180366140168198092766974201433936205272956998081177816",
			},
		},
		{
			{
				"6542985608222806190361240322586112750744169038454362455181422643027100751666",
				"3478427836468552423396868478117894008061261013954248157992395910462939736589",
				"1904980799580062506738911865015687096398867595589699208837816975692422464009",
				"11971464497515232077059236682405357499403220967704831154657374522418385384151",
			},
			{
				"21792249080447013894140672594027696524030291802493510986509431008224624594361",
				"3536096706123550619294332177231935214243656967137545251021848527424156573335",
				"14869351042206255711434675256184369368509719143073814271302931417334356905217",
				"5027523131326906886284185656868809493297314443444919363729302983434650240523",
			},
		},
	}

	for _, testCase := range testCases {
		var in [BN254_SPONGE_WIDTH]frontend.Variable
		var out [BN254_SPONGE_WIDTH]frontend.Variable
		copy(in[:], gl.StrArrayToFrontendVariableArray(testCase[0]))
		copy(out[:], gl.StrArrayToFrontendVariableArray(testCase[1]))
		testCaseFn(in, out)
	}
}
