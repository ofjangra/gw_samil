package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SP_AutomobileDealer_Details struct {
	RegisterationFee   float64 `json:"reg_fee" bson:"reg_fee"`
	Road_tax           float64 `json:"road_tax" bson:"road_tax"`
	Hypothecation      float64 `json:"hypothecation"`
	FitnessFee         float64 `json:"fitness_fee" bson:"fitness_fee"`
	UserCharges        float64 `json:"user_charges" bson:"user_charges"`
	PostalFee          float64 `json:"postal_fee" bson:"postal_fee"`
	RCCard             float64 `json:"rc_card" bson:"rc_card"`
	PermitApplication  float64 `json:"permit_application" bson:"permit_application"`
	AuthorizationFee   float64 `json:"authorization_fee" bson:"authorization_fee"`
	LocalPermitFee     float64 `json:"local_permit_fee" bson:"local_permit_fee"`
	NationalPermitFee  float64 `json:"national_permit_fee" bson:"national_permit_fee"`
	DepartmentExpenses float64 `json:"department_expenses" bson:"department_expenses"`
	ServiceCharges     float64 `json:"service_charges" bson:"service_charges"`
	GSTCharges         float64 `json:"gst_charges" bson:"gst_charges"`
}

type SP_AutomobileDealer struct {
	ID             primitive.ObjectID          `json:"_id" bson:"_id"`
	Passing        float64                     `json:"passing"`
	RTO            string                      `json:"rto"`
	Location       string                      `json:"location"`
	PricingDetails SP_AutomobileDealer_Details `json:"pricing_details" bson:"pricing_details"`
	CreatedBy      string                      `json:"created_by" bson:"created_by"`
	UpdatedBy      string                      `json:"updated_by" bson:"update_by"`
	CreatedOn      primitive.DateTime          `json:"created_on" bson:"created_on"`
	UpdatedOn      primitive.DateTime          `json:"updated_on" bson:"updated_on"`
}
