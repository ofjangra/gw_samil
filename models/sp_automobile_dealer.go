package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SPRegDetails struct {
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
	SCWithGST          float64 `json:"sc_with_gst" bson:"sc_with_gst"`
}

type SP_FE_Details struct {
	RegisterationFee   float64 `json:"reg_fee" bson:"reg_fee"`
	Road_tax           float64 `json:"road_tax" bson:"road_tax"`
	Hypothecation      float64 `json:"hypothecation"`
	FitnessFee         float64 `json:"fitness_fee" bson:"fitness_fee"`
	UserCharges        float64 `json:"user_charges" bson:"user_charges"`
	PostalFee          float64 `json:"postal_fee" bson:"postal_fee"`
	RCCard             float64 `json:"rc_card" bson:"rc_card"`
	DepartmentExpenses float64 `json:"department_expenses" bson:"department_expenses"`
	SCWithGST          float64 `json:"sc_with_gst" bson:"sc_with_gst"`
}

type RenewalDetails struct {
	RecieptAmount    float64 `json:"reciept_amount" bson:"reciept_amount"`
	NonRecieptAmount float64 `json:"non_reciept_amount" bson:"non_reciept_amount"`
	SCWithGST        float64 `json:"sc_with_gst" bson:"sc_with_gst"`
}

type TotalsRegLp struct {
	RecieptAmount    float64 `json:"reciept_amount" bson:"reciept_amount"`
	NonRecieptAmount float64 `json:"non_reciept_amount" bson:"non_reciept_amount"`
	ServiceCharges   float64 `json:"service_charges" bson:"service_charges"`
}

type TotalsRegNp struct {
	RecieptAmount    float64 `json:"reciept_amount" bson:"reciept_amount"`
	NonRecieptAmount float64 `json:"non_reciept_amount" bson:"non_reciept_amount"`
	ServiceCharges   float64 `json:"service_charges" bson:"service_charges"`
}

type TotalsRenewals struct {
	FitnessNormal  float64 `json:"fitness_normal" bson:"fitness_normal"`
	FitnessPremium float64 `json:"fitness_premium" bson:"fitness_premium"`
	PermitNP       float64 `json:"permit_np" bson:"permit_np"`
	PermitLP       float64 `json:"permit_lp" bson:"permit_lp"`
}

type OtherServices struct {
	NOC         float64 `json:"noc"`
	Transfer    float64 `json:"transfer"`
	HPCancel    float64 `json:"hp_cancel" bson:"hp_cancel"`
	HPAddition  float64 `json:"hp_add" bson:"hp_add"`
	DuplicateRC float64 `json:"duplicate_rc" bson:"duplicate_rc"`
	Insurance   float64 `json:"insurance"`
	Pollution   float64 `json:"pollution"`
}

type Totals struct {
	LocalPermitRegistration    TotalsRegLp    `json:"reg_lp" bson:"reg_lp"`
	NationalPermitRegistration TotalsRegNp    `json:"reg_np" bson:"reg_np"`
	Renewals                   TotalsRenewals `json:"renewals"`
	OtherServices              OtherServices  `json:"other_services" bson:"other_services"`
}

type TotalsFE struct {
	NewRegistration float64 `json:"new_reg" bson:"new_reg"`

	Renewals struct {
		FitnessNormal  float64 `json:"fitness_normal" bson:"fitness_normal"`
		FitnessPremium float64 `json:"fitness_premium" bson:"fitness_premium"`
	} `json:"renewals"`
	OtherServices struct {
		NOC         float64 `json:"noc"`
		Transfer    float64 `json:"transfer"`
		HPCancel    float64 `json:"hp_cancel" bson:"hp_cancel"`
		HPAddition  float64 `json:"hp_add" bson:"hp_add"`
		DuplicateRC float64 `json:"duplicate_rc" bson:"duplicate_rc"`
		Insurance   float64 `json:"insurance"`
		Pollution   float64 `json:"pollution"`
	} `json:"other_services" bson:"other_services"`
}

