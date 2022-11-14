// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	"reflect"

	slack "github.com/slack-go/slack"
)

func AnyPtrToSlackGetConversationsParameters() *slack.GetConversationsParameters {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*slack.GetConversationsParameters))(nil)).Elem()))
	var nullValue *slack.GetConversationsParameters
	return nullValue
}

func EqPtrToSlackGetConversationsParameters(value *slack.GetConversationsParameters) *slack.GetConversationsParameters {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *slack.GetConversationsParameters
	return nullValue
}

func NotEqPtrToSlackGetConversationsParameters(value *slack.GetConversationsParameters) *slack.GetConversationsParameters {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue *slack.GetConversationsParameters
	return nullValue
}

func PtrToSlackGetConversationsParametersThat(matcher pegomock.ArgumentMatcher) *slack.GetConversationsParameters {
	pegomock.RegisterMatcher(matcher)
	var nullValue *slack.GetConversationsParameters
	return nullValue
}
