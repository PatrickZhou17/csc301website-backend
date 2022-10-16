package typespec

// CommonRequest set initialization structure 
type CommonRequest struct {
	UserID        int64  `form:"userId" json:"userId"`               // user ID
	SessID        string `form:"sessId" json:"sessId"`               // user token
	UniqueID      string `form:"uniqueId" json:"uniqueId"`           // eqyipment ID
	BundleType    string `form:"bundleType" json:"bundleType"`       
	Role          string `form:"role" json:"role"`                   //
	ClientFrom    string `form:"clientFrom" json:"clientFrom"`       // android ios
	ClientVersion string `form:"clientVersion" json:"clientVersion"` //  6.14.00
}

// request to turn page
type PagerRequest struct {
	Offset int `form:"offset,default=0" json:"offset"`  // the first time to login page set to 0
	Length int `form:"length,default=10" json:"length"` // volume of paage
}
