package call

import (
	"fmt"

	"github.com/ipron-ne/client-sdk-go/types"
	"github.com/pkg/errors"
)

const (
	API_PREFIX  = "/webapi"
	API_MODULE  = "/action/call"
	API_VERSION = "/v1"
	API_NAME    = API_PREFIX + API_MODULE + API_VERSION
)

type Call struct {
	types.Client
}

func NewFromClient(client types.Client) *Call {
	return &Call{
		Client: client,
	}
}

func (c *Call) MakeCall(tntID, userID, callID, ani, dnis, userANI, mediaType string) (types.MakeCallResponse, error) {
	var respData types.MakeCallResponse

	url := fmt.Sprintf("%s/makecall/%s", API_NAME, tntID)
	body := types.MakeCallRequest{
		Tenant:    tntID,
		UserID:    userID,
		CallID:    callID,
		ANI:       ani,
		DNIS:      dnis,
		UserANI:   userANI,
		MediaType: mediaType,
	}

	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCall")
}

func (c *Call) MakeCallEx(tntID, userID, ani, dnis, userANI, mediaType, uei, uui string, routeOption types.RouteOption) (types.MakeCallExResponse, error) {
	var respData types.MakeCallExResponse

	url := fmt.Sprintf("%s/makecallex/%s", API_NAME, tntID)
	body := types.MakeCallExRequest{
		Tenant:    tntID,
		UserID:    userID,
		ANI:       ani,
		DNIS:      dnis,
		UserANI:   userANI,
		MediaType: mediaType,
		UEI:       uei,
		UUI:       uui,
		RouteOpt:  routeOption,
	}

	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) Answer(tntID, callID, connID string) (types.AnswerResponse, error) {
	var respData types.AnswerResponse

	url := fmt.Sprintf("%s/answer/%s", API_NAME, tntID)
	body := types.AnswerRequest{
		Tenant: tntID,
		CallID: callID,
		ConnID: connID,
	}

	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) ReleaseCall(tntID, callID, connID string) (types.ReleaseCallResponse, error) {
	var respData types.ReleaseCallResponse

	url := fmt.Sprintf("%s/release/%s", API_NAME, tntID)
	body := types.ReleaseCallRequest{
		Tenant: tntID,
		CallID: callID,
		ConnID: connID,
	}

	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) Hold(tntID, callID, connID string) (types.HoldResponse, error) {
	var respData types.HoldResponse

	url := fmt.Sprintf("%s/hold/%s", API_NAME, tntID)
	body := types.HoldRequest{
		Tenant: tntID,
		CallID: callID,
		ConnID: connID,
	}

	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) Unhold(tntID, callID, connID string) (types.UnholdResponse, error) {
	var respData types.UnholdResponse

	url := fmt.Sprintf("%s/unhold/%s", API_NAME, tntID)
	body := types.UnholdRequest{
		Tenant: tntID,
		CallID: callID,
		ConnID: connID,
	}

	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) SingleStepTransfer(tntID, callID, connID, dnis, userANI, uui, uei string, routeOption types.RouteOption) (types.SingleStepTransferResponse, error) {
	var respData types.SingleStepTransferResponse

	url := fmt.Sprintf("%s/singlestep-transfer/%s", API_NAME, tntID)
	body := types.SingleStepTransferRequest{
		Tenant:   tntID,
		CallID:   callID,
		ConnID:   connID,
		DNIS:     dnis,
		UserAni:  userANI,
		UUI:      uui,
		UEI:      uei,
		RouteOpt: routeOption,
	}

	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) MuteTransfer(tntID, holdCallID, holdConnID, activeCallID string) (types.MuteTransferResponse, error) {
	var respData types.MuteTransferResponse

	url := fmt.Sprintf("%s/mute-transfer/%s", API_NAME, tntID)
	body := types.MuteTransferRequest{
		Tenant:       tntID,
		HoldCallID:   holdCallID,
		HoldConnID:   holdConnID,
		ActiveCallID: activeCallID,
	}

	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) SingleStepConference(tntID, callID, connID, dnis, userANI, uui, uei, partyType string) (types.SingleStepConferenceResponse, error) {
	var respData types.SingleStepConferenceResponse

	url := fmt.Sprintf("%s/singlestep-conference/%s", API_NAME, tntID)
	body := types.SingleStepConferenceRequest{
		Tenant:    tntID,
		CallID:    callID,
		ConnID:    connID,
		DNIS:      dnis,
		UserAni:   userANI,
		UUI:       uui,
		UEI:       uei,
		PartyType: partyType,
	}

	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) MuteConference(tntID, holdCallID, holdConnID, activeCallID, partyType string) (types.MuteConferenceResponse, error) {
	var respData types.MuteConferenceResponse

	url := fmt.Sprintf("%s/mute-conference/%s", API_NAME, tntID)
	body := types.MuteConferenceRequest{
		Tenant:       tntID,
		HoldCallID:   holdCallID,
		HoldConnID:   holdConnID,
		ActiveCallID: activeCallID,
		PartyType:    partyType,
	}

	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) JoinCall(tntID, userID, joinCallID, joinConnID, joinType string) (types.JoinCallResponse, error) {
	var respData types.JoinCallResponse

	url := fmt.Sprintf("%s/join/%s", API_NAME, tntID)
	body := types.JoinCallRequest{
		Tenant:     tntID,
		UserID:     userID,
		JoinCallID: joinCallID,
		JoinConnID: joinConnID,
		JoinType:   joinType,
	}

	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) Route(tntID, callID, partID, partType, ani, dnis, callbackURI string, timeout int, retryCall, autoAnswer bool) (types.RouteResponse, error) {
	var respData types.RouteResponse

	url := fmt.Sprintf("%s/route/%s", API_NAME, tntID)
	body := types.RouteRequest{
		Tenant:      tntID,
		CallID:      callID,
		PartID:      partID,
		PartType:    partType,
		ANI:         ani,
		DNIS:        dnis,
		CallbackURI: callbackURI,
		Timeout:     timeout,
		RetryCall:   retryCall,
		AutoAnswer:  autoAnswer,
	}

	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) Numberplan(tntID, siteID, dnis string) (types.NumberplanResponse, error) {
	var respData types.NumberplanResponse

	url := fmt.Sprintf("%s/numberplan/%s", API_NAME, tntID)
	body := types.NumberplanRequest{
		Tenant: tntID,
		SiteID: siteID,
		DNIS:   dnis,
	}

	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) SetUserdata(tntID, callID, uei, uui string) (types.SetUserdataResponse, error) {
	var respData types.SetUserdataResponse

	url := fmt.Sprintf("%s/userdata/%s", API_NAME, tntID)
	body := types.SetUserdataRequest{
		Tenant: tntID,
		CallID: callID,
		UEI:    uei,
		UUI:    uui,
	}

	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) GetUserdata(tntID, callID string) (types.GetUserdataResponse, error) {
	var respData types.GetUserdataResponse

	url := fmt.Sprintf("%s/userdata/%s/%s", API_NAME, tntID, callID)
	resp, err := c.GetRequest().Post(url, nil)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}
