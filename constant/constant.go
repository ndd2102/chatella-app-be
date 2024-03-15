package constant

const (
	ACCOUNT_STATUS_ACTIVATED   = "activated"
	ACCOUNT_STATUS_INACTIVATED = "inactivated"
	ACCOUNT_STATUS_DEACTIVATED = "deactivated"

	ACCOUNT_TYPE_NORMAL = "normal"
	ACCOUNT_TYPE_ADMIN  = "admin"
	ACCOUNT_TYPE_MOD    = "mod"
)

const (
	CLIENT_BASE_URL      = "http://localhost:3000"
	CONFIRM_EMAIL_PATH   = "/confirm-email"
	FORGOT_PASSWORD_PATH = "/forgot-password"
)

const (
	CHANNEL_ROLE_HOST   = "host"
	CHANNEL_ROLE_MEMBER = "member"
)

var DEFAULT_USER_AVATAR_LIST = []string{
	"https://img.lovepik.com/free-png/20220107/lovepik-cartoon-avatar-png-image_401205251_wh860.png",
	"https://img.lovepik.com/element/40144/0421.png_860.png",
	"https://img.lovepik.com/free-png/20220109/lovepik-cartoon-avatar-png-image_401349915_wh860.png",
	"https://img.lovepik.com/free-png/20220108/lovepik-cartoon-avatar-png-image_401302777_wh860.png",
}

var DEFAULT_CHANNEL_AVATAR_LIST = []string{
	"https://img.lovepik.com/free-png/20220109/lovepik-animal-avatar-icon-png-image_401351325_wh860.png",
	"https://banner2.cleanpng.com/20190629/gkj/kisspng-animal-clip-art-cartoon-avatar-image-5d172a123835c5.3066887415617991862302.jpg",
}
