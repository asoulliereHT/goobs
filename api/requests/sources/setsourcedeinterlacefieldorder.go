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
type SetSourceDeinterlaceFieldOrderParams struct {
	// Name of the source to take a DeinterlaceFieldOrder of
	SourceName string `json:"sourceName,omitempty"`

	SourceDeinterlaceFieldOrder string `json:"sourceDeinterlaceFieldOrder"`
}

// Returns the associated request.
func (o *SetSourceDeinterlaceFieldOrderParams) GetRequestName() string {
	return "SetSourceDeinterlaceFieldOrder"
}

// Represents the response body for the SetSourceDeinterlaceFieldOrder request.
type SetSourceDeinterlaceFieldOrderResponse struct {
	_response
}

/*
Gets a Base64-encoded DeinterlaceFieldOrder of a source.

The `imageWidth` and `imageHeight` parameters are treated as "scale to inner", meaning the smallest ratio will be used and the aspect ratio of the original resolution is kept.
If `imageWidth` and `imageHeight` are not specified, the compressed image will use the full resolution of the source.

**Compatible with inputs and scenes.**
*/
func (c *Client) SetSourceDeinterlaceFieldOrder(params *SetSourceDeinterlaceFieldOrderParams) (*SetSourceDeinterlaceFieldOrderResponse, error) {
	data := &SetSourceDeinterlaceFieldOrderResponse{}
	return data, c.client.SendRequest(params, data)
}
