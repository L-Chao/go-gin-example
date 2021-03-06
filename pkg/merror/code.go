package merror

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_EXIST_TAG                = 10001
	ERROR_NOT_EXIST_TAG            = 10002
	ERROR_NOT_EXIST_ARTICLE        = 10003
	ERROR_CHECK_EXIST_ARTICLE_FAIL = 10004
	ERROR_GET_ARTICLE_FAIL         = 10005
	ERROR_COUNT_ARTICLE_FAIL       = 10006
	ERROR_GET_ALL_ARTICLE_FAIL     = 10007
	ERROR_ADD_ARTICLE_FAIL         = 10008
	ERROR_EDIE_ARTICLE_FAIL        = 10009
	ERROR_DELETE_ARTICLE_FAIL      = 10010

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
)
