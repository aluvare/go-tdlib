// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
	"fmt"
)

// Messages Contains a list of messages
type Messages struct {
	tdCommon
	TotalCount int32     `json:"total_count"` // Approximate total count of messages found
	Messages   []Message `json:"messages"`    // List of messages; messages may be null
}

// MessageType return the string telegram-type of Messages
func (messages *Messages) MessageType() string {
	return "messages"
}

// NewMessages creates a new Messages
//
// @param totalCount Approximate total count of messages found
// @param messages List of messages; messages may be null
func NewMessages(totalCount int32, messages []Message) *Messages {
	messagesTemp := Messages{
		tdCommon:   tdCommon{Type: "messages"},
		TotalCount: totalCount,
		Messages:   messages,
	}

	return &messagesTemp
}

// GetMessages Returns information about messages. If a message is not found, returns null on the corresponding position of the result
// @param chatID Identifier of the chat the messages belong to
// @param messageIDs Identifiers of the messages to get
func (client *Client) GetMessages(chatID int64, messageIDs []int64) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "getMessages",
		"chat_id":     chatID,
		"message_ids": messageIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// GetChatHistory Returns messages in a chat. The messages are returned in a reverse chronological order (i.e., in order of decreasing message_id). For optimal performance the number of returned messages is chosen by the library. This is an offline request if only_local is true
