package constants

const (
	CTX_ID_PARAM      = "id"
	CTX_KEY_ROLE_NAME = "user_id"

	ENUM_ROLE_ADMIN = "admin"
	ENUM_ROLE_USER  = "user"

	ENUM_RUN_PRODUCTION = "production"
	ENUM_RUN_TESTING    = "testing"

	ENUM_PAGINATION_LIMIT = 10
	ENUM_PAGINATION_PAGE  = 1

	JWT_EXPIRE_TIME_IN_MINS = 120

	SA_CODE_LIMIT             = 15
	SA_DISCOUNT_PERCENTAGE    = 0.8 // WARNING: 100 - X , X = actual discount
	INVOICE_DURATION          = 900 // 15m
	SUCCESS_REDIRECT_PATH     = "/dashboard/events"
	FAILED_REDIRECT_PATH      = "/dashboard/events"
	SUCCESS_REDIRECT_PATH_WSN = "/events/welcome/registration"

	DAMEN_UID = "29d49bd7-e597-4cd3-8bb6-db7e222e0021"

	TPS_ZOOM_1 = "http://its.id/m/ZoomTPS1"
	TPS_ZOOM_2 = "http://its.id/m/ZoomTPS2"

	TKA_SAINTEK_1 = "http://its.id/m/ZoomSaintek"
	TKA_SAINTEK_2 = "http://its.id/m/ZoomSaintek"
	TKA_SOSHUM    = "http://its.id/m/ZoomSoshum"

	INVALID_TCOUNT_FORDA = 999999 // forda ga jual tiket segini
)
