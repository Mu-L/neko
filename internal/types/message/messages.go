package message

import (
	"github.com/pion/webrtc/v3"

	"demodesk/neko/internal/types"
)

type Message struct {
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"` // TODO: New.
}

/////////////////////////////
// System
/////////////////////////////

type SystemWebRTC struct {
	Event  string   `json:"event,omitempty"`
	Videos []string `json:"videos"`
}

type SystemInit struct {
	Event             string                 `json:"event,omitempty"`
	SessionId         string                 `json:"session_id"`
	ControlHost       ControlHost            `json:"control_host"`
	ScreenSize        ScreenSize             `json:"screen_size"`
	Sessions          map[string]SessionData `json:"sessions"`
	ImplicitHosting   bool                   `json:"implicit_hosting"`
	ScreencastEnabled bool                   `json:"screencast_enabled"`
	WebRTC            SystemWebRTC           `json:"webrtc"`
}

type SystemAdmin struct {
	Event           string          `json:"event,omitempty"`
	ScreenSizesList []ScreenSize    `json:"screen_sizes_list"`
	BroadcastStatus BroadcastStatus `json:"broadcast_status"`
}

type SystemDisconnect struct {
	Event   string `json:"event,omitempty"`
	Message string `json:"message"`
}

/////////////////////////////
// Signal
/////////////////////////////

type SignalProvide struct {
	Event      string            `json:"event,omitempty"`
	SDP        string            `json:"sdp"`
	ICEServers []types.ICEServer `json:"iceservers"`
	Video      string            `json:"video"`
}

type SignalCandidate struct {
	Event string `json:"event,omitempty"`
	*webrtc.ICECandidateInit
}

type SignalAnswer struct {
	Event string `json:"event,omitempty"`
	SDP   string `json:"sdp"`
}

type SignalVideo struct {
	Event string `json:"event,omitempty"`
	Video string `json:"video"`
}

/////////////////////////////
// Session
/////////////////////////////

type SessionID struct {
	Event string `json:"event,omitempty"`
	ID    string `json:"id"`
}

type MemberProfile struct {
	Event string `json:"event,omitempty"`
	ID    string `json:"id"`
	*types.MemberProfile
}

type SessionState struct {
	Event string `json:"event,omitempty"`
	ID    string `json:"id"`
	*types.SessionState
}

type SessionData struct {
	Event   string              `json:"event,omitempty"`
	ID      string              `json:"id"`
	Profile types.MemberProfile `json:"profile"`
	State   types.SessionState  `json:"state"`
}

/////////////////////////////
// Control
/////////////////////////////

type ControlHost struct {
	Event   string `json:"event,omitempty"`
	HasHost bool   `json:"has_host"`
	HostID  string `json:"host_id,omitempty"`
}

// TODO: New.
type ControlMove struct {
	Event string `json:"event,omitempty"`
	X     uint16 `json:"x"`
	Y     uint16 `json:"y"`
}

// TODO: New.
type ControlScroll struct {
	Event string `json:"event,omitempty"`
	X     int16  `json:"x"`
	Y     int16  `json:"y"`
}

// TODO: New.
type ControlKey struct {
	Event string `json:"event,omitempty"`
	Key   uint32 `json:"key"`
}

/////////////////////////////
// Screen
/////////////////////////////

type ScreenSize struct {
	Event  string `json:"event,omitempty"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Rate   int16  `json:"rate"`
}

/////////////////////////////
// Clipboard
/////////////////////////////

type ClipboardData struct {
	Event string `json:"event,omitempty"`
	Text  string `json:"text"`
}

/////////////////////////////
// Keyboard
/////////////////////////////

type KeyboardMap struct {
	Event   string `json:"event,omitempty"`
	Layout  string `json:"layout"`
	Variant string `json:"variant"`
}

type KeyboardModifiers struct {
	Event    string `json:"event,omitempty"`
	CapsLock *bool  `json:"capslock"`
	NumLock  *bool  `json:"numlock"`
}

/////////////////////////////
// Broadcast
/////////////////////////////

type BroadcastStatus struct {
	Event    string `json:"event,omitempty"`
	IsActive bool   `json:"is_active"`
	URL      string `json:"url,omitempty"`
}

/////////////////////////////
// Send (opaque comunication channel)
/////////////////////////////

type SendUnicast struct {
	Event    string      `json:"event,omitempty"`
	Sender   string      `json:"sender"`
	Receiver string      `json:"receiver"`
	Subject  string      `json:"subject"`
	Body     interface{} `json:"body"`
}

type SendBroadcast struct {
	Event   string      `json:"event,omitempty"`
	Sender  string      `json:"sender"`
	Subject string      `json:"subject"`
	Body    interface{} `json:"body"`
}