type SP_Loading struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	Passing        float32            `json:"passing"`
	RTO            string             `json:"rto"`
	Location       string             `json:"location"`
	PricingDetails SPRegDetails       `json:"pricing_details" bson:"pricing_details"`
	Renewals       struct {
		FitnessNormal  RenewalDetails `json:"fitness_normal" bson:"fitness_normal"`
		FitnessPremium RenewalDetails `json:"fitness_premium" bson:"fitness_premium"`
		PermitNP       RenewalDetails `json:"permit_np" bson:"permit_np"`
		PermitLP       RenewalDetails `json:"permit_lp" bson:"permit_lp"`
	} `json:"renewals"`

	OtherServices struct {
		NOC         RenewalDetails `json:"noc"`
		Transfer    RenewalDetails `json:"transfer"`
		HPCancel    RenewalDetails `json:"hp_cancel" bson:"hp_cancel"`
		HPAddition  RenewalDetails `json:"hp_add" bson:"hp_add"`
		DuplicateRC RenewalDetails `json:"duplicate_rc" bson:"duplicate_rc"`
		Insurance   float64        `json:"insurance"`
		Pollution   float64        `json:"pollution"`
	} `json:"other_services" bson:"other_services"`

	Totals    Totals             `json:"totals" bson:"totals"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedBy string             `json:"updated_by" bson:"update_by"`
	CreatedOn primitive.DateTime `json:"created_on" bson:"created_on"`
	UpdatedOn primitive.DateTime `json:"updated_on" bson:"updated_on"`
}

type SP_CE struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	Passing        float32            `json:"passing"`
	RTO            string             `json:"rto"`
	Location       string             `json:"location"`
	PricingDetails SPRegDetails       `json:"pricing_details" bson:"pricing_details"`
	Renewals       struct {
		FitnessNormal  RenewalDetails `json:"fitness_normal" bson:"fitness_normal"`
		FitnessPremium RenewalDetails `json:"fitness_premium" bson:"fitness_premium"`
		PermitNP       RenewalDetails `json:"permit_np" bson:"permit_np"`
		PermitLP       RenewalDetails `json:"permit_lp" bson:"permit_lp"`
	} `json:"renewals"`

	OtherServices struct {
		NOC         RenewalDetails `json:"noc"`
		Transfer    RenewalDetails `json:"transfer"`
		HPCancel    RenewalDetails `json:"hp_cancel" bson:"hp_cancel"`
		HPAddition  RenewalDetails `json:"hp_add" bson:"hp_add"`
		DuplicateRC RenewalDetails `json:"duplicate_rc" bson:"duplicate_rc"`
		Insurance   float64        `json:"insurance"`
		Pollution   float64        `json:"pollution"`
	} `json:"other_services" bson:"other_services"`
	Totals    Totals             `json:"totals" bson:"totals"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedBy string             `json:"updated_by" bson:"update_by"`
	CreatedOn primitive.DateTime `json:"created_on" bson:"created_on"`
	UpdatedOn primitive.DateTime `json:"updated_on" bson:"updated_on"`
}

// Bus Registration Pricing

