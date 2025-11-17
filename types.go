package main

type AuthResponse struct {
	Token string `json:"token"`
}


type IndividualResp struct {
	Type  string `json:"type"`
	Nodes []any  `json:"nodes"`
}
