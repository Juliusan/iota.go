package iotago

import (
	"encoding/json"
	"fmt"

	"github.com/iotaledger/hive.go/v2/serializer"
)

// TreasuryOutput is an output which holds the treasury of a network.
type TreasuryOutput struct {
	// The currently residing funds in the treasury.
	Amount uint64 `json:"deposit"`
}

func (t *TreasuryOutput) VByteCost(_ *RentStructure, _ VByteCostFunc) uint64 {
	return 0
}

func (t *TreasuryOutput) Deposit() uint64 {
	return t.Amount
}

func (t *TreasuryOutput) Type() OutputType {
	return OutputTreasury
}

func (t *TreasuryOutput) Deserialize(data []byte, _ serializer.DeSerializationMode, _ interface{}) (int, error) {
	return serializer.NewDeserializer(data).
		CheckTypePrefix(uint32(OutputTreasury), serializer.TypeDenotationByte, func(err error) error {
			return fmt.Errorf("unable to deserialize treasury output: %w", err)
		}).
		ReadNum(&t.Amount, func(err error) error {
			return fmt.Errorf("unable to deserialize amount for treasury output: %w", err)
		}).
		Done()
}

func (t *TreasuryOutput) Serialize(_ serializer.DeSerializationMode, _ interface{}) ([]byte, error) {
	return serializer.NewSerializer().
		WriteNum(OutputTreasury, func(err error) error {
			return fmt.Errorf("unable to serialize treasury output type ID: %w", err)
		}).
		WriteNum(t.Amount, func(err error) error {
			return fmt.Errorf("unable to serialize treasury output amount: %w", err)
		}).Serialize()
}

func (t *TreasuryOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(&jsonTreasuryOutput{
		Type:   int(OutputTreasury),
		Amount: int(t.Amount),
	})
}

func (t *TreasuryOutput) UnmarshalJSON(bytes []byte) error {
	jTreasuryOutput := &jsonTreasuryOutput{}
	if err := json.Unmarshal(bytes, jTreasuryOutput); err != nil {
		return err
	}
	seri, err := jTreasuryOutput.ToSerializable()
	if err != nil {
		return err
	}
	*t = *seri.(*TreasuryOutput)
	return nil
}

// jsonTreasuryOutput defines the json representation of a TreasuryOutput.
type jsonTreasuryOutput struct {
	Type   int `json:"type"`
	Amount int `json:"amount"`
}

func (j *jsonTreasuryOutput) ToSerializable() (serializer.Serializable, error) {
	return &TreasuryOutput{Amount: uint64(j.Amount)}, nil
}
