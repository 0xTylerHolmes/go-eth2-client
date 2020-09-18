// Copyright © 2020 Attestant Limited.
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

package lighthousehttp

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

// CurrentEpoch is a helper that calculates the current epoch.
func (s *Service) CurrentEpoch(ctx context.Context) (uint64, error) {
	genesisTime, err := s.GenesisTime(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to obtain genesis time for current epoch")
	}
	slotDuration, err := s.SlotDuration(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to obtain slot duration for current epoch")
	}
	slotsPerEpoch, err := s.SlotsPerEpoch(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to obtain slots per epoch for current epoch")
	}
	var currentEpoch uint64
	if genesisTime.After(time.Now()) {
		currentEpoch = 0
	} else {
		currentEpoch = uint64(time.Since(genesisTime).Seconds()) / (uint64(slotDuration.Seconds()) * slotsPerEpoch)
	}

	return currentEpoch, nil
}

// CurrentSlot is a helper that calculates the current slot.
func (s *Service) CurrentSlot(ctx context.Context) (uint64, error) {
	genesisTime, err := s.GenesisTime(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to obtain genesis time for current slot")
	}
	slotDuration, err := s.SlotDuration(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to obtain slot duration for current slot")
	}
	var currentSlot uint64
	if genesisTime.After(time.Now()) {
		currentSlot = 0
	} else {
		currentSlot = uint64(time.Since(genesisTime).Seconds()) / uint64(slotDuration.Seconds())
	}

	return currentSlot, nil
}

//func (s *Service) beaconHead(ctx context.Context) ([]byte, []byte, uint64, error) {
//	respBodyReader, err := s.get(ctx, "/beacon/head")
//	if err != nil {
//		log.Trace().Err(err).Msg("Request failed")
//		return nil, nil, 0, errors.Wrap(err, "failed to request beacon head")
//	}
//	specReader, err := lhToSpec(ctx, respBodyReader)
//	beaconHeadResponse := &struct {
//		Slot      string `json:"slot"`
//		StateRoot string `json:"state_root"`
//		BlockRoot string `json:"block_root"`
//	}{}
//	if err := json.NewDecoder(specReader).Decode(&beaconHeadResponse); err != nil {
//		return nil, nil, 0, errors.Wrap(err, "failed to parse beacon head")
//	}
//
//	slot, err := strconv.ParseUint(beaconHeadResponse.Slot, 10, 64)
//	if err != nil {
//		return nil, nil, 0, errors.Wrap(err, "failed to parse beacon head slot")
//	}
//	stateRoot, err := hex.DecodeString(strings.TrimPrefix(beaconHeadResponse.StateRoot, "0x"))
//	if err != nil {
//		return nil, nil, 0, errors.Wrap(err, "failed to parse beacon head state root")
//	}
//	blockRoot, err := hex.DecodeString(strings.TrimPrefix(beaconHeadResponse.BlockRoot, "0x"))
//	if err != nil {
//		return nil, nil, 0, errors.Wrap(err, "failed to parse beacon head block root")
//	}
//
//	return stateRoot, blockRoot, slot, nil
//}