// @param chatID Chat identifier
// @param fromMessageID Identifier of the message starting from which history must be fetched; use 0 to get results from the last message
// @param offset Specify 0 to get results from exactly the from_message_id or a negative offset up to 99 to get additionally some newer messages
// @param limit The maximum number of messages to be returned; must be positive and can't be greater than 100. If the offset is negative, the limit must be greater than or equal to -offset. Fewer messages may be returned than specified by the limit, even if the end of the message history has not been reached
// @param onlyLocal If true, returns only messages that are available locally without sending network requests
func (client *Client) GetChatHistory(chatID int64, fromMessageID int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "getChatHistory",
		"chat_id":         chatID,
		"from_message_id": fromMessageID,
		"offset":          offset,
		"limit":           limit,
		"only_local":      onlyLocal,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// GetMessageThreadHistory Returns messages in a message thread of a message. Can be used only if message.can_get_message_thread == true. Message thread of a channel message is in the channel's linked supergroup. The messages are returned in a reverse chronological order (i.e., in order of decreasing message_id). For optimal performance the number of returned messages is chosen by the library
// @param chatID Chat identifier
// @param messageID Message identifier, which thread history needs to be returned
// @param fromMessageID Identifier of the message starting from which history must be fetched; use 0 to get results from the last message
// @param offset Specify 0 to get results from exactly the from_message_id or a negative offset up to 99 to get additionally some newer messages
// @param limit The maximum number of messages to be returned; must be positive and can't be greater than 100. If the offset is negative, the limit must be greater than or equal to -offset. Fewer messages may be returned than specified by the limit, even if the end of the message thread history has not been reached
func (client *Client) GetMessageThreadHistory(chatID int64, messageID int64, fromMessageID int64, offset int32, limit int32) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "getMessageThreadHistory",
		"chat_id":         chatID,
		"message_id":      messageID,
		"from_message_id": fromMessageID,
		"offset":          offset,
		"limit":           limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SearchChatMessages Searches for messages with given words in the chat. Returns the results in reverse chronological order, i.e. in order of decreasing message_id. Cannot be used in secret chats with a non-empty query (searchSecretMessages should be used instead), or without an enabled message database. For optimal performance the number of returned messages is chosen by the library
// @param chatID Identifier of the chat in which to search messages
// @param query Query to search for
// @param sender If not null, only messages sent by the specified sender will be returned. Not supported in secret chats
// @param fromMessageID Identifier of the message starting from which history must be fetched; use 0 to get results from the last message
// @param offset Specify 0 to get results from exactly the from_message_id or a negative offset to get the specified message and some newer messages
// @param limit The maximum number of messages to be returned; must be positive and can't be greater than 100. If the offset is negative, the limit must be greater than -offset. Fewer messages may be returned than specified by the limit, even if the end of the message history has not been reached
// @param filter Filter for message content in the search results
// @param messageThreadID If not 0, only messages in the specified thread will be returned; supergroups only
func (client *Client) SearchChatMessages(chatID int64, query string, sender MessageSender, fromMessageID int64, offset int32, limit int32, filter SearchMessagesFilter, messageThreadID int64) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "searchChatMessages",
		"chat_id":           chatID,
		"query":             query,
		"sender":            sender,
		"from_message_id":   fromMessageID,
		"offset":            offset,
		"limit":             limit,
		"filter":            filter,
		"message_thread_id": messageThreadID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SearchMessages Searches for messages in all chats except secret chats. Returns the results in reverse chronological order (i.e., in order of decreasing (date, chat_id, message_id)). For optimal performance the number of returned messages is chosen by the library
// @param chatList Chat list in which to search messages; pass null to search in all chats regardless of their chat list
// @param query Query to search for
// @param offsetDate The date of the message starting from which the results should be fetched. Use 0 or any date in the future to get results from the last message
// @param offsetChatID The chat identifier of the last found message, or 0 for the first request
// @param offsetMessageID The message identifier of the last found message, or 0 for the first request
// @param limit The maximum number of messages to be returned; up to 100. Fewer messages may be returned than specified by the limit, even if the end of the message history has not been reached
// @param filter Filter for message content in the search results; searchMessagesFilterCall, searchMessagesFilterMissedCall, searchMessagesFilterMention, searchMessagesFilterUnreadMention, searchMessagesFilterFailedToSend and searchMessagesFilterPinned are unsupported in this function
// @param minDate If not 0, the minimum date of the messages to return
// @param maxDate If not 0, the maximum date of the messages to return
func (client *Client) SearchMessages(chatList ChatList, query string, offsetDate int32, offsetChatID int64, offsetMessageID int64, limit int32, filter SearchMessagesFilter, minDate int32, maxDate int32) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "searchMessages",
		"chat_list":         chatList,
		"query":             query,
		"offset_date":       offsetDate,
		"offset_chat_id":    offsetChatID,
		"offset_message_id": offsetMessageID,
		"limit":             limit,
		"filter":            filter,
		"min_date":          minDate,
		"max_date":          maxDate,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SearchCallMessages Searches for call messages. Returns the results in reverse chronological order (i. e., in order of decreasing message_id). For optimal performance the number of returned messages is chosen by the library
// @param fromMessageID Identifier of the message from which to search; use 0 to get results from the last message
// @param limit The maximum number of messages to be returned; up to 100. Fewer messages may be returned than specified by the limit, even if the end of the message history has not been reached
// @param onlyMissed If true, returns only messages with missed calls
func (client *Client) SearchCallMessages(fromMessageID int64, limit int32, onlyMissed bool) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "searchCallMessages",
		"from_message_id": fromMessageID,
		"limit":           limit,
		"only_missed":     onlyMissed,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SearchChatRecentLocationMessages Returns information about the recent locations of chat members that were sent to the chat. Returns up to 1 location message per user
// @param chatID Chat identifier
// @param limit The maximum number of messages to be returned
func (client *Client) SearchChatRecentLocationMessages(chatID int64, limit int32) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "searchChatRecentLocationMessages",
		"chat_id": chatID,
		"limit":   limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// GetActiveLiveLocationMessages Returns all active live locations that should be updated by the application. The list is persistent across application restarts only if the message database is used
func (client *Client) GetActiveLiveLocationMessages() (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getActiveLiveLocationMessages",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// GetChatScheduledMessages Returns all scheduled messages in a chat. The messages are returned in a reverse chronological order (i.e., in order of decreasing message_id)
// @param chatID Chat identifier
func (client *Client) GetChatScheduledMessages(chatID int64) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatScheduledMessages",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SendMessageAlbum Sends messages grouped together into an album. Currently only audio, document, photo and video messages can be grouped into an album. Documents and audio files can be only grouped in an album with messages of the same type. Returns sent messages
// @param chatID Target chat
// @param messageThreadID If not 0, a message thread identifier in which the messages will be sent
// @param replyToMessageID Identifier of a message to reply to or 0
// @param options Options to be used to send the messages
// @param inputMessageContents Contents of messages to be sent
func (client *Client) SendMessageAlbum(chatID int64, messageThreadID int64, replyToMessageID int64, options *MessageSendOptions, inputMessageContents []InputMessageContent) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                  "sendMessageAlbum",
		"chat_id":                chatID,
		"message_thread_id":      messageThreadID,
		"reply_to_message_id":    replyToMessageID,
		"options":                options,
		"input_message_contents": inputMessageContents,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// ForwardMessages Forwards previously sent messages. Returns the forwarded messages in the same order as the message identifiers passed in message_ids. If a message can't be forwarded, null will be returned instead of the message
// @param chatID Identifier of the chat to which to forward messages
// @param fromChatID Identifier of the chat from which to forward messages
// @param messageIDs Identifiers of the messages to forward. Message identifiers must be in a strictly increasing order
// @param options Options to be used to send the messages
// @param sendCopy True, if content of the messages needs to be copied without links to the original messages. Always true if the messages are forwarded to a secret chat
// @param removeCaption True, if media caption of message copies needs to be removed. Ignored if send_copy is false
func (client *Client) ForwardMessages(chatID int64, fromChatID int64, messageIDs []int64, options *MessageSendOptions, sendCopy bool, removeCaption bool) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "forwardMessages",
		"chat_id":        chatID,
		"from_chat_id":   fromChatID,
		"message_ids":    messageIDs,
		"options":        options,
		"send_copy":      sendCopy,
		"remove_caption": removeCaption,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// ResendMessages Resends messages which failed to send. Can be called only for messages for which messageSendingStateFailed.can_retry is true and after specified in messageSendingStateFailed.retry_after time passed. If a message is re-sent, the corresponding failed to send message is deleted. Returns the sent messages in the same order as the message identifiers passed in message_ids. If a message can't be re-sent, null will be returned instead of the message
// @param chatID Identifier of the chat to send messages
// @param messageIDs Identifiers of the messages to resend. Message identifiers must be in a strictly increasing order
func (client *Client) ResendMessages(chatID int64, messageIDs []int64) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "resendMessages",
		"chat_id":     chatID,
		"message_ids": messageIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}
