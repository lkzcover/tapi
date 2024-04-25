package tapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lkzcover/tapi/inline"
	"io"
	"net/http"
)

// AnswerInlineQuery - is implement https://core.telegram.org/bots/api#answerinlinequery
type AnswerInlineQuery struct {
	InlineQueryID string                     `json:"inline_query_id"`       // Unique identifier for the answered query
	Results       []inline.QueryResult       `json:"results"`               // A JSON-serialized array of results for the inline query
	CacheTime     *int64                     `json:"cache_time,omitempty"`  // The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
	IsPersonal    *bool                      `json:"is_personal,omitempty"` // Pass True if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query.
	NextOffset    *int64                     `json:"next_offset,omitempty"` // Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don't support pagination. Offset length can't exceed 64 bytes.
	Button        *inline.QueryResultsButton `json:"button,omitempty"`      // A JSON-serialized object describing a button to be shown above inline query results
}

// AnswerInlineQuery - send answer to inline query, implement https://core.telegram.org/bots/api#answerinlinequery
func (obj *Engine) AnswerInlineQuery(answer AnswerInlineQuery) (bool, error) {
	body, err := json.Marshal(answer)
	if err != nil {
		return false, fmt.Errorf("marshal answer error: %w", err)
	}

	resp, err := http.Post(obj.telegramApiURL+obj.telegramBotToken+obj.telegramEnvironment+"/answerInlineQuery", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, fmt.Errorf("send answer to user error: %w", err)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("answer to user error: %d %s", resp.StatusCode, string(body))
	}

	var resultMsg struct {
		Ok     bool `json:"ok"`
		Result bool `json:"result"`
	}

	err = json.Unmarshal(body, &resultMsg)
	if err != nil {
		return false, fmt.Errorf("unmarshal answer %s error: %w", string(body), err)
	}

	return resultMsg.Result, nil
}
