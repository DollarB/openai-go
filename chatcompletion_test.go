// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openai_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/internal/testutil"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/shared"
)

func TestChatCompletionNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{openai.ChatCompletionSystemMessageParam{
			Content: openai.F[openai.ChatCompletionSystemMessageParamContentUnion](shared.UnionString("string")),
			Role:    openai.F(openai.ChatCompletionSystemMessageParamRoleSystem),
			Name:    openai.F("name"),
		}}),
		Model:            openai.F(openai.ChatModelGPT4o),
		FrequencyPenalty: openai.F(-2.000000),
		FunctionCall:     openai.F[openai.ChatCompletionNewParamsFunctionCallUnion](openai.ChatCompletionNewParamsFunctionCallString(openai.ChatCompletionNewParamsFunctionCallStringNone)),
		Functions: openai.F([]openai.ChatCompletionNewParamsFunction{{
			Description: openai.F("description"),
			Name:        openai.F("name"),
			Parameters: openai.F(shared.FunctionParameters{
				"foo": "bar",
			}),
		}}),
		LogitBias: openai.F(map[string]int64{
			"foo": int64(0),
		}),
		Logprobs:          openai.F(true),
		MaxTokens:         openai.F(int64(0)),
		N:                 openai.F(int64(1)),
		ParallelToolCalls: openai.F(true),
		PresencePenalty:   openai.F(-2.000000),
		ResponseFormat: openai.F[openai.ChatCompletionNewParamsResponseFormatUnion](shared.ResponseFormatTextParam{
			Type: openai.F(shared.ResponseFormatTextTypeText),
		}),
		Seed:        openai.F(int64(-9007199254740991)),
		ServiceTier: openai.F(openai.ChatCompletionNewParamsServiceTierAuto),
		Stop:        openai.F[openai.ChatCompletionNewParamsStopUnion](shared.UnionString("string")),
		StreamOptions: openai.F(openai.ChatCompletionStreamOptionsParam{
			IncludeUsage: openai.F(true),
		}),
		Temperature: openai.F(1.000000),
		ToolChoice:  openai.F[openai.ChatCompletionToolChoiceOptionUnionParam](openai.ChatCompletionToolChoiceOptionString(openai.ChatCompletionToolChoiceOptionStringNone)),
		Tools: openai.F([]openai.ChatCompletionToolParam{{
			Type: openai.F(openai.ChatCompletionToolTypeFunction),
			Function: openai.F(shared.FunctionDefinitionParam{
				Description: openai.F("description"),
				Name:        openai.F("name"),
				Parameters: openai.F(shared.FunctionParameters{
					"foo": "bar",
				}),
				Strict: openai.F(true),
			}),
		}, {
			Type: openai.F(openai.ChatCompletionToolTypeFunction),
			Function: openai.F(shared.FunctionDefinitionParam{
				Description: openai.F("description"),
				Name:        openai.F("name"),
				Parameters: openai.F(shared.FunctionParameters{
					"foo": "bar",
				}),
				Strict: openai.F(true),
			}),
		}, {
			Type: openai.F(openai.ChatCompletionToolTypeFunction),
			Function: openai.F(shared.FunctionDefinitionParam{
				Description: openai.F("description"),
				Name:        openai.F("name"),
				Parameters: openai.F(shared.FunctionParameters{
					"foo": "bar",
				}),
				Strict: openai.F(true),
			}),
		}}),
		TopLogprobs: openai.F(int64(0)),
		TopP:        openai.F(1.000000),
		User:        openai.F("user-1234"),
	})
	if err != nil {
		var apierr *openai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}