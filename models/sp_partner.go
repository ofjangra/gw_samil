package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SP_Partner_Details struct {
	GovtFee        float64 `json:"govt_fee" bson:"govt_fee"`
	Reimbursement  float64 `json:"reimbursement" bson:"reimbursement"`
	ServiceCharges float64 `json:"service_charges" bson:"service_charges"`
	GSTCharges     float64 `json:"gst_charges" bson:"gst_charges"`
	VendorCharges  float64 `json:"vendor_charges" bson:"vendor_charges"`
	OPE            float64 `json:"ope" bson:"ope"`
	Partner        float64 `json:"partner" bson:"partner"`
}

type SP_Partner_4W struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	VehicleCategory string             `json:"vehicle_category" bson:"vehicle_category"`
	RTO             string             `json:"rto" bson:"rto"`
	Location        string             `json:"location"`
	TransferLCV     SP_Partner_Details `json:"transfer_lcv" bson:"transfer_lcv"`
	TransferHCV     SP_Partner_Details `json:"transfer_hcv" bson:"transfer_hcv"`
	HPCancel        SP_Partner_Details `json:"hp_cancel" bson:"hp_cancel"`
	HPAdd           SP_Partner_Details `json:"hp_add" bson:"hp_add"`
	NOC             SP_Partner_Details `json:"noc"`
	DRC             SP_Partner_Details `json:"drc"`
	PVandVV         bool               `json:"pv_and_vv" bson:"pv_and_vv"`
	PartyVideo      SP_Partner_Details `json:"party_video" bson:"party_video"`
	VehicleVideo    SP_Partner_Details `json:"vehicle_video" bson:"vehicle_video"`
	Insurance       struct {
		Reimbursement float64 `json:"reimbursement" bson:"reimbursement"`
		GSTCharges    float64 `json:"gst_charges" bson:"gst_charges"`
	} `json:"insurance"`
	Pollution struct {
		Reimbursement float64 `json:"reimbursement" bson:"reimbursement"`
		GSTCharges    float64 `json:"gst_charges" bson:"gst_charges"`
	} `json:"pollution"`

	Renewals struct {
		FitnessNormal  SP_Partner_Details `json:"fitness_normal" bson:"fitness_normal"`
		FitnessPremium SP_Partner_Details `json:"fitness_premium" bson:"fitness_premium"`
		PermitNP       SP_Partner_Details `json:"permit_np" bson:"permit_np"`
		PermitLP       SP_Partner_Details `json:"permit_lp" bson:"permit_lp"`
	} `json:"renewals"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedBy string             `json:"updated_by" bson:"update_by"`
	CreatedOn primitive.DateTime `json:"created_on" bson:"created_on"`
	UpdatedOn primitive.DateTime `json:"updated_on" bson:"updated_on"`
}

type SP_Partner_2W struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	VehicleCategory string             `json:"vehicle_category" bson:"vehicle_category"`
	RTO             string             `json:"rto" bson:"rto"`
	Location        string             `json:"location"`
	NOC             SP_Partner_Details `json:"noc"`
	Transfer        SP_Partner_Details `json:"transfer"`
	HPCancel        SP_Partner_Details `json:"hp_cancel" bson:"hp_cancel"`
	HPAdd           SP_Partner_Details `json:"hp_add" bson:"hp_add"`
	DRC             SP_Partner_Details `json:"drc"`
	Pollution       struct {
		Reimbursement float64 `json:"reimbursement" bson:"reimbursement"`
		GSTCharges    float64 `json:"gst_charges" bson:"gst_charges"`
	} `json:"pollution"`
	Insurance struct {
		Reimbursement float64 `json:"reimbursement" bson:"reimbursement"`
		GSTCharges    float64 `json:"gst_charges" bson:"gst_charges"`
	} `json:"insurance"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedBy string             `json:"updated_by" bson:"update_by"`
	CreatedOn primitive.DateTime `json:"created_on" bson:"created_on"`
	UpdatedOn primitive.DateTime `json:"updated_on" bson:"updated_on"`
}

type SP_Partner_3W struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	VehicleCategory string             `json:"vehicle_category" bson:"vehicle_category"`
	RTO             string             `json:"rto" bson:"rto"`
	Location        string             `json:"location"`
	NOC             SP_Partner_Details `json:"noc"`
	Transfer        SP_Partner_Details `json:"transfer"`
	HPCancel        SP_Partner_Details `json:"hp_cancel" bson:"hp_cancel"`
	HPAdd           SP_Partner_Details `json:"hp_add" bson:"hp_add"`
	DRC             SP_Partner_Details `json:"drc"`
	Pollution       struct {
		Reimbursement float64 `json:"reimbursement" bson:"reimbursement"`
		GSTCharges    float64 `json:"gst_charges" bson:"gst_charges"`
	} `json:"pollution"`
	Insurance struct {
		Reimbursement float64 `json:"reimbursement" bson:"reimbursement"`
		GSTCharges    float64 `json:"gst_charges" bson:"gst_charges"`
	} `json:"insurance"`

	Renewals struct {
		FitnessNormal  SP_Partner_Details `json:"fitness_normal" bson:"fitness_normal"`
		FitnessPremium SP_Partner_Details `json:"fitness_premium" bson:"fitness_premium"`
	} `json:"renewals"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedBy string             `json:"updated_by" bson:"update_by"`
	CreatedOn primitive.DateTime `json:"created_on" bson:"created_on"`
	UpdatedOn primitive.DateTime `json:"updated_on" bson:"updated_on"`
}
