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

func (c *Call) MakeCall(tntID, userID, callID, ani, dnis, userANI, mediaType string) (MakeCallResponse, error) {
	var respData MakeCallResponse

	url := fmt.Sprintf("%s/makecall/%s", API_NAME, tntID)
	body := MakeCallRequest{
		Tenant:    tntID,
		UserID:    userID,
		CallID:    callID,
		ANI:       ani,
		DNIS:      dnis,
		UserANI:   userANI,
		MediaType: mediaType,
	}

	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCall")
}

func (c *Call) MakeCallEx(tntID, userID, ani, dnis, userANI, mediaType, uei, uui string, routeOption RouteOption) (MakeCallExResponse, error) {
	var respData MakeCallExResponse

	url := fmt.Sprintf("%s/makecallex/%s", API_NAME, tntID)
	body := MakeCallExRequest{
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
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) Answer(tntID, callID, connID string) (AnswerResponse, error) {
	var respData AnswerResponse

	url := fmt.Sprintf("%s/answer/%s", API_NAME, tntID)
	body := AnswerRequest{
		Tenant: tntID,
		CallID: callID,
		ConnID: connID,
	}

	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) ReleaseCall(tntID, callID, connID string) (ReleaseCallResponse, error) {
	var respData ReleaseCallResponse

	url := fmt.Sprintf("%s/release/%s", API_NAME, tntID)
	body := ReleaseCallRequest{
		Tenant: tntID,
		CallID: callID,
		ConnID: connID,
	}

	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) Hold(tntID, callID, connID string) (HoldResponse, error) {
	var respData HoldResponse

	url := fmt.Sprintf("%s/hold/%s", API_NAME, tntID)
	body := HoldRequest{
		Tenant: tntID,
		CallID: callID,
		ConnID: connID,
	}

	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) Unhold(tntID, callID, connID string) (UnholdResponse, error) {
	var respData UnholdResponse

	url := fmt.Sprintf("%s/unhold/%s", API_NAME, tntID)
	body := UnholdRequest{
		Tenant: tntID,
		CallID: callID,
		ConnID: connID,
	}

	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) SingleStepTransfer(tntID, callID, connID, dnis, userANI, uui, uei string, routeOption RouteOption) (SingleStepTransferResponse, error) {
	var respData SingleStepTransferResponse

	url := fmt.Sprintf("%s/singlestep-transfer/%s", API_NAME, tntID)
	body := SingleStepTransferRequest{
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
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) MuteTransfer(tntID, holdCallID, holdConnID, activeCallID string) (MuteTransferResponse, error) {
	var respData MuteTransferResponse

	url := fmt.Sprintf("%s/mute-transfer/%s", API_NAME, tntID)
	body := MuteTransferRequest{
		Tenant:       tntID,
		HoldCallID:   holdCallID,
		HoldConnID:   holdConnID,
		ActiveCallID: activeCallID,
	}

	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) SingleStepConference(tntID, callID, connID, dnis, userANI, uui, uei, partyType string) (SingleStepConferenceResponse, error) {
	var respData SingleStepConferenceResponse

	url := fmt.Sprintf("%s/singlestep-conference/%s", API_NAME, tntID)
	body := SingleStepConferenceRequest{
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
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) MuteConference(tntID, holdCallID, holdConnID, activeCallID, partyType string) (MuteConferenceResponse, error) {
	var respData MuteConferenceResponse

	url := fmt.Sprintf("%s/mute-conference/%s", API_NAME, tntID)
	body := MuteConferenceRequest{
		Tenant:       tntID,
		HoldCallID:   holdCallID,
		HoldConnID:   holdConnID,
		ActiveCallID: activeCallID,
		PartyType:    partyType,
	}

	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) JoinCall(tntID, userID, joinCallID, joinConnID, joinType string) (JoinCallResponse, error) {
	var respData JoinCallResponse

	url := fmt.Sprintf("%s/join/%s", API_NAME, tntID)
	body := JoinCallRequest{
		Tenant:     tntID,
		UserID:     userID,
		JoinCallID: joinCallID,
		JoinConnID: joinConnID,
		JoinType:   joinType,
	}

	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) Route(tntID, callID, partID, partType, ani, dnis, callbackURI string, timeout int, retryCall, autoAnswer bool) (RouteResponse, error) {
	var respData RouteResponse

	url := fmt.Sprintf("%s/route/%s", API_NAME, tntID)
	body := RouteRequest{
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
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) Numberplan(tntID, siteID, dnis string) (NumberplanResponse, error) {
	var respData NumberplanResponse

	url := fmt.Sprintf("%s/numberplan/%s", API_NAME, tntID)
	body := NumberplanRequest{
		Tenant: tntID,
		SiteID: siteID,
		DNIS:   dnis,
	}

	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) SetUserdata(tntID, callID, uei, uui string) (SetUserdataResponse, error) {
	var respData SetUserdataResponse

	url := fmt.Sprintf("%s/userdata/%s", API_NAME, tntID)
	body := SetUserdataRequest{
		Tenant: tntID,
		CallID: callID,
		UEI:    uei,
		UUI:    uui,
	}

	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

func (c *Call) GetUserdata(tntID, callID string) (GetUserdataResponse, error) {
	var respData GetUserdataResponse

	url := fmt.Sprintf("%s/userdata/%s/%s", API_NAME, tntID, callID)
	resp, err := c.GetRequest().Get(url, nil)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}

// Supported v1.2
func (c *Call) RemoteTransfer(tntID, callID, connID, dnis, userANI, uui string) (RemoteTransferResponse, error) {
	var respData RemoteTransferResponse

	url := fmt.Sprintf("%s/remote-transfer/%s", API_NAME, tntID)
	body := RemoteTransferRequest{
		Tenant:  tntID,
		CallID:  callID,
		ConnID:  connID,
		DNIS:    dnis,
		UserANI: userANI,
		UUI:     uui,
	}

	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "MakeCallEx")
}
