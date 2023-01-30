package test

import (
	"github.com/grokkos/finnio/internal/app"
	"math"
	"testing"
)

//TODO mock the finnio http requests so integration test can created to test the whole flow with calculations

func TestCalcSimpleProfit(t *testing.T) {
	gotLoss, gotProfit := app.CalcProfitAndLoss(10.4, 9.5)
	gotProfit = float32(math.Ceil(float64(gotProfit*100)) / 100)
	wantLoss := float32(0)
	wantProfit := float32(9)

	if gotProfit != wantProfit || gotLoss != wantLoss {
		t.Errorf("gotLoss %+v,wantLoss %+v, gotProfit %+v, wantProfit %+v", gotLoss, wantLoss, gotProfit, wantProfit)
	}
}

func TestCalcSimpleLoss(t *testing.T) {
	gotLoss, gotProfit := app.CalcProfitAndLoss(11, 13.75)
	gotLoss = float32(math.Ceil(float64(gotLoss*100)) / 100)
	wantLoss := float32(27.5)
	wantProfit := float32(0)

	if gotProfit != wantProfit || gotLoss != wantLoss {
		t.Errorf("gotLoss %+v,wantLoss %+v, gotProfit %+v, wantProfit %+v", gotLoss, wantLoss, gotProfit, wantProfit)
	}
}
