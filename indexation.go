package iotago

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/iotaledger/hive.go/v2/serializer"
)

const (
	// IndexationBinSerializedMinSize is the minimum size of an Indexation.
	// 	type bytes + index prefix + one char + data length
	IndexationBinSerializedMinSize = serializer.TypeDenotationByteSize + serializer.UInt16ByteSize + serializer.OneByte + serializer.UInt32ByteSize
	// IndexationIndexMaxLength defines the max length of the index within an Indexation.
	IndexationIndexMaxLength = 64
	// IndexationIndexMinLength defines the min length of the index within an Indexation.
	IndexationIndexMinLength = 1
)

var (
	// ErrIndexationIndexExceedsMaxSize gets returned when an Indexation's index exceeds IndexationIndexMaxLength.
	ErrIndexationIndexExceedsMaxSize = errors.New("index exceeds max size")
	// ErrIndexationIndexUnderMinSize gets returned when an Indexation's index is under IndexationIndexMinLength.
	ErrIndexationIndexUnderMinSize = errors.New("index is below min size")
)

// Indexation is a payload which holds an index and associated data.
type Indexation struct {
	// The index to use to index the enclosing message and data.
	Index []byte `json:"index"`
	// The data within the payload.
	Data []byte `json:"data"`
}

func (u *Indexation) PayloadType() PayloadType {
	return PayloadIndexation
}

func (u *Indexation) Deserialize(data []byte, deSeriMode serializer.DeSerializationMode, deSeriCtx interface{}) (int, error) {
	return serializer.NewDeserializer(data).
		CheckTypePrefix(uint32(PayloadIndexation), serializer.TypeDenotationUint32, func(err error) error {
			return fmt.Errorf("unable to deserialize indexation: %w", err)
		}).
		ReadVariableByteSlice(&u.Index, serializer.SeriLengthPrefixTypeAsUint16, func(err error) error {
			return fmt.Errorf("unable to deserialize indexation index: %w", err)
		}, IndexationIndexMaxLength).
		WithValidation(deSeriMode, func(_ []byte,err error) error {
			switch {
			case len(u.Index) < IndexationIndexMinLength:
				return fmt.Errorf("unable to deserialize indexation index: %w", ErrIndexationIndexUnderMinSize)
			}
			return nil
		}).
		ReadVariableByteSlice(&u.Data, serializer.SeriLengthPrefixTypeAsUint32, func(err error) error {
			return fmt.Errorf("unable to deserialize indexation data: %w", err)
		}, MessageBinSerializedMaxSize). // obviously can never be that size
		Done()
}

func (u *Indexation) Serialize(deSeriMode serializer.DeSerializationMode, deSeriCtx interface{}) ([]byte, error) {
	return serializer.NewSerializer().
		WithValidation(deSeriMode, func(_ []byte,err error) error {
			if deSeriMode.HasMode(serializer.DeSeriModePerformValidation) {
				switch {
				case len(u.Index) > IndexationIndexMaxLength:
					return fmt.Errorf("unable to serialize indexation index: %w", ErrIndexationIndexExceedsMaxSize)
				case len(u.Index) < IndexationIndexMinLength:
					return fmt.Errorf("unable to serialize indexation index: %w", ErrIndexationIndexUnderMinSize)
				}
				// we do not check the length of the data field as in any circumstance
				// the max size it can take up is dependent on how big the enclosing
				// parent object is
			}
			return nil
		}).
		WriteNum(PayloadIndexation, func(err error) error {
			return fmt.Errorf("unable to serialize indexation payload ID: %w", err)
		}).
		WriteVariableByteSlice(u.Index, serializer.SeriLengthPrefixTypeAsUint16, func(err error) error {
			return fmt.Errorf("unable to serialize indexation index: %w", err)
		}).
		WriteVariableByteSlice(u.Data, serializer.SeriLengthPrefixTypeAsUint32, func(err error) error {
			return fmt.Errorf("unable to serialize indexation data: %w", err)
		}).
		Serialize()
}

func (u *Indexation) MarshalJSON() ([]byte, error) {
	jIndexation := &jsonIndexation{}
	jIndexation.Type = int(PayloadIndexation)
	jIndexation.Index = hex.EncodeToString(u.Index)
	jIndexation.Data = hex.EncodeToString(u.Data)
	return json.Marshal(jIndexation)
}

func (u *Indexation) UnmarshalJSON(bytes []byte) error {
	jIndexation := &jsonIndexation{}
	if err := json.Unmarshal(bytes, jIndexation); err != nil {
		return err
	}
	seri, err := jIndexation.ToSerializable()
	if err != nil {
		return err
	}
	*u = *seri.(*Indexation)
	return nil
}

// jsonIndexation defines the json representation of an Indexation.
type jsonIndexation struct {
	Type  int    `json:"type"`
	Index string `json:"index"`
	Data  string `json:"data"`
}

func (j *jsonIndexation) ToSerializable() (serializer.Serializable, error) {
	indexBytes, err := hex.DecodeString(j.Index)
	if err != nil {
		return nil, fmt.Errorf("unable to decode index from JSON for indexation: %w", err)
	}

	dataBytes, err := hex.DecodeString(j.Data)
	if err != nil {
		return nil, fmt.Errorf("unable to decode data from JSON for indexation: %w", err)
	}

	return &Indexation{Index: indexBytes, Data: dataBytes}, nil
}
