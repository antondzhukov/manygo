package manygo

type RateLimit uint16

const (
	RateLimitPageGetInfo           RateLimit = 100
	RateLimitPageCreateTag         RateLimit = 10
	RateLimitPageGetTags           RateLimit = 100
	RateLimitPageRemoveTag         RateLimit = 10
	RateLimitPageRemoveTageByName  RateLimit = 10
	RateLimitPageCreateCustomField RateLimit = 10
	RateLimitPageGetGrowthTools    RateLimit = 100
	RateLimitPageGetFlows          RateLimit = 10
	RateLimitPageGetCustomFields   RateLimit = 100
	RateLimitPageGetOtnTopics      RateLimit = 100
	RateLimitPageGetBotFields      RateLimit = 100
	RateLimitPageCreateBotField    RateLimit = 10
	RateLimitPageSetBotField       RateLimit = 10
	RateLimitPageSetBotFieldByName RateLimit = 10
	RateLimitPageSetBotFields      RateLimit = 10

	RateLimitSubscriberGetInfo               RateLimit = 10
	RateLimitSubscriberFindByName            RateLimit = 100
	RateLimitSubscriberGetInfoByUserRef      RateLimit = 1000
	RateLimitSubscriberFindByCustomField     RateLimit = 100
	RateLimitSubscriberFindBySystemField     RateLimit = 100
	RateLimitSubscriberAddTag                RateLimit = 10
	RateLimitSubscriberAddTagByName          RateLimit = 10
	RateLimitSubscriberRemoveTag             RateLimit = 10
	RateLimitSubscriberRemoveTagByName       RateLimit = 10
	RateLimitSubscriberSetCustomField        RateLimit = 10
	RateLimitSubscriberSetCustomFields       RateLimit = 10
	RateLimitSubscriberSetCustomFieldByName  RateLimit = 10
	RateLimitSubscriberVerifyBySignedRequest RateLimit = 10
	RateLimitSubscriberCreateSubscriber      RateLimit = 10
	RateLimitSubscriberUpdateSubscriber      RateLimit = 10

	RateLimitSendingSendContent          RateLimit = 25
	RateLimitSendingSendContentByUserRef RateLimit = 25
	RateLimitSendingSendFlow             RateLimit = 25
)
