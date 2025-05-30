package errcode

var (
	ChannelIdIsNull                  = New(703001, "channel id is null")
	ChannelOwnerIdIsNull             = New(703002, "channel owner id is null")
	ChannelOwnerTypeIsNull           = New(703003, "channel owner type is null")
	ChannelTypeIsNull                = New(703004, "channel type is null")
	ChannelCurrencyIsNull            = New(703005, "channel currency is null")
	ChannelAmountIsNull              = New(703006, "channel amount is null")
	ChannelAmountInvalid             = New(703007, "channel amount invalid")
	ChannelCreateErr                 = New(703008, "channel create err")
	ChannelNotActiveErr              = New(703009, "channel not active")
	ChannelExistsErr                 = New(703010, "channel already exists")
	ChannelNoFoundErr                = New(703011, "channel no found")
	ChannelLoginPwdErr               = New(703012, "channel login pwd err")
	ChannelUpdateReqParamsErr        = New(703013, "channel update req params err")
	ChannelUpdateErr                 = New(703014, "channel update err")
	ChannelUpdateStatusReqParamsErr  = New(703015, "channel update status req params err")
	ChannelUpdateStatusErr           = New(703016, "channel update status err")
	ChannelOwnerTypeNotMatchErr      = New(703017, "channel owner type not match")
	ChannelIdGenerateErr             = New(703018, "channel id generate err")
	ChannelGetReqParamsErr           = New(703019, "channel get req params err")
	ChannelCloseReqParamsErr         = New(703019, "channel close req params err")
	ChannelAmountMustGreaterThanZero = New(703020, "amount must greater than zero")
	ChannelTxTypeIsNull              = New(703021, "channel tx type is null")
	ChannelTxTimeIsNull              = New(703022, "channel tx time is null")
	ChannelTxOrderIdIsNull           = New(703023, "channel tx order id is null")
	ChannelTxDetailIsNull            = New(703024, "channel tx detail is null")
	ChannelCreateReqParamsErr        = New(703025, "channel create req params err")
	ChannelNameIsNull                = New(703026, "channel name is null")
	ChannelAvatarIsNull              = New(703027, "channel avatar is null")
	ChannelSupportRegionIsNull       = New(703028, "channel support region is null")
	ChannelSupportCurrencyIsNull     = New(703029, "channel support currency is null")
	ChannelSupportPayModeIsNull      = New(703030, "channel support pay mode is null")
	ChannelStatusIsNull              = New(703031, "channel status is null")
	ChannelPhoneIsNull               = New(703032, "channel phone is null")
	ChannelEmailIsNull               = New(703033, "channel email is null")
	ChannelRemarkIsNull              = New(703034, "channel remark is null")
	ChannelParamsIsNull              = New(703035, "channel params is null")
	ChannelParamsCreateErr           = New(703036, "channel params create err")
	ChannelParamsUpdateErr           = New(703037, "channel params update err")
	ChannelPayUrlIsNull              = New(703038, "channel pay url is null")
	ChannelCallbackUrlIsNull         = New(703039, "channel callback url is null")
	ChannelActivateErr               = New(703040, "channel activate err")
	ChannelListErr                   = New(703041, "channel list err")
	ChannelOrderIdIsNull             = New(703042, "channel order id is null")
	ChannelOrderIdGenerateErr        = New(703043, "channel order id generate err")
	ChannelOrderIdInvalid            = New(703044, "channel order id invalid")
	ChannelOrderMerchantIdIsNull     = New(703045, "channel order merchant id is null")
	ChannelOrderMerchantIdInvalid    = New(703046, "channel order merchant id invalid")
	ChannelOrderTxOrderIdIsNull      = New(703047, "channel order tx order id is null")
	ChannelOrderTxOrderIdInvalid     = New(703048, "channel order tx order id invalid")
	ChannelOrderTxTypeIsNull         = New(703049, "channel order tx type is null")
	ChannelOrderTxTypeInvalid        = New(703050, "channel order tx type invalid")
	ChannelOrderTxAmountIsNull       = New(703051, "channel order tx amount is null")
	ChannelOrderTxAmountInvalid      = New(703052, "channel order tx amount invalid")
	ChannelOrderTxFeeAmountIsNull    = New(703053, "channel order tx fee amount is null")
	ChannelOrderTxFeeAmountInvalid   = New(703054, "channel order tx fee amount invalid")
	ChannelOrderTxCurrencyIsNull     = New(703055, "channel order tx currency is null")
	ChannelOrderTxCurrencyInvalid    = New(703056, "channel order tx currency invalid")
	ChannelOrderTxTimeIsNull         = New(703057, "channel order tx time is null")
	ChannelOrderTxTimeInvalid        = New(703058, "channel order tx time invalid")
	ChannelOrderTxDetailIsNull       = New(703059, "channel order tx detail is null")
	ChannelOrderTxDetailInvalid      = New(703060, "channel order tx detail invalid")
	ChannelOrderStatusIsNull         = New(703061, "channel order status is null")
	ChannelOrderStatusInvalid        = New(703062, "channel order status invalid")
	ChannelOrderCreateErr            = New(703063, "channel order create err")
	ChannelOrderReqParamsErr         = New(703064, "channel order  req params err")
	ChannelOrderNoFoundErr           = New(703065, "channel order no found")
	ChannelOrderUpdateErr            = New(703066, "channel order update err")
	ChannelOrderStatusNotInitErr     = New(703067, "channel order status not init")
)