type SP_Bus struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	Seating        int32              `json:"seating"`
	RTO            string             `json:"rto"`
	Location       string             `json:"location"`
	PricingDetails SPRegDetails       `json:"pricing_details" bson:"pricing_details"`
	Renewals       struct {
		FitnessNormal  RenewalDetails `json:"fitness_normal" bson:"fitness_normal"`
		FitnessPremium RenewalDetails `json:"fitness_premium" bson:"fitness_premium"`
		PermitNP       RenewalDetails `json:"permit_np" bson:"permit_np"`
		PermitLP       RenewalDetails `json:"permit_lp" bson:"permit_lp"`
	} `json:"renewals"`

	OtherServices struct {
		NOC         RenewalDetails `json:"noc"`
		Transfer    RenewalDetails `json:"transfer"`
		HPCancel    RenewalDetails `json:"hp_cancel" bson:"hp_cancel"`
		HPAddition  RenewalDetails `json:"hp_add" bson:"hp_add"`
		DuplicateRC RenewalDetails `json:"duplicate_rc" bson:"duplicate_rc"`
		Insurance   float64        `json:"insurance"`
		Pollution   float64        `json:"pollution"`
	} `json:"other_services" bson:"other_services"`
	Totals    Totals             `json:"totals" bson:"totals"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedBy string             `json:"updated_by" bson:"update_by"`
	CreatedOn primitive.DateTime `json:"created_on" bson:"created_on"`
	UpdatedOn primitive.DateTime `json:"updated_on" bson:"updated_on"`
}

// Motor Cab Registration pricing

type SP_Auto_Cab struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	Seating        int32              `json:"seating"`
	RTO            string             `json:"rto"`
	Location       string             `json:"location"`
	PricingDetails SPRegDetails       `json:"pricing_details" bson:"pricing_details"`
	Renewals       struct {
		FitnessNormal  RenewalDetails `json:"fitness_normal" bson:"fitness_normal"`
		FitnessPremium RenewalDetails `json:"fitness_premium" bson:"fitness_premium"`
		PermitNP       RenewalDetails `json:"permit_np" bson:"permit_np"`
		PermitLP       RenewalDetails `json:"permit_lp" bson:"permit_lp"`
	} `json:"renewals"`

	OtherServices struct {
		NOC         RenewalDetails `json:"noc"`
		Transfer    RenewalDetails `json:"transfer"`
		HPCancel    RenewalDetails `json:"hp_cancel" bson:"hp_cancel"`
		HPAddition  RenewalDetails `json:"hp_add" bson:"hp_add"`
		DuplicateRC RenewalDetails `json:"duplicate_rc" bson:"duplicate_rc"`
		Insurance   float64        `json:"insurance"`
		Pollution   float64        `json:"pollution"`
	} `json:"other_services" bson:"other_services"`
	Totals    Totals             `json:"totals" bson:"totals"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedBy string             `json:"updated_by" bson:"update_by"`
	CreatedOn primitive.DateTime `json:"created_on" bson:"created_on"`
	UpdatedOn primitive.DateTime `json:"updated_on" bson:"updated_on"`
}

// Farm Equipment Pricing

type SP_FE struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	Passing        float32            `json:"passing" bson:"passing"`
	RTO            string             `json:"rto"`
	Location       string             `json:"location"`
	PricingDetails SP_FE_Details      `json:"pricing_details" bson:"pricing_details"`
	Renewals       struct {
		FitnessNormal  RenewalDetails `json:"fitness_normal" bson:"fitness_normal"`
		FitnessPremium RenewalDetails `json:"fitness_premium" bson:"fitness_premium"`
	} `json:"renewals"`

	OtherServices struct {
		NOC         RenewalDetails `json:"noc"`
		Transfer    RenewalDetails `json:"transfer"`
		HPCancel    RenewalDetails `json:"hp_cancel" bson:"hp_cancel"`
		HPAddition  RenewalDetails `json:"hp_add" bson:"hp_add"`
		DuplicateRC RenewalDetails `json:"duplicate_rc" bson:"duplicate_rc"`
		Insurance   float64        `json:"insurance"`
		Pollution   float64        `json:"pollution"`
	} `json:"other_services" bson:"other_services"`
	Totals    TotalsFE           `json:"totals" bson:"totals"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedBy string             `json:"updated_by" bson:"update_by"`
	CreatedOn primitive.DateTime `json:"created_on" bson:"created_on"`
	UpdatedOn primitive.DateTime `json:"updated_on" bson:"updated_on"`
}
