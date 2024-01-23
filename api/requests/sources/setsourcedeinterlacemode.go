/**
 * Deinterlace modes:
 *
 * - `OBS_DEINTERLACE_MODE_DISABLE`
 * - `OBS_DEINTERLACE_MODE_DISCARD`
 * - `OBS_DEINTERLACE_MODE_RETRO`
 * - `OBS_DEINTERLACE_MODE_BLEND`
 * - `OBS_DEINTERLACE_MODE_BLEND_2X`
 * - `OBS_DEINTERLACE_MODE_LINEAR`
 * - `OBS_DEINTERLACE_MODE_LINEAR_2X`
 * - `OBS_DEINTERLACE_MODE_YADIF`
 * - `OBS_DEINTERLACE_MODE_YADIF_2X`
 *
 */

package sources

// Represents the request body for the SetSourceDeinterlaceMode request.
type SetSourceDeinterlaceModeParams struct {
	// Name of the source to take a DeinterlaceMode of
	SourceName string `json:"sourceName,omitempty"`

	SourceDeinterlaceMode string `json:"sourceDeinterlaceMode"`
}

// Returns the associated request.
func (o *SetSourceDeinterlaceModeParams) GetRequestName() string {
	return "SetSourceDeinterlaceMode"
}

// Represents the response body for the SetSourceDeinterlaceMode request.
type SetSourceDeinterlaceModeResponse struct {
	_response
}

/*
Gets a Base64-encoded DeinterlaceMode of a source.

The `imageWidth` and `imageHeight` parameters are treated as "scale to inner", meaning the smallest ratio will be used and the aspect ratio of the original resolution is kept.
If `imageWidth` and `imageHeight` are not specified, the compressed image will use the full resolution of the source.

**Compatible with inputs and scenes.**
*/
func (c *Client) SetSourceDeinterlaceMode(params *SetSourceDeinterlaceModeParams) (*SetSourceDeinterlaceModeResponse, error) {
	data := &SetSourceDeinterlaceModeResponse{}
	return data, c.client.SendRequest(params, data)
}
