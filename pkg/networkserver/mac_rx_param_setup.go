// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
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

package networkserver

import (
	"context"

	"go.thethings.network/lorawan-stack/pkg/errors/common"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
)

func handleRxParamSetupAns(ctx context.Context, dev *ttnpb.EndDevice, pld *ttnpb.MACCommand_RxParamSetupAns) (err error) {
	if pld == nil {
		return common.ErrMissingPayload.New(nil)
	}

	dev.PendingMACRequests, err = handleMACResponse(ttnpb.CID_RX_PARAM_SETUP, func(cmd *ttnpb.MACCommand) {
		if !pld.Rx1DataRateOffsetAck || !pld.Rx2DataRateIndexAck || !pld.Rx2FrequencyAck {
			// TODO: Handle NACK, modify desired state
			// (https://github.com/TheThingsIndustries/ttn/issues/834)
			return
		}

		req := cmd.GetRxParamSetupReq()

		dev.MACState.Rx1DataRateOffset = req.Rx1DataRateOffset
		dev.MACState.Rx2DataRateIndex = req.Rx2DataRateIndex
		dev.MACState.Rx2Frequency = req.Rx2Frequency

	}, dev.PendingMACRequests...)
	return
}
