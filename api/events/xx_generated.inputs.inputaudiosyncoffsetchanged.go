// This file has been automatically generated. Don't edit it.

package events

// Represents the event body for the InputAudioSyncOffsetChanged event.
type InputAudioSyncOffsetChanged struct {
	// New sync offset in milliseconds
	InputAudioSyncOffset float64 `json:"inputAudioSyncOffset,omitempty"`

	// Name of the input
	InputName string `json:"inputName,omitempty"`
}
