// Copyright 2020 Consensys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by consensys/gnark-crypto DO NOT EDIT

package bandersnatch

import (
	"math/big"
	"sync"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
)

// CurveParams curve parameters: ax^2 + y^2 = 1 + d*x^2*y^2
type CurveParams struct {
	A, D     fr.Element
	Cofactor fr.Element
	Order    big.Int
	Base     PointAffine
	// endomorphism
	endo     [2]fr.Element
	lambda   big.Int
	glvBasis ecc.Lattice
}

// GetEdwardsCurve returns the twisted Edwards curve on bls12-381/Fr
func GetEdwardsCurve() CurveParams {
	initOnce.Do(initCurveParams)
	// copy to keep Order private
	var res CurveParams

	res.A.Set(&curveParams.A)
	res.D.Set(&curveParams.D)
	res.Cofactor.Set(&curveParams.Cofactor)
	res.Order.Set(&curveParams.Order)
	res.Base.Set(&curveParams.Base)
	res.endo[0].Set(&curveParams.endo[0])
	res.endo[1].Set(&curveParams.endo[1])
	res.lambda.Set(&curveParams.lambda)
	res.glvBasis = curveParams.glvBasis // TODO @gbotrel do proper copy of that

	return res
}

var (
	initOnce    sync.Once
	curveParams CurveParams
)

func initCurveParams() {
	curveParams.A.SetString("-5")
	curveParams.D.SetString("45022363124591815672509500913686876175488063829319466900776701791074614335719")
	curveParams.Cofactor.SetString("4")
	curveParams.Order.SetString("13108968793781547619861935127046491459309155893440570251786403306729687672801", 10)

	curveParams.Base.X.SetString("18886178867200960497001835917649091219057080094937609519140440539760939937304")
	curveParams.Base.Y.SetString("19188667384257783945677642223292697773471335439753913231509108946878080696678")
	curveParams.endo[0].SetString("37446463827641770816307242315180085052603635617490163568005256780843403514036")
	curveParams.endo[1].SetString("49199877423542878313146170939139662862850515542392585932876811575731455068989")
	curveParams.lambda.SetString("8913659658109529928382530854484400854125314752504019737736543920008458395397", 10)
	ecc.PrecomputeLattice(&curveParams.Order, &curveParams.lambda, &curveParams.glvBasis)
}

// mulByA multiplies fr.Element by curveParams.A
func mulByA(x *fr.Element) {
	x.Neg(x)
	fr.MulBy5(x)
}
